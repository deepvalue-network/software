package executables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/scripts"
)

type executable struct {
	application applications.Application
	script      scripts.Script
}

func createExecutableWithApplication(
	application applications.Application,
) Executable {
	return createExecutableInternally(application, nil)
}

func createExecutableWithScript(
	script scripts.Script,
) Executable {
	return createExecutableInternally(nil, script)
}

func createExecutableInternally(
	application applications.Application,
	script scripts.Script,
) Executable {
	out := executable{
		application: application,
		script:      script,
	}

	return &out
}

// IsApplication returns true if there is an application, false otherwise
func (obj *executable) IsApplication() bool {
	return obj.application != nil
}

// Application returns the application, if any
func (obj *executable) Application() applications.Application {
	return obj.application
}

// IsScript returns true if there is a script, false otherwise
func (obj *executable) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *executable) Script() scripts.Script {
	return obj.script
}
