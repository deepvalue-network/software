package hydro

type managerFactory struct {
}

func createManagerFactory() ManagerFactory {
	out := managerFactory{}
	return &out
}

// Create creates a manager instance
func (app *managerFactory) Create() Manager {
	return createManager(map[string]Bridge{})
}
