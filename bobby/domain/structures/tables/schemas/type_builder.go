package schemas

import "errors"

type typeBuilder struct {
	isString  bool
	isInt     bool
	isFloat32 bool
	isFloat64 bool
	isData    bool
}

func createTypeBuilder() TypeBuilder {
	out := typeBuilder{
		isString:  false,
		isInt:     false,
		isFloat32: false,
		isFloat64: false,
		isData:    false,
	}

	return &out
}

// Create initializes the builder
func (app *typeBuilder) Create() TypeBuilder {
	return createTypeBuilder()
}

// IsString flags the builder as string
func (app *typeBuilder) IsString() TypeBuilder {
	app.isString = true
	return app
}

// IsInt flags the builder as int
func (app *typeBuilder) IsInt() TypeBuilder {
	app.isInt = true
	return app
}

// IsFloat32 flags the builder as float32
func (app *typeBuilder) IsFloat32() TypeBuilder {
	app.isFloat32 = true
	return app
}

// IsFloat64 flags the builder as float64
func (app *typeBuilder) IsFloat64() TypeBuilder {
	app.isFloat64 = true
	return app
}

// IsData flags the builder as data
func (app *typeBuilder) IsData() TypeBuilder {
	app.isData = true
	return app
}

// Now builds a new Type instance
func (app *typeBuilder) Now() (Type, error) {
	if app.isString {
		return createTypeWithString(), nil
	}

	if app.isInt {
		return createTypeWithInt(), nil
	}

	if app.isFloat32 {
		return createTypeWithFloat32(), nil
	}

	if app.isFloat64 {
		return createTypeWithFloat64(), nil
	}

	if app.isData {
		return createTypeWithData(), nil
	}

	return nil, errors.New("the Type is invalid")
}
