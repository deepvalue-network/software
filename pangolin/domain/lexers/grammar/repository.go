package grammar

import (
	"errors"
	"fmt"
	"path/filepath"
)

type repository struct {
	retrieverCriteriaRepositoryBuilder RetrieverCriteriaRepositoryBuilder
	fileFetcher                        FileFetcher
	builder                            Builder
	ruleAdapter                        RuleAdapter
	tokensAdapter                      TokensAdapter
	tokenPattern                       string
	channelPattern                     string
	replacementTokenPattern            string
}

type grammarJSON struct {
	Root     string            `json:"root"`
	Tokens   string            `json:"tokens"`
	Channels string            `json:"channels"`
	Rules    string            `json:"rules"`
	Extends  map[string]string `json:"extends"`
}

func createRepository(
	retrieverCriteriaRepositoryBuilder RetrieverCriteriaRepositoryBuilder,
	fileFetcher FileFetcher,
	builder Builder,
	ruleAdapter RuleAdapter,
	tokensAdapter TokensAdapter,
) Repository {

	anythingExceptEnd := fmt.Sprintf(anythingExcept, end)
	tokPattern := fmt.Sprintf(
		"(%s)%s%s(%s)%s",
		tokenPattern,
		potentialWhitespaces,
		begin,
		anythingExceptEnd,
		end,
	)

	channelPattern := fmt.Sprintf(
		"(%s)%s%s(%s)%s",
		channelPattern,
		potentialWhitespaces,
		begin,
		anythingExceptEnd,
		end,
	)

	replacementTokenPattern := fmt.Sprintf(
		"(%s)%s<%s(%s)%s",
		grammarNamePattern,
		potentialWhitespaces,
		potentialWhitespaces,
		tokenPattern,
		end,
	)

	out := repository{
		retrieverCriteriaRepositoryBuilder: retrieverCriteriaRepositoryBuilder,
		fileFetcher:                        fileFetcher,
		builder:                            builder,
		ruleAdapter:                        ruleAdapter,
		tokensAdapter:                      tokensAdapter,
		tokenPattern:                       tokPattern,
		channelPattern:                     channelPattern,
		replacementTokenPattern:            replacementTokenPattern,
	}

	return &out
}

// Retrieve retrieves a grammar from criteria
func (app *repository) Retrieve(criteria RetrieverCriteria) (Grammar, error) {
	// if the grammar extends another grammar:
	subGrammars := map[string]Grammar{}
	extends := map[string]Grammar{}
	if criteria.HasExtends() {
		criteriaExtends := criteria.Extends()
		for keyname, criteria := range criteriaExtends {
			gr, err := app.Retrieve(criteria)
			if err != nil {
				return nil, err
			}

			grammar, err := app.builder.Create().WithName(keyname).WithGrammar(gr).Now()
			if err != nil {
				return nil, err
			}

			extends[keyname] = grammar
			subGrammars[grammar.Name()] = grammar
		}
	}

	// create the builder:
	name := criteria.Name()
	root := criteria.Root()
	builder := app.builder.Create().WithName(name).WithRoot(root).WithSubGrammars(subGrammars)

	// create the dir path:
	dirPath := criteria.BaseDirPath()

	// retrieve the rules:
	rulesPath := criteria.RulesPath()
	rulesFilePath := filepath.Join(dirPath, rulesPath)
	rulesScript, err := app.fileFetcher(rulesFilePath)
	if err != nil {
		str := fmt.Sprintf("there was an error reading the rules file (path: %s): %s", rulesFilePath, err.Error())
		return nil, errors.New(str)
	}

	rules, err := app.ruleAdapter.ToRules(string(rulesScript), name)
	if err != nil {
		str := fmt.Sprintf("there was an error while transforming the rules script to []Rule: %s", err.Error())
		return nil, errors.New(str)
	}

	// add the rules to the builder:
	builder.WithRules(rules)

	// retrieve the tokens:
	tokensPath := criteria.TokensPath()
	tokenFilesPath := filepath.Join(dirPath, tokensPath)
	tokenScript, err := app.fileFetcher(tokenFilesPath)
	if err != nil {
		str := fmt.Sprintf("there was an error reading the tokens file (path: %s): %s", tokenFilesPath, err.Error())
		return nil, errors.New(str)
	}

	script := string(tokenScript)
	tokens, err := app.tokensAdapter.ToTokens(script, app.tokenPattern, app.replacementTokenPattern, name, extends, rules)
	if err != nil {
		str := fmt.Sprintf("there was an error while transforming the tokens script to []Token: %s", err.Error())
		return nil, errors.New(str)
	}

	// add the tokens to the builder:
	builder.WithTokens(tokens)

	// retrieve the channels:
	if criteria.HasChannelsPath() {
		channelsPath := criteria.ChannelsPath()
		channelsFilePath := filepath.Join(dirPath, channelsPath)
		channelsScript, err := app.fileFetcher(channelsFilePath)
		if err != nil {
			str := fmt.Sprintf("there was an error reading the channels file (path: %s): %s", channelsFilePath, err.Error())
			return nil, errors.New(str)
		}

		channels, err := app.tokensAdapter.ToTokens(string(channelsScript), app.channelPattern, app.replacementTokenPattern, name, extends, rules)
		if err != nil {
			str := fmt.Sprintf("there was an error while transforming the channels script to []Token: %s", err.Error())
			return nil, errors.New(str)
		}

		// add the channels to the builder:
		builder.WithChannels(channels)
	}

	// returns:
	return builder.Now()
}

// RetrieveFromFile retrieves a grammar from file
func (app *repository) RetrieveFromFile(rootPattern string, name string, filePath string) (Grammar, error) {
	criteriaRepository, err := app.retrieverCriteriaRepositoryBuilder.Create().WithFileFetcher(app.fileFetcher).WithName(name).WithRoot(rootPattern).Now()
	if err != nil {
		return nil, err
	}

	criteria, err := criteriaRepository.Retrieve(filePath)
	if err != nil {
		return nil, err
	}

	return app.Retrieve(criteria)
}
