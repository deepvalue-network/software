package computable

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
)

type builder struct {
	bl             *bool
	str            *string
	intHeight      *int8
	intSixteen     *int16
	intThirtyTwo   *int32
	intSixtyFour   *int64
	floatThirtyTwo *float32
	floatSixtyFour *float64
	uintHeight     *uint8
	uintSixteen    *uint16
	uintThirtyTwo  *uint32
	uintSixtyFour  *uint64
	isToken        bool
	token          lexers.NodeTree
	isStackFrame   bool
	isFrame        bool
}

func createBuilder() Builder {
	out := builder{
		bl:             nil,
		str:            nil,
		intHeight:      nil,
		intSixteen:     nil,
		intThirtyTwo:   nil,
		intSixtyFour:   nil,
		floatThirtyTwo: nil,
		floatSixtyFour: nil,
		uintHeight:     nil,
		uintSixteen:    nil,
		uintThirtyTwo:  nil,
		uintSixtyFour:  nil,
		isToken:        false,
		token:          nil,
		isStackFrame:   false,
		isFrame:        false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBool adds a bool to the builder
func (app *builder) WithBool(bl bool) Builder {
	app.bl = &bl
	return app
}

// WithInt8 adds an int8 to the builder
func (app *builder) WithInt8(intHeight int8) Builder {
	app.intHeight = &intHeight
	return app
}

// WithInt16 adds an int16 to the builder
func (app *builder) WithInt16(intSixteen int16) Builder {
	app.intSixteen = &intSixteen
	return app
}

// WithInt32 adds an int32 to the builder
func (app *builder) WithInt32(intThirtyTwo int32) Builder {
	app.intThirtyTwo = &intThirtyTwo
	return app
}

// WithInt64 adds an int64 to the builder
func (app *builder) WithInt64(intSixtyFour int64) Builder {
	app.intSixtyFour = &intSixtyFour
	return app
}

// WithUint8 adds an uint8 to the builder
func (app *builder) WithUint8(uintHeight uint8) Builder {
	app.uintHeight = &uintHeight
	return app
}

// WithUint16 adds an uint16 to the builder
func (app *builder) WithUint16(uintSixteen uint16) Builder {
	app.uintSixteen = &uintSixteen
	return app
}

// WithUint32 adds an uint32 to the builder
func (app *builder) WithUint32(uintThirtyTwo uint32) Builder {
	app.uintThirtyTwo = &uintThirtyTwo
	return app
}

// WithUint64 adds an uint64 to the builder
func (app *builder) WithUint64(uintSixtyFour uint64) Builder {
	app.uintSixtyFour = &uintSixtyFour
	return app
}

// WithFloat32 adds a float32 to the builder
func (app *builder) WithFloat32(floatThirtyTwo float32) Builder {
	app.floatThirtyTwo = &floatThirtyTwo
	return app
}

// WithFloat64 adds a float64 to the builder
func (app *builder) WithFloat64(floatSixtyFour float64) Builder {
	app.floatSixtyFour = &floatSixtyFour
	return app
}

// WithString adds a string to the builder
func (app *builder) WithString(str string) Builder {
	app.str = &str
	return app
}

// IsToken flags the builder as a token
func (app *builder) IsToken() Builder {
	app.isToken = true
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(tok lexers.NodeTree) Builder {
	app.IsToken()
	app.token = tok
	return app
}

// IsStackFrame adds a stackFrame to the builder
func (app *builder) IsStackFrame() Builder {
	app.isStackFrame = true
	return app
}

// IsFrame adds a frame to the builder
func (app *builder) IsFrame() Builder {
	app.isFrame = true
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.intHeight != nil {
		return createValueWithInt8(app.intHeight), nil
	}

	if app.intSixteen != nil {
		return createValueWithInt16(app.intSixteen), nil
	}

	if app.intThirtyTwo != nil {
		return createValueWithInt32(app.intThirtyTwo), nil
	}

	if app.intSixtyFour != nil {
		return createValueWithInt64(app.intSixtyFour), nil
	}

	if app.uintHeight != nil {
		return createValueWithUint8(app.uintHeight), nil
	}

	if app.uintSixteen != nil {
		return createValueWithUint16(app.uintSixteen), nil
	}

	if app.uintThirtyTwo != nil {
		return createValueWithUint32(app.uintThirtyTwo), nil
	}

	if app.uintSixtyFour != nil {
		return createValueWithUint64(app.uintSixtyFour), nil
	}

	if app.floatThirtyTwo != nil {
		return createValueWithFloat32(app.floatThirtyTwo), nil
	}

	if app.floatSixtyFour != nil {
		return createValueWithFloat64(app.floatSixtyFour), nil
	}

	if app.bl != nil {
		return createValueWithBool(app.bl), nil
	}

	if app.str != nil {
		return createValueWithString(app.str), nil
	}

	if app.isToken {
		if app.token != nil {
			return createValueWithToken(app.token), nil
		}

		return createValueWithNilToken(), nil
	}

	if app.isStackFrame {
		return createValueWithStackFrame(), nil
	}

	if app.isFrame {
		return createValueWithFrame(), nil
	}

	panic(errors.New("voila"))

	return nil, errors.New("the computable Value is invalid")
}
