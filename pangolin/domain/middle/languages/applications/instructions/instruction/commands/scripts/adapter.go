package scripts

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type adapter struct {
	builder      Builder
	valueBuilder ValueBuilder
}

func createAdapter(
	builder Builder,
	valueBuilder ValueBuilder,
) Adapter {
	out := adapter{
		builder:      builder,
		valueBuilder: valueBuilder,
	}

	return &out
}

// ToScript converts a parsed script command to a script instance
func (app *adapter) ToScript(parsed parsers.ScriptCommand) (Script, error) {
	values := []Value{}
	parsedValues := parsed.Values()
	for _, oneValue := range parsedValues {
		valueBuilder := app.valueBuilder.Create()
		if oneValue.IsName() {
			name := oneValue.Name()
			valueBuilder.WithName(name)
		}

		if oneValue.IsVersion() {
			version := oneValue.Version()
			valueBuilder.WithVersion(version)
		}

		if oneValue.IsScript() {
			scriptPath := oneValue.Script().String()
			valueBuilder.WithScriptPath(scriptPath)
		}

		if oneValue.IsLanguage() {
			langPath := oneValue.Language().String()
			valueBuilder.WithLanguagePath(langPath)
		}

		value, err := valueBuilder.Now()
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	variable := parsed.Variable()
	return app.builder.Create().
		WithValues(values).
		WithVariable(variable).
		Now()
}
