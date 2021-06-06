package linkers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/scripts"
	language_applications "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type linker struct {
	middleAdapter                   middle.Adapter
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder
	programBuilder                  ProgramBuilder
	testableBuilder                 TestableBuilder
	executableBuilder               ExecutableBuilder
	applicationBuilder              ApplicationBuilder
	languageDefinitionBuilder       LanguageDefinitionBuilder
	pathsBuilder                    PathsBuilder
	scriptBuilder                   ScriptBuilder
	testBuilder                     TestBuilder
	languageReferenceBuilder        LanguageReferenceBuilder
	languageApplicationBuilder      LanguageApplicationBuilder
	parser                          parsers.Parser
	dirPath                         string
	currentPath                     string
}

func createLinker(
	middleAdapter middle.Adapter,
	grammarRetrieverCriteriaBuilder grammar.RetrieverCriteriaBuilder,
	programBuilder ProgramBuilder,
	testableBuilder TestableBuilder,
	executableBuilder ExecutableBuilder,
	applicationBuilder ApplicationBuilder,
	languageDefinitionBuilder LanguageDefinitionBuilder,
	pathsBuilder PathsBuilder,
	scriptBuilder ScriptBuilder,
	testBuilder TestBuilder,
	languageReferenceBuilder LanguageReferenceBuilder,
	languageApplicationBuilder LanguageApplicationBuilder,
	parser parsers.Parser,
	dirPath string,
) Linker {
	out := linker{
		middleAdapter:                   middleAdapter,
		grammarRetrieverCriteriaBuilder: grammarRetrieverCriteriaBuilder,
		programBuilder:                  programBuilder,
		testableBuilder:                 testableBuilder,
		executableBuilder:               executableBuilder,
		applicationBuilder:              applicationBuilder,
		languageDefinitionBuilder:       languageDefinitionBuilder,
		pathsBuilder:                    pathsBuilder,
		scriptBuilder:                   scriptBuilder,
		testBuilder:                     testBuilder,
		languageReferenceBuilder:        languageReferenceBuilder,
		languageApplicationBuilder:      languageApplicationBuilder,
		parser:                          parser,
		dirPath:                         dirPath,
		currentPath:                     dirPath,
	}

	return &out
}

// Execute links a parsed program into a linked program
func (app *linker) Execute(parsed parsers.Program) (Program, error) {
	return app.parsedProgram(parsed)
}

func (app *linker) parsedProgram(parsed parsers.Program) (Program, error) {
	middleProgram, err := app.middleAdapter.ToProgram(parsed)
	if err != nil {
		return nil, err
	}

	return app.program(middleProgram)
}

func (app *linker) program(
	program middle.Program,
) (Program, error) {
	builder := app.programBuilder.Create()
	if program.IsTestable() {
		testable := program.Testable()
		linkedTestable, err := app.testable(testable)
		if err != nil {
			return nil, err
		}

		builder.WithTestable(linkedTestable)
	}

	if program.IsLanguage() {
		langApp := program.Language()
		linkedLangAppli, err := app.languageApplication(langApp)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(linkedLangAppli)
	}

	return builder.Now()
}

func (app *linker) testable(
	testable testables.Testable,
) (Testable, error) {
	builder := app.testableBuilder.Create()
	if testable.IsExecutable() {
		executable := testable.Executable()
		linkedExecutable, err := app.executable(executable)
		if err != nil {
			return nil, err
		}

		builder.WithExecutable(linkedExecutable)
	}

	if testable.IsLanguage() {
		langDef := testable.Language()
		linkedLangRef, err := app.languageReference(langDef)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(linkedLangRef)
	}

	return builder.Now()
}

