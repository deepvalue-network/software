package transform

import "errors"

type operationBuilder struct {
	miscBuilder MiscBuilder
	misc        Misc
	isPop       bool
}

func createOperationBuilder(miscBuilder MiscBuilder) OperationBuilder {
	out := operationBuilder{
		miscBuilder: miscBuilder,
		misc:        nil,
		isPop:       false,
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

// IsPop flags the builder as pop
func (app *operationBuilder) IsPop() OperationBuilder {
	app.isPop = true
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.misc != nil {
		return createOperationWithMisc(app.misc), nil
	}

	if app.isPop {
		misc, err := app.miscBuilder.Create().IsPop().Now()
		if err != nil {
			return nil, err
		}

		return createOperationWithMisc(misc), nil
	}

	return nil, errors.New("the Operation is invalid")
}
