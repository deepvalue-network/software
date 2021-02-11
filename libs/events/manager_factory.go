package events

type managerFactory struct {
}

func createManagerFactory() ManagerFactory {
	out := managerFactory{}
	return &out
}

// Create creates a new manager instance
func (app *managerFactory) Create() Manager {
	evts := []map[string]Event{}
	return createManager(evts)
}
