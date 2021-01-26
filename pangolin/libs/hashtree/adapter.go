package hashtree

type adapter struct {
}

func createAdapter() Adapter {
	out := adapter{}
	return &out
}

// FromJSON converts a jsonCompact to Compact
func (app *adapter) FromJSON(js *JSONCompact) (Compact, error) {
	return createCompactFromJSON(js)
}

// ToJSON converts a Compact to jsonCompact
func (app *adapter) ToJSON(ht Compact) *JSONCompact {
	return createJSONCompactFromCompact(ht)
}
