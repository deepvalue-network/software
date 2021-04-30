package application

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	language_applications "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
)

type linker struct {
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	parserBuilder                   ParserBuilder
	parser                          Parser
	applicationBuilder              linkers.ApplicationBuilder
	languageBuilder                 linkers.LanguageBuilder
	programBuilder                  linkers.ProgramBuilder
	languageDefinitionBuilder       linkers.LanguageDefinitionBuilder
	pathsBuilder                    linkers.PathsBuilder
	scriptBuilder                   linkers.ScriptBuilder
	languageReferenceBuilder        linkers.LanguageReferenceBuilder
	languageApplicationBuilder      linkers.LanguageApplicationBuilder
	dirPath                         string
	currentPath                     string
}

func createLinker(
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	parserBuilder ParserBuilder,
	parser Parser,
	applicationBuilder linkers.ApplicationBuilder,
	languageBuilder linkers.LanguageBuilder,
	programBuilder linkers.ProgramBuilder,
	languageDefinitionBuilder linkers.LanguageDefinitionBuilder,
	pathsBuilder linkers.PathsBuilder,
	scriptBuilder linkers.ScriptBuilder,
	languageReferenceBuilder linkers.LanguageReferenceBuilder,
	languageApplicationBuilder linkers.LanguageApplicationBuilder,
	dirPath string,
) Linker {
	out := linker{
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		parserBuilder:                   parserBuilder,
		parser:                          parser,
		applicationBuilder:              applicationBuilder,
		languageBuilder:                 languageBuilder,
		programBuilder:                  programBuilder,
		languageDefinitionBuilder:       languageDefinitionBuilder,
		pathsBuilder:                    pathsBuilder,
		scriptBuilder:                   scriptBuilder,
		languageReferenceBuilder:        languageReferenceBuilder,
		languageApplicationBuilder:      languageApplicationBuilder,
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
		appli := program.Application()
		linkedApp, err := app.application(appli)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(linkedApp)
	}

	// language:
	if program.IsLanguage() {
		language := program.Language()
		app, err := app.language(language, app.parser)
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

func (app *linker) application(
	appli applications.Application,
) (linkers.Application, error) {
	head := appli.Head()
	if head.HasImports() {
		return nil, errors.New("the internal linker cannot link an application that contain imports")
	}

	name := head.Name()
	version := head.Version()

	ins := appli.Main()
	tests := appli.Tests()
	labels := appli.Labels()
	return app.applicationBuilder.Create().WithName(name).
		WithInstructions(ins).
		WithTests(tests).
		WithLabels(labels).
		WithVersion(version).
		Now()
}

func (app *linker) script(
	script scripts.Script,
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
	language languages.Language,
	prevParser Parser,
) (linkers.Language, error) {
	builder := app.languageBuilder.Create()
	if language.IsDefinition() {
		def := language.Definition()
		langRef, err := app.languageReference(def, prevParser)
		if err != nil {
			return nil, err
		}

		builder.WithReference(langRef)
	}

	if language.IsApplication() {
		langApp := language.Application()
		app, err := app.languageApplication(langApp, prevParser)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(app)
	}

	return builder.Now()
}

func (app *linker) languageApplication(
	langApp language_applications.Application,
	prevParser Parser,
) (linkers.LanguageApplication, error) {
	head := langApp.Head()
	if head.HasImports() {
		return nil, errors.New("the internal linker cannot link a language application that contain imports")
	}

	name := head.Name()
	version := head.Version()

	ins := langApp.Main()
	tests := langApp.Tests()
	labels := langApp.Labels()
	return app.languageApplicationBuilder.Create().WithName(name).
		WithInstructions(ins).
		WithTests(tests).
		WithLabels(labels).
		WithVersion(version).
		Now()
}

func (app *linker) languageDefinition(
	def definitions.Definition,
	prevParser Parser,
) (linkers.LanguageDefinition, error) {
	// parse the logic with the parser:
	relLogicsPath := def.LogicsPath()
	absLogicsPath := filepath.Join(app.currentPath, relLogicsPath)
	parsedProgram, err := prevParser.File(absLogicsPath)
	if err != nil {
		return nil, err
	}

	program, err := app.program(parsedProgram)
	if err != nil {
		return nil, err
	}

	strErr := errors.New("the language definition was expected to contain a language application")
	if !program.IsLanguage() {
		return nil, strErr
	}

	progLang := program.Language()
	if !progLang.IsApplication() {
		return nil, strErr
	}

	tokensPath := def.TokensPath()
	rulesPath := def.RulesPath()
	pathBuider := app.pathsBuilder.Create().WithBaseDir(app.currentPath).WithTokens(tokensPath).WithRules(rulesPath).WithLogics(relLogicsPath)
	if def.HasChannelsPath() {
		chanPath := def.ChannelsPath()
		pathBuider.WithChannels(chanPath)
	}

	paths, err := pathBuider.Now()
	if err != nil {
		return nil, err
	}

	root := def.Root()
	langApp := progLang.Application()
	patternMatches := def.PatternMatches()
	return app.languageDefinitionBuilder.Create().WithApplication(langApp).WithPaths(paths).WithRoot(root).WithPatternMatches(patternMatches).Now()
}

func (app *linker) languageReference(
	def definitions.Definition,
	prevParser Parser,
) (linkers.LanguageReference, error) {
	langDef, err := app.languageDefinition(def, prevParser)
	if err != nil {
		return nil, err
	}

	input := def.InputVariable()
	return app.languageReferenceBuilder.Create().WithDefinition(langDef).WithInputVariable(input).Now()
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
		str := fmt.Sprintf("the language file (%s) was expected to contain a Language Definition program", absLangPath)
		return nil, errors.New(str)
	}

	language := program.Language()
	if !language.IsApplication() {
		str := fmt.Sprintf("the language file (%s) was expected to contain a Language Definition program", absLangPath)
		return nil, errors.New(str)
	}

	return language.Reference(), nil
}

func (app *linker) buildGrammarRetrieverCriteria(
	def definitions.Definition,
	parser Parser,
) (grammar.RetrieverCriteria, error) {
	root := def.Root()
	tokensPath := def.TokensPath()
	rulesPath := def.RulesPath()
	builder := app.grammarRetrieverCriteriaBuilder.Create().
		WithBaseDirPath(app.currentPath).
		WithName(scriptName).
		WithRoot(root).
		WithTokensPath(tokensPath).
		WithRulesPath(rulesPath)

	if def.HasChannelsPath() {
		channelsPath := def.ChannelsPath()
		builder.WithChannelsPath(channelsPath)
	}

	if def.HasExtends() {
		return nil, errors.New("the internal linker cannot link a language that contains extends")
	}

	return builder.Now()
}
