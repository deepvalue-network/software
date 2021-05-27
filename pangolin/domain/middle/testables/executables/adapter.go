package executables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	applicationAdapter applications.Adapter
	scriptAdapter      scripts.Adapter
	builder            Builder
}

func createAdapter(
	applicationAdapter applications.Adapter,
	scriptAdapter scripts.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		applicationAdapter: applicationAdapter,
		scriptAdapter:      scriptAdapter,
		builder:            builder,
	}

	return &out
}

// ToExecutable converts a parsed executable to an executable
func (app *adapter) ToExecutable(parsed parsers.Executable) (Executable, error) {
	builder := app.builder.Create()
	if parsed.IsApplication() {
		parsedApp := parsed.Application()
		app, err := app.applicationAdapter.ToApplication(parsedApp)
		if err != nil {
			return nil, err
		}

		builder.WithApplication(app)
	}

	if parsed.IsScript() {
		parsedScript := parsed.Script()
		script, err := app.scriptAdapter.ToScript(parsedScript)
		if err != nil {
			return nil, err
		}

		builder.WithScript(script)
	}

	return builder.Now()
}
