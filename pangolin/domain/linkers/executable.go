package linkers

type executable struct {
	app    Application
	script Script
}

func createExecutableWithApplication(
	app Application,
) Executable {
	return createExecutableInternally(app, nil)
}

func createExecutableWithScript(
	script Script,
) Executable {
	return createExecutableInternally(nil, script)
}

func createExecutableInternally(
	app Application,
	script Script,
) Executable {
	out := executable{
		app:    app,
		script: script,
	}

	return &out
}

// IsApplication returns true if the executable is an application, false otherwise
func (obj *executable) IsApplication() bool {
	return obj.app != nil
}

// Application returns the application, if any
func (obj *executable) Application() Application {
	return obj.app
}

// IsScript returns true if the executable is a script, false otherwise
func (obj *executable) IsScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *executable) Script() Script {
	return obj.script
}
