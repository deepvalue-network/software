package scripts

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
)

type linker struct {
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	parserBuilder                   ParserBuilder
	parser                          Parser
	applicationBuilder              linkers.ApplicationBuilder
	programBuilder                  linkers.ProgramBuilder
	languageBuilder                 linkers.LanguageBuilder
	pathsBuilder                    linkers.PathsBuilder
	scriptBuilder                   linkers.ScriptBuilder
	languageReferenceBuilder        linkers.LanguageReferenceBuilder
	dirPath                         string
	currentPath                     string
}

func createLinker(
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	parserBuilder ParserBuilder,
	parser Parser,
	applicationBuilder linkers.ApplicationBuilder,
	programBuilder linkers.ProgramBuilder,
	languageBuilder linkers.LanguageBuilder,
	pathsBuilder linkers.PathsBuilder,
	scriptBuilder linkers.ScriptBuilder,
	languageReferenceBuilder linkers.LanguageReferenceBuilder,
	dirPath string,
) Linker {
	out := linker{
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		parserBuilder:                   parserBuilder,
		parser:                          parser,
		applicationBuilder:              applicationBuilder,
		programBuilder:                  programBuilder,
		languageBuilder:                 languageBuilder,
		pathsBuilder:                    pathsBuilder,
		scriptBuilder:                   scriptBuilder,
		languageReferenceBuilder:        languageReferenceBuilder,
		dirPath:                         dirPath,
		currentPath:                     "",
	}

	return &out
}

// Execute downloads the dependencies of the program and build a linked Program instance
func (app *linker) Execute(
	program middle.Program,
) (linkers.Program, error) {

	// set the current path:
	currentPath, err := filepath.Abs(app.dirPath)
	if err != nil {
		return nil, err
	}

	app.currentPath = currentPath
	defer func() {
		app.currentPath = ""
	}()

	return app.program(program)
}

func (app *linker) program(
	program middle.Program,
) (linkers.Program, error) {

	// create the program builder:
	builder := app.programBuilder.Create()

	// application:
	if program.IsApplication() {
		appBuilder := app.applicationBuilder.Create()
		application := program.Application()
		if application.HasImports() {
			return nil, errors.New("the internal linker cannot link an application that contain imports")
		}

		name := application.Name()
		ins := application.Instructions()
		tests := application.Tests()
		labels := application.Labels()
		vars := application.Variables()
		version := application.Version()
		app, err := appBuilder.WithName(name).
			WithInstructions(ins).
			WithTests(tests).
			WithLabels(labels).
			WithVariables(vars).
			WithVersion(version).
			Now()

		if err != nil {
			return nil, err
		}

		builder.WithApplication(app)
	}

	// language:
	if program.IsLanguage() {
		language := program.Language()
		app, err := app.languageReference(language, app.parser)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(app)
	}

	// script:
	if program.IsScript() {
		script := program.Script()
		app, err := app.script(script, app.parser)
		if err != nil {
			return nil, err
		}

		builder.WithScript(app)
	}

	// return the built program:
	return builder.Now()
}

func (app *linker) script(
	script middle.Script,
	prevParser Parser,
) (linkers.Script, error) {
	relLangPath := script.LanguagePath()
	langRef, err := app.fileLanguageReference(relLangPath, prevParser)
	if err != nil {
		return nil, err
	}

	relCodePath := script.ScriptPath()
	absCodePath, err := filepath.Abs(filepath.Join(app.dirPath, relCodePath))
	if err != nil {
		return nil, err
	}

	// set the current path:
	app.currentPath = filepath.Dir(absCodePath)

	// read the content:
	content, err := ioutil.ReadFile(absCodePath)
	if err != nil {
		return nil, err
	}

	name := script.Name()
	version := script.Version()
	return app.scriptBuilder.Create().WithLanguage(langRef).WithName(name).WithVersion(version).WithCode(string(content)).Now()
}

func (app *linker) language(
	language middle.Language,
	prevParser Parser,
) (linkers.Language, error) {
	// parse the logic with the parser:
	relLogicsPath := language.LogicsPath()
	absLogicsPath := filepath.Join(app.currentPath, relLogicsPath)
	parsedProgram, err := prevParser.File(absLogicsPath)
	if err != nil {
		return nil, err
	}

	program, err := app.program(parsedProgram)
	if err != nil {
		return nil, err
	}

	if !program.IsApplication() {
		str := fmt.Sprintf("the language was expected to contain an application")
		return nil, errors.New(str)
	}

	tokensPath := language.TokensPath()
	rulesPath := language.RulesPath()
	pathBuider := app.pathsBuilder.Create().WithBaseDir(app.currentPath).WithTokens(tokensPath).WithRules(rulesPath).WithLogics(relLogicsPath)
	if language.HasChannelsPath() {
		chanPath := language.ChannelsPath()
		pathBuider.WithChannels(chanPath)
	}

	paths, err := pathBuider.Now()
	if err != nil {
		return nil, err
	}

	progApp := program.Application()
	root := language.Root()
	patternMatches := language.PatternMatches()
	return app.languageBuilder.Create().WithApplication(progApp).WithPaths(paths).WithRoot(root).WithPatternMatches(patternMatches).Now()
}

func (app *linker) languageReference(
	language middle.Language,
	prevParser Parser,
) (linkers.LanguageReference, error) {
	lang, err := app.language(language, prevParser)
	if err != nil {
		return nil, err
	}

	input := language.InputVariable()
	output := language.OutputVariable()
	return app.languageReferenceBuilder.Create().WithLanguage(lang).WithInputVariable(input).WithOutputVariable(output).Now()
}

func (app *linker) fileLanguageReference(
	relLangPath string,
	prevParser Parser,
) (linkers.LanguageReference, error) {
	// parse the language with the parser:
	absLangPath := filepath.Join(app.currentPath, relLangPath)

	// set the current path:
	app.currentPath = filepath.Dir(absLangPath)

	parsedProgram, err := prevParser.File(absLangPath)
	if err != nil {
		return nil, err
	}

	program, err := app.program(parsedProgram)
	if err != nil {
		return nil, err
	}

	if !program.IsLanguage() {
		str := fmt.Sprintf("the language file (%s) was expected to contain a Language program", absLangPath)
		return nil, errors.New(str)
	}

	return program.Language(), nil
}

func (app *linker) buildGrammarRetrieverCriteria(
	language middle.Language,
	parser Parser,
) (grammar.RetrieverCriteria, error) {
	root := language.Root()
	tokensPath := language.TokensPath()
	rulesPath := language.RulesPath()
	builder := app.grammarRetrieverCriteriaBuilder.Create().
		WithBaseDirPath(app.currentPath).
		WithName(scriptName).
		WithRoot(root).
		WithTokensPath(tokensPath).
		WithRulesPath(rulesPath)

	if language.HasChannelsPath() {
		channelsPath := language.ChannelsPath()
		builder.WithChannelsPath(channelsPath)
	}

	if language.HasExtends() {
		return nil, errors.New("the internal linker cannot link a language that contains extends")
	}

	return builder.Now()
}
