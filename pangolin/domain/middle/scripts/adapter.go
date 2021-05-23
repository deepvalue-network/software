package scripts

import (
	"path"

	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	builder      Builder
	testsBuilder TestsBuilder
	testBuilder  TestBuilder
}

func createAdapter(
	builder Builder,
	testsBuilder TestsBuilder,
	testBuilder TestBuilder,
) Adapter {
	out := adapter{
		builder:      builder,
		testsBuilder: testsBuilder,
		testBuilder:  testBuilder,
	}

	return &out
}

// ToScript converts a parsed script to a script
func (app *adapter) ToScript(parsed parsers.Script) (Script, error) {
	name := parsed.Name()
	version := parsed.Version()
	scriptPath := parsed.Script().String()
	languagePath := parsed.Language().String()
	output := parsed.Output()
	builder := app.builder.Create().
		WithName(name).
		WithVersion(version).
		WithLanguagePath(languagePath).
		WithScriptPath(scriptPath).
		WithOutput(output)

	if parsed.HasTests() {
		list := []Test{}
		parsedTests := parsed.Tests().All()
		for _, oneTest := range parsedTests {
			name := oneTest.Name()
			relPath := oneTest.Path().String()
			path := path.Join(scriptPath, relPath)
			test, err := app.testBuilder.Create().WithName(name).WithPath(path).Now()
			if err != nil {
				return nil, err
			}

			list = append(list, test)
		}

		tests, err := app.testsBuilder.Create().WithTests(list).Now()
		if err != nil {
			return nil, err
		}

		builder.WithTests(tests)
	}

	return builder.Now()
}
