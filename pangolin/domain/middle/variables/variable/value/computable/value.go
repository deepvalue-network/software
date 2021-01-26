package computable

import (
	"strconv"

	"github.com/steve-care-software/products/pangolin/domain/lexers"
)

type value struct {
	isNil          bool
	bl             *bool
	str            *string
	intHeight      *int8
	intSixteen     *int16
	intThirtyTwo   *int32
	intSixtyFour   *int64
	uintHeight     *uint8
	uintSixteen    *uint16
	uintThirtyTwo  *uint32
	uintSixtyFour  *uint64
	floatThirtyTwo *float32
	floatSixtyFour *float64
	isToken        bool
	token          lexers.NodeTree
	isStackFrame   bool
	isFrame        bool
}

func createValueWithNil() Value {
	return createValueInternally(
		true,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithBool(bl *bool) Value {
	return createValueInternally(
		false,
		bl,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithString(str *string) Value {
	return createValueInternally(
		false,
		nil,
		str,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithInt8(intHeight *int8) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		intHeight,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithInt16(intSixteen *int16) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		intSixteen,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithInt32(intThirtyTwo *int32) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		intThirtyTwo,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithInt64(intSixtyFour *int64) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		intSixtyFour,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithUint8(uintHeight *uint8) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		uintHeight,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithUint16(uintSixteen *uint16) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		uintSixteen,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithUint32(uintThirtyTwo *uint32) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		uintThirtyTwo,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithUint64(uintSixtyFour *uint64) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		uintSixtyFour,
		nil,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithFloat32(floatThirtyTwo *float32) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		floatThirtyTwo,
		nil,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithFloat64(floatSixtyFour *float64) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		floatSixtyFour,
		false,
		nil,
		false,
		false,
	)
}

func createValueWithNilToken() Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		true,
		nil,
		false,
		false,
	)
}

func createValueWithToken(token lexers.NodeTree) Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		true,
		token,
		false,
		false,
	)
}

func createValueWithStackFrame() Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		true,
		false,
	)
}

func createValueWithFrame() Value {
	return createValueInternally(
		false,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		false,
		nil,
		false,
		true,
	)
}

func createValueInternally(
	isNil bool,
	bl *bool,
	str *string,
	intHeight *int8,
	intSixteen *int16,
	intThirtyTwo *int32,
	intSixtyFour *int64,
	uintHeight *uint8,
	uintSixteen *uint16,
	uintThirtyTwo *uint32,
	uintSixtyFour *uint64,
	floatThirtyTwo *float32,
	floatSixtyFour *float64,
	isToken bool,
	token lexers.NodeTree,
	isStackFrame bool,
	isFrame bool,
) Value {
	out := value{
		isNil:          isNil,
		bl:             bl,
		str:            str,
		intHeight:      intHeight,
		intSixteen:     intSixteen,
		intThirtyTwo:   intThirtyTwo,
		intSixtyFour:   intSixtyFour,
		uintHeight:     uintHeight,
		uintSixteen:    uintSixteen,
		uintThirtyTwo:  uintThirtyTwo,
		uintSixtyFour:  uintSixtyFour,
		floatThirtyTwo: floatThirtyTwo,
		floatSixtyFour: floatSixtyFour,
		isToken:        isToken,
		token:          token,
		isStackFrame:   isStackFrame,
		isFrame:        isFrame,
	}

	return &out
}

// IsNil returns true if nil, false otherwise
func (obj *value) IsNil() bool {
	return obj.isNil
}

// IsBool returns true if the value is a boolean, false otherwise
func (obj *value) IsBool() bool {
	return obj.bl != nil
}

// Bool returns the bool, if any
func (obj *value) Bool() *bool {
	return obj.bl
}

// IsString returns true if the value is a string, false otherwise
func (obj *value) IsString() bool {
	return obj.str != nil
}

// String returns the string, if any
func (obj *value) String() *string {
	return obj.str
}

// IsIntHeight returns true if the value is an intHeight, false otherwise
func (obj *value) IsIntHeight() bool {
	return obj.intHeight != nil
}

// IntHeight returns the int8, if any
func (obj *value) IntHeight() *int8 {
	return obj.intHeight
}

// IsIntSixteen returns true if the value is an intSixteen, false otherwise
func (obj *value) IsIntSixteen() bool {
	return obj.intSixteen != nil
}

// IntSixteen returns the int16, if any
func (obj *value) IntSixteen() *int16 {
	return obj.intSixteen
}

// IsIntThirtyTwo returns true if the value is an intThirtyTwo, false otherwise
func (obj *value) IsIntThirtyTwo() bool {
	return obj.intThirtyTwo != nil
}

// IntThirtyTwo returns the int32, if any
func (obj *value) IntThirtyTwo() *int32 {
	return obj.intThirtyTwo
}

