package linkers

type external struct {
	name        string
	application Application
	script      Script
}

func createExternalWithApplication(
	name string,
	application Application,
) External {
	return createExternalInternally(name, application, nil)
}

func createExternalWithScript(
	name string,
	script Script,
) External {
	return createExternalInternally(name, nil, script)
}

func createExternalInternally(
	name string,
	application Application,
	script Script,
) External {
	out := external{
		name:        name,
		application: application,
		script:      script,
	}

	return &out
}

// Name returns the name
func (obj *external) Name() string {
	return obj.name
}

// HasApplication returns true if there is an application, false otherwise
func (obj *external) HasApplication() bool {
	return obj.application != nil
}

// Application returns the application
func (obj *external) Application() Application {
	return obj.application
}

// HasScript returns true if there is a script, false otherwise
func (obj *external) HasScript() bool {
	return obj.script != nil
}

// Script returns the script, if any
func (obj *external) Script() Script {
	return obj.script
}
