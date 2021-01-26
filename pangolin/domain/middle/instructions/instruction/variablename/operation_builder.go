package variablename

import "errors"

type operationBuilder struct {
	miscBuilder MiscBuilder
	misc        Misc
	isPush      bool
}

func createOperationBuilder(miscBuilder MiscBuilder) OperationBuilder {
	out := operationBuilder{
		miscBuilder: miscBuilder,
		misc:        nil,
		isPush:      false,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder(app.miscBuilder)
}

// WithMisc adds a misc to the builder
func (app *operationBuilder) WithMisc(misc Misc) OperationBuilder {
	app.misc = misc
	return app
}

// IsPush flags the builder as pop
func (app *operationBuilder) IsPush() OperationBuilder {
	app.isPush = true
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.misc != nil {
		return createOperationWithMisc(app.misc), nil
	}

	if app.isPush {
		misc, err := app.miscBuilder.Create().IsPush().Now()
		if err != nil {
			return nil, err
		}

		return createOperationWithMisc(misc), nil
	}

	return nil, errors.New("the Operation is invalid")
}