// IsIntSixtyFour returns true if the value is an intSixtyFour, false otherwise
func (obj *value) IsIntSixtyFour() bool {
	return obj.intSixtyFour != nil
}

// IntSixtyFour returns the int64, if any
func (obj *value) IntSixtyFour() *int64 {
	return obj.intSixtyFour
}

// IsUintHeight returns true if the value is an uintHeight, false otherwise
func (obj *value) IsUintHeight() bool {
	return obj.uintHeight != nil
}

// UintHeight returns the uint8, if any
func (obj *value) UintHeight() *uint8 {
	return obj.uintHeight
}

// IsUintSixteen returns true if the value is an uintSixteen, false otherwise
func (obj *value) IsUintSixteen() bool {
	return obj.uintSixteen != nil
}

// UintSixteen returns the uint16, if any
func (obj *value) UintSixteen() *uint16 {
	return obj.uintSixteen
}

// IsUintThirtyTwo returns true if the value is an uintThirtyTwo, false otherwise
func (obj *value) IsUintThirtyTwo() bool {
	return obj.uintThirtyTwo != nil
}

// UintThirtyTwo returns the uint32, if any
func (obj *value) UintThirtyTwo() *uint32 {
	return obj.uintThirtyTwo
}

// IsUintSixtyFour returns true if the value is an uintSixtyFour, false otherwise
func (obj *value) IsUintSixtyFour() bool {
	return obj.uintSixtyFour != nil
}

// UintSixtyFour returns the uint64, if any
func (obj *value) UintSixtyFour() *uint64 {
	return obj.uintSixtyFour
}

// IsFloatThirtyTwo returns true if the value is a floatThirtyTwo, false otherwise
func (obj *value) IsFloatThirtyTwo() bool {
	return obj.floatThirtyTwo != nil
}

// FloatThirtyTwo returns the float32, if any
func (obj *value) FloatThirtyTwo() *float32 {
	return obj.floatThirtyTwo
}

// IsFloatSixtyFour returns true if the value is a floatSixtyFour, false otherwise
func (obj *value) IsFloatSixtyFour() bool {
	return obj.floatSixtyFour != nil
}

// FloatSixtyFour returns the float64, if any
func (obj *value) FloatSixtyFour() *float64 {
	return obj.floatSixtyFour
}

// IsToken returns true if the value is a token, false otherwise
func (obj *value) IsToken() bool {
	return obj.isToken
}

// Token returns the token, if any
func (obj *value) Token() lexers.NodeTree {
	return obj.token
}

// IsStackFrame returns true if the value is a stackframe, false otherwise
func (obj *value) IsStackFrame() bool {
	return obj.isStackFrame
}

// IsFrame returns true if the value is a frame, false otherwise
func (obj *value) IsFrame() bool {
	return obj.isFrame
}

// StringRepresentation returns the string representation
func (obj *value) StringRepresentation() string {
	if obj.IsNil() {
		return "nil"
	}

	if obj.IsBool() {
		val := obj.Bool()
		return strconv.FormatBool(*val)
	}

	if obj.IsString() {
		val := obj.String()
		return *val
	}

	if obj.IsIntHeight() {
		val := obj.IntHeight()
		return strconv.Itoa(int(*val))
	}

	if obj.IsIntSixteen() {
		val := obj.IntSixteen()
		return strconv.Itoa(int(*val))
	}

	if obj.IsIntThirtyTwo() {
		val := obj.IntThirtyTwo()
		return strconv.Itoa(int(*val))
	}

	if obj.IsIntSixtyFour() {
		val := obj.IntSixtyFour()
		return strconv.Itoa(int(*val))
	}

	if obj.IsUintHeight() {
		val := obj.UintHeight()
		return strconv.FormatUint(uint64(*val), 8)
	}

	if obj.IsUintSixteen() {
		val := obj.UintSixteen()
		return strconv.FormatUint(uint64(*val), 16)
	}

	if obj.IsUintThirtyTwo() {
		val := obj.UintThirtyTwo()
		return strconv.FormatUint(uint64(*val), 32)
	}

	if obj.IsUintSixtyFour() {
		val := obj.UintSixtyFour()
		return strconv.FormatUint(uint64(*val), 64)
	}

	if obj.IsFloatThirtyTwo() {
		val := obj.FloatThirtyTwo()
		return strconv.FormatFloat(float64(*val), 'f', -1, 32)
	}

	if obj.IsFloatSixtyFour() {
		val := obj.FloatSixtyFour()
		return strconv.FormatFloat(*val, 'f', -1, 64)
	}

	if obj.IsToken() {
		if obj.Token() != nil {
			return obj.Token().Code()
		}

		return "nil"
	}

	if obj.IsStackFrame() {
		return "stackframe"
	}

	return "frame"
}
