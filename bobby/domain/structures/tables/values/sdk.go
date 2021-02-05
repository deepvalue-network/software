package values

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/bobby/domain/resources"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the value builder
type Builder interface {
	Create() Builder
	WithResource(res resources.Resource) Builder
	WithID(id *uuid.UUID) Builder
	WithString(stringVal string) Builder
	WithInt(intVal int) Builder
	WithFloat32(float32Val float32) Builder
	WithFloat64(float64Val float64) Builder
	WithData(data []byte) Builder
	Now() (Value, error)
}

// Value represents an element's value
type Value interface {
	Resource() resources.Resource
	Content() ValueContent
}

// ValueContent represents the content of a value
type ValueContent interface {
	IsID() bool
	ID() *uuid.UUID
	IsString() bool
	String() *string
	IsInt() bool
	Int() *int
	IsFloat32() bool
	Float32() *float32
	IsFloat64() bool
	Float64() *float64
	IsData() bool
	Data() []byte
}
