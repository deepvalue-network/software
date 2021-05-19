package linkers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	language_applications "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type linker struct {
	prevParser                      parsers.Parser
	middleAdapter                   middle.Adapter
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	applicationBuilder              ApplicationBuilder
	languageBuilder                 LanguageBuilder
	programBuilder                  ProgramBuilder
	languageDefinitionBuilder       LanguageDefinitionBuilder
	pathsBuilder                    PathsBuilder
	scriptBuilder                   ScriptBuilder
	languageReferenceBuilder        LanguageReferenceBuilder
	languageApplicationBuilder      LanguageApplicationBuilder
	dirPath                         string
	currentPath                     string
}

func createLinker(
	middleAdapter middle.Adapter,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	applicationBuilder ApplicationBuilder,
	languageBuilder LanguageBuilder,
	programBuilder ProgramBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	pathsBuilder PathsBuilder,
	scriptBuilder ScriptBuilder,
	languageReferenceBuilder LanguageReferenceBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
	prevParser parsers.Parser,
	dirPath string,
) Linker {
	out := linker{
		middleAdapter:                   middleAdapter,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		applicationBuilder:              applicationBuilder,
		languageBuilder:                 languageBuilder,
		programBuilder:                  programBuilder,
		languageDefinitionBuilder:       languageDefinitionBuilder,
		pathsBuilder:                    pathsBuilder,
		scriptBuilder:                   scriptBuilder,
		languageReferenceBuilder:        languageReferenceBuilder,
		languageApplicationBuilder:      languageApplicationBuilder,
		prevParser:                      prevParser,
		dirPath:                         dirPath,
		currentPath:                     "",
	}

	return &out
}

// Execute links a parsed program into a linked program
func (app *linker) Execute(parsed parsers.Program) (Program, error) {
	middleProgram, err := app.middleAdapter.ToProgram(parsed)
	if err != nil {
		return nil, err
	}

	return app.program(middleProgram)
}

func (app *linker) program(
	program middle.Program,
) (Program, error) {

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
		app, err := app.language(language, app.prevParser)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(app)
	}

	// script:
	if program.IsScript() {
		script := program.Script()
		app, err := app.script(script, app.prevParser)
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
) (Application, error) {
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
	prevParser parsers.Parser,
) (Script, error) {
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
	return app.scriptBuilder.Create().
		WithLanguage(langRef).
		WithName(name).
		WithVersion(version).
		WithCode(string(content)).
		Now()
}

func (app *linker) language(
	language languages.Language,
	prevParser parsers.Parser,
) (Language, error) {
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
	prevParser parsers.Parser,
) (LanguageApplication, error) {
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
	prevParser parsers.Parser,
) (LanguageDefinition, error) {
	// parse the logic with the parser:
	relLogicsPath := def.LogicsPath()
	absLogicsPath := filepath.Join(app.currentPath, relLogicsPath)
	parsedProgram, err := prevParser.ExecuteFile(absLogicsPath)
	if err != nil {
		return nil, err
	}

	if castedParsedProgram, ok := parsedProgram.(parsers.Program); ok {
		program, err := app.Execute(castedParsedProgram)
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
		pathBuider := app.pathsBuilder.Create().
			WithBaseDir(app.currentPath).
			WithTokens(tokensPath).
			WithRules(rulesPath).
			WithLogics(relLogicsPath)

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
		return app.languageDefinitionBuilder.Create().
			WithApplication(langApp).
			WithPaths(paths).
			WithRoot(root).
			WithPatternMatches(patternMatches).
			Now()
	}

	str := fmt.Sprintf("the parsed program (relative path: %s) is invalid", absLogicsPath)
	return nil, errors.New(str)

}

func (app *linker) languageReference(
	def definitions.Definition,
	prevParser parsers.Parser,
) (LanguageReference, error) {
	langDef, err := app.languageDefinition(def, prevParser)
	if err != nil {
		return nil, err
	}

	input := def.InputVariable()
	return app.languageReferenceBuilder.Create().WithDefinition(langDef).WithInputVariable(input).Now()
}

func (app *linker) fileLanguageReference(
	relLangPath string,
	prevParser parsers.Parser,
) (LanguageReference, error) {
	// parse the language with the parser:
	absLangPath := filepath.Join(app.currentPath, relLangPath)

	// set the current path:
	app.currentPath = filepath.Dir(absLangPath)

	parsedProgram, err := prevParser.ExecuteFile(absLangPath)
	if err != nil {
		return nil, err
	}

	if castedProgram, ok := parsedProgram.(parsers.Program); ok {
		program, err := app.Execute(castedProgram)
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

	str := fmt.Sprintf("the parsed program (relative path: %s) is invalid", absLangPath)
	return nil, errors.New(str)
}

func (app *linker) buildGrammarRetrieverCriteria(
	def definitions.Definition,
	parser parsers.Parser,
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
