package scripts

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
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
	return app.builder.Create().
		WithName(name).
		WithVersion(version).
		WithLanguagePath(languagePath).
		WithScriptPath(scriptPath).
		WithOutput(output).
		Now()
}
