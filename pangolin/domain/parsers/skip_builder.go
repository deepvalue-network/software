package parsers

import "errors"

type skipBuilder struct {
	pointer IntPointer
}

func createSkipBuilder() SkipBuilder {
	out := skipBuilder{
		pointer: nil,
	}

	return &out
}

// Create initializes the builder
func (app *skipBuilder) Create() SkipBuilder {
	return createSkipBuilder()
}

// WithPointer adds a pointer to the builder
func (app *skipBuilder) WithPointer(pointer IntPointer) SkipBuilder {
	app.pointer = pointer
	return app
}

// Now builds a new Skip instance
func (app *skipBuilder) Now() (Skip, error) {
	if app.pointer == nil {
		return nil, errors.New("the IntPointer is mandatory in order to build a Skip instance")
	}

	return createSkip(app.pointer), nil
}
