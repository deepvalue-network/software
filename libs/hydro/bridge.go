package hydro

type bridge struct {
	hydrated   Hydrated
	dehydrated Dehydrated
}

func createBridge(
	hydrated Hydrated,
	dehydrated Dehydrated,
) Bridge {
	out := bridge{
		hydrated:   hydrated,
		dehydrated: dehydrated,
	}

	return &out
}

// Hydrated returns the hydrated instance
func (obj *bridge) Hydrated() Hydrated {
	return obj.hydrated
}

// Dehydrated returns the dehydrated instance
func (obj *bridge) Dehydrated() Dehydrated {
	return obj.dehydrated
}
