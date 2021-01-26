package computable

import (
	"github.com/steve-care-software/products/pangolin/domain/lexers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a computable value builder
type Builder interface {
	Create() Builder
	WithBool(bl bool) Builder
	WithString(str string) Builder
	WithInt8(intHeight int8) Builder
	WithInt16(intSixteen int16) Builder
	WithInt32(intThirtyTwo int32) Builder
	WithInt64(intSixtyFour int64) Builder
	WithUint8(uintHeight uint8) Builder
	WithUint16(uintSixteen uint16) Builder
	WithUint32(uintThirtyTwo uint32) Builder
	WithUint64(uintSixtyFour uint64) Builder
	WithFloat32(floatThirtyTwo float32) Builder
	WithFloat64(floatSixtyFour float64) Builder
	IsToken() Builder
	WithToken(tok lexers.NodeTree) Builder
	IsStackFrame() Builder
	IsFrame() Builder
	Now() (Value, error)
}

// Value represnets a computable value
type Value interface {
	IsNil() bool
	IsBool() bool
	Bool() *bool
	IsString() bool
	String() *string
	IsIntHeight() bool
	IntHeight() *int8
	IsIntSixteen() bool
	IntSixteen() *int16
	IsIntThirtyTwo() bool
	IntThirtyTwo() *int32
	IsIntSixtyFour() bool
	IntSixtyFour() *int64
	IsUintHeight() bool
	UintHeight() *uint8
	IsUintSixteen() bool
	UintSixteen() *uint16
	IsUintThirtyTwo() bool
	UintThirtyTwo() *uint32
	IsUintSixtyFour() bool
	UintSixtyFour() *uint64
	IsFloatThirtyTwo() bool
	FloatThirtyTwo() *float32
	IsFloatSixtyFour() bool
	FloatSixtyFour() *float64
	IsToken() bool
	Token() lexers.NodeTree
	IsStackFrame() bool
	IsFrame() bool
	StringRepresentation() string
}
