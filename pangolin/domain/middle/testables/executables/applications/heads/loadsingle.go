package heads

type loadSingle struct {
	internal string
	external string
}

func createLoadSingle(
	internal string,
	external string,
) LoadSingle {
	out := loadSingle{
		internal: internal,
		external: external,
	}

	return &out
}

// Internal returns the internal
func (obj *loadSingle) Internal() string {
	return obj.internal
}

// External returns the external
func (obj *loadSingle) External() string {
	return obj.external
}
