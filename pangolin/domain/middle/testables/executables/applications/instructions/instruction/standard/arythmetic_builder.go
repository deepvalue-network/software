package standard

import "errors"

type arythmeticBuilder struct {
	isAdd bool
	isSub bool
	isMul bool
}

func createArythmericBuilder() ArythmeticBuilder {
	out := arythmeticBuilder{
		isAdd: false,
		isSub: false,
		isMul: false,
	}

	return &out
}

// Create intializes the builder
func (app *arythmeticBuilder) Create() ArythmeticBuilder {
	return createArythmericBuilder()
}

// IsAdd flags the builder as an add
func (app *arythmeticBuilder) IsAdd() ArythmeticBuilder {
	app.isAdd = true
	return app
}

// IsSub flags the builder as a sub
func (app *arythmeticBuilder) IsSub() ArythmeticBuilder {
	app.isSub = true
	return app
}

// IsMul flags the builder as a mul
func (app *arythmeticBuilder) IsMul() ArythmeticBuilder {
	app.isMul = true
	return app
}

// Now builds a new Arythmetic instance
func (app *arythmeticBuilder) Now() (Arythmetic, error) {
	if app.isAdd {
		return createArythmeticWithAdd(), nil
	}

	if app.isSub {
		return createArythmeticWithSub(), nil
	}

	if app.isMul {
		return createArythmeticWithMul(), nil
	}

	return nil, errors.New("the Arythmetic is invalid")
}