func (app *linker) executable(
	executable executables.Executable,
) (Executable, error) {
	builder := app.executableBuilder.Create()
	if executable.IsApplication() {
		appli := executable.Application()
		linkedApp, err := app.application(appli)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(linkedApp)
	}

	if executable.IsScript() {
		script := executable.Script()
		linkedScript, err := app.script(script)
		if err != nil {
			return nil, err
		}

		builder.WithScript(linkedScript)
	}

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
) (Script, error) {
	relLangPath := script.LanguagePath()
	langRef, err := app.fileLanguageReference(relLangPath)
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
	output := script.Output()
	builder := app.scriptBuilder.Create().
		WithLanguage(langRef).
		WithName(name).
		WithVersion(version).
		WithCode(string(content)).
		WithOutput(output)

	if script.HasTests() {
		tests := []Test{}
		parsedTests := script.Tests().All()
		for _, oneParsedTest := range parsedTests {
			parsedPath := oneParsedTest.Path()
			parsedProg, err := app.parser.ExecuteFile(parsedPath)
			if err != nil {
				return nil, err
			}

			if castedParsedProg, ok := parsedProg.(parsers.Program); ok {
				prog, err := app.Execute(castedParsedProg)
				if err != nil {
					return nil, err
				}

				if !prog.IsTestable() {
					return nil, errors.New("the test was expected to contain a testable application")
				}

				testable := prog.Testable()
				if !testable.IsExecutable() {
					return nil, errors.New("the test was expected to contain an executable application")
				}

				name := oneParsedTest.Name()
				executable := testable.Executable()
				test, err := app.testBuilder.Create().WithName(name).WithExecutable(executable).Now()
				if err != nil {
					return nil, err
				}

				tests = append(tests, test)
				continue
			}

			str := fmt.Sprintf("the test script (path: %s) is not a valid Program", parsedPath)
			return nil, errors.New(str)
		}

		builder.WithTests(tests)
	}

	return builder.Now()
}

func (app *linker) languageApplication(
	langApp language_applications.Application,
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
) (LanguageDefinition, error) {
	// parse the logic with the parser:
	relLogicsPath := def.LogicsPath()
	absLogicsPath := filepath.Join(app.currentPath, relLogicsPath)
	parsedProgram, err := app.parser.ExecuteFile(absLogicsPath)
	if err != nil {
		return nil, err
	}

	if castedParsedProgram, ok := parsedProgram.(parsers.Program); ok {
		program, err := app.parsedProgram(castedParsedProgram)
		if err != nil {
			return nil, err
		}

		strErr := errors.New("the language definition was expected to contain a language application")
		if !program.IsLanguage() {
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
		langApp := program.Language()
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
) (LanguageReference, error) {
	langDef, err := app.languageDefinition(def)
	if err != nil {
		return nil, err
	}

	input := def.InputVariable()
	return app.languageReferenceBuilder.Create().WithDefinition(langDef).WithInputVariable(input).Now()
}

func (app *linker) fileLanguageReference(
	relLangPath string,
) (LanguageReference, error) {
	// parse the language with the parser:
	absLangPath := filepath.Join(app.currentPath, relLangPath)

	// set the current path:
	app.currentPath = filepath.Dir(absLangPath)

	parsedProgram, err := app.parser.ExecuteFile(absLangPath)
	if err != nil {
		return nil, err
	}

	if castedProgram, ok := parsedProgram.(parsers.Program); ok {
		program, err := app.parsedProgram(castedProgram)
		if err != nil {
			return nil, err
		}

		if !program.IsTestable() {
			str := fmt.Sprintf("the language file (%s) was expected to contain a testable program", absLangPath)
			return nil, errors.New(str)
		}

		testable := program.Testable()
		if !testable.IsLanguage() {
			str := fmt.Sprintf("the language file (%s) was expected to contain a testable language program", absLangPath)
			return nil, errors.New(str)
		}

		return testable.Language(), nil
	}

	str := fmt.Sprintf("the parsed program (relative path: %s) is invalid", absLangPath)
	return nil, errors.New(str)
}

func (app *linker) buildGrammarRetrieverCriteria(
	def definitions.Definition,
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
