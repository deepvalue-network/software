package condition

import "errors"

type operationBuilder struct {
	isJump bool
}

func createOperationBuilder() OperationBuilder {
	out := operationBuilder{
		isJump: false,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// IsJump flags the builder as a jump
func (app *operationBuilder) IsJump() OperationBuilder {
	app.isJump = true
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.isJump {
		return createOperationWithJump(), nil
	}

	return nil, errors.New("the Operation is invalid")
}
