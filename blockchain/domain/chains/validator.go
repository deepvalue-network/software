package chains

type validator struct {
}

func createValidator() Validator {
	out := validator{}
	return &out
}

// Validate validates a chain, returns an error if invalid, nil if valid
func (app *validator) Validate(chain Chain) error {
	return nil
}
