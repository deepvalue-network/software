package parsers

import "errors"

type typeBuilder struct {
	isBool    bool
	isInt8    bool
	isInt16   bool
	isInt32   bool
	isInt64   bool
	isFloat32 bool
	isFloat64 bool
	isUint8   bool
	isUint16  bool
	isUint32  bool
	isUint64  bool
	isString  bool
}

func createTypeBuilder() TypeBuilder {
	out := typeBuilder{
		isBool:    false,
		isInt8:    false,
		isInt16:   false,
		isInt32:   false,
		isInt64:   false,
		isFloat32: false,
		isFloat64: false,
		isUint8:   false,
		isUint16:  false,
		isUint32:  false,
		isUint64:  false,
		isString:  false,
	}

	return &out
}

// Create initializes the builder
func (app *typeBuilder) Create() TypeBuilder {
	return createTypeBuilder()
}

// IsBool flags the type as bool
func (app *typeBuilder) IsBool() TypeBuilder {
	app.isBool = true
	return app
}

// IsInt8 flags the type as int8
func (app *typeBuilder) IsInt8() TypeBuilder {
	app.isInt8 = true
	return app
}

// IsInt16 flags the type as int16
func (app *typeBuilder) IsInt16() TypeBuilder {
	app.isInt16 = true
	return app
}

// IsInt32 flags the type as int32
func (app *typeBuilder) IsInt32() TypeBuilder {
	app.isInt32 = true
	return app
}

// IsInt64 flags the type as int64
func (app *typeBuilder) IsInt64() TypeBuilder {
	app.isInt64 = true
	return app
}

// IsFloat32 flags the type as float32
func (app *typeBuilder) IsFloat32() TypeBuilder {
	app.isFloat32 = true
	return app
}

// IsFloat64 flags the type as float64
func (app *typeBuilder) IsFloat64() TypeBuilder {
	app.isFloat64 = true
	return app
}

// IsUint8 flags the type as uint8
func (app *typeBuilder) IsUint8() TypeBuilder {
	app.isUint8 = true
	return app
}

// IsUint16 flags the type as uint16
func (app *typeBuilder) IsUint16() TypeBuilder {
	app.isUint16 = true
	return app
}

// IsUint32 flags the type as uint32
func (app *typeBuilder) IsUint32() TypeBuilder {
	app.isUint32 = true
	return app
}

// IsUint64 flags the type as uint64
func (app *typeBuilder) IsUint64() TypeBuilder {
	app.isUint64 = true
	return app
}

// IsString flags the type as char
func (app *typeBuilder) IsString() TypeBuilder {
	app.isString = true
	return app
}

// Now builds a new Type instance
func (app *typeBuilder) Now() (Type, error) {
	if app.isBool {
		return createTypeWithBool(), nil
	}

	if app.isInt8 {
		return createTypeWithInt8(), nil
	}

	if app.isInt16 {
		return createTypeWithInt16(), nil
	}

	if app.isInt32 {
		return createTypeWithInt32(), nil
	}

	if app.isInt64 {
		return createTypeWithInt64(), nil
	}

	if app.isFloat32 {
		return createTypeWithFloat32(), nil
	}

	if app.isFloat64 {
		return createTypeWithFloat64(), nil
	}

	if app.isUint8 {
		return createTypeWithUint8(), nil
	}

	if app.isUint16 {
		return createTypeWithUint16(), nil
	}

	if app.isUint32 {
		return createTypeWithUint32(), nil
	}

	if app.isUint64 {
		return createTypeWithUint64(), nil
	}

	if app.isString {
		return createTypeWithString(), nil
	}

	return nil, errors.New("the Type is invalid")
}
