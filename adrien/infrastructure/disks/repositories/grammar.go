package repositories

import (
	"io/ioutil"

	"github.com/deepvalue-network/software/adrien/domain/grammars"
	"github.com/deepvalue-network/software/adrien/domain/rules"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

type grammar struct {
	builder              grammars.Builder
	fileRepository       grammars.FileRepository
	ruleAdapter          rules.Adapter
	tokensAdapterBuilder tokens.AdapterBuilder
}

func createGrammar(
	builder grammars.Builder,
	fileRepository grammars.FileRepository,
	ruleAdapter rules.Adapter,
	tokensAdapterBuilder tokens.AdapterBuilder,
) grammars.Repository {
	out := grammar{
		builder:              builder,
		fileRepository:       fileRepository,
		ruleAdapter:          ruleAdapter,
		tokensAdapterBuilder: tokensAdapterBuilder,
	}

	return &out
}

// Retrieve retrieves a grammar from path, from disk
func (app *grammar) Retrieve(relativePath string) (grammars.Grammar, error) {
	file, err := app.fileRepository.Retrieve(relativePath)
	if err != nil {
		return nil, err
	}

	rulesData, err := ioutil.ReadFile(file.RulesPath())
	if err != nil {
		return nil, err
	}

	tokensData, err := ioutil.ReadFile(file.TokensPath())
	if err != nil {
		return nil, err
	}

	rules, err := app.ruleAdapter.ToRules(string(rulesData))
	if err != nil {
		return nil, err
	}

	tokensAdapter, err := app.tokensAdapterBuilder.Create().WithRules(rules).Now()
	if err != nil {
		return nil, err
	}

	tokens, err := tokensAdapter.ToTokens(string(tokensData))
	if err != nil {
		return nil, err
	}

	root := file.Root()
	builder := app.builder.Create().WithRoot(root).WithRules(rules).WithTokens(tokens)
	if file.HasChannelsPath() {
		channelsData, err := ioutil.ReadFile(file.ChannelsPath())
		if err != nil {
			return nil, err
		}

		channels, err := tokensAdapter.ToTokens(string(channelsData))
		if err != nil {
			return nil, err
		}

		builder.WithChannels(channels)
	}

	return builder.Now()
}
