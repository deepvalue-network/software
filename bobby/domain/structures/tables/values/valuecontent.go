package values

import uuid "github.com/satori/go.uuid"

type valueContent struct {
	id         *uuid.UUID
	stringVal  *string
	intVal     *int
	float32Val *float32
	float64Val *float64
	data       []byte
}

func createValueContentWithID(id *uuid.UUID) ValueContent {
	return createValueContentInternally(id, nil, nil, nil, nil, nil)
}

func createValueContentWithString(stringVal *string) ValueContent {
	return createValueContentInternally(nil, stringVal, nil, nil, nil, nil)
}

func createValueContentWithInt(intVal *int) ValueContent {
	return createValueContentInternally(nil, nil, intVal, nil, nil, nil)
}

func createValueContentWithFloat32(float32Val *float32) ValueContent {
	return createValueContentInternally(nil, nil, nil, float32Val, nil, nil)
}

func createValueContentWithFloat64(float64Val *float64) ValueContent {
	return createValueContentInternally(nil, nil, nil, nil, float64Val, nil)
}

func createValueContentWithData(data []byte) ValueContent {
	return createValueContentInternally(nil, nil, nil, nil, nil, data)
}

func createValueContentInternally(
	id *uuid.UUID,
	stringVal *string,
	intVal *int,
	float32Val *float32,
	float64Val *float64,
	data []byte,
) ValueContent {
	out := valueContent{
		id:         id,
		stringVal:  stringVal,
		intVal:     intVal,
		float32Val: float32Val,
		float64Val: float64Val,
		data:       data,
	}

	return &out
}

// IsID returns true if there is an ID, false otherwise
func (obj *valueContent) IsID() bool {
	return obj.id != nil
}

// ID returns the ID, if any
func (obj *valueContent) ID() *uuid.UUID {
	return obj.id
}

// IsString returns true if there is a string, false otherwise
func (obj *valueContent) IsString() bool {
	return obj.stringVal != nil
}

// String returns the string, if any
func (obj *valueContent) String() *string {
	return obj.stringVal
}

// IsInt returns true if there is an int, false otherwise
func (obj *valueContent) IsInt() bool {
	return obj.intVal != nil
}

// Int returns the int, if any
func (obj *valueContent) Int() *int {
	return obj.intVal
}

// IsFloat32 returns true if there is a float32, false otherwise
func (obj *valueContent) IsFloat32() bool {
	return obj.float32Val != nil
}

// Float32 returns the float32, if any
func (obj *valueContent) Float32() *float32 {
	return obj.float32Val
}

// IsFloat64 returns true if there is a float64, false otherwise
func (obj *valueContent) IsFloat64() bool {
	return obj.float64Val != nil
}

// Float64 returns the float64, if any
func (obj *valueContent) Float64() *float64 {
	return obj.float64Val
}

// IsData returns true if there is data, false otherwise
func (obj *valueContent) IsData() bool {
	return obj.data != nil
}

// Data returns the float64, if any
func (obj *valueContent) Data() []byte {
	return obj.data
}
