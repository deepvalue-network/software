package parsers

type executable struct {
	application Application
	script      Script
}

func createExecutableWithApplication(
	application Application,
) Executable {
	return createExecutableInternally(application, nil)
}

func createExecutableWithScript(
	script Script,
) Executable {
	return createExecutableInternally(nil, script)
}

func createExecutableInternally(
	application Application,
	script Script,
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
func (obj *executable) Application() Application {
	return obj.application
}

// IsScript returns true if there is a script, false otherwise
func (obj *executable) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *executable) Script() Script {
	return obj.script
}
