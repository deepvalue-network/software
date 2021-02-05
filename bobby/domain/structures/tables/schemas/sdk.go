package schemas

import (
	"github.com/deepvalue-network/software/bobby/domain/resources"
)

// NewBuilder creates a new schema builder
func NewBuilder() Builder {
	propertiesBuilder := NewPropertiesBuilder()
	return createBuilder(propertiesBuilder)
}

// NewPropertiesBuilder creates a new properties builder
func NewPropertiesBuilder() PropertiesBuilder {
	return createPropertiesBuilder()
}

// NewPropertyBuilder creates a new property builder
func NewPropertyBuilder() PropertyBuilder {
	return createPropertyBuilder()
}

// NewTypeBuilder creates a new type builder instance
func NewTypeBuilder() TypeBuilder {
	return createTypeBuilder()
}

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithResource(res resources.Accessible) Builder
	WithName(name string) Builder
	WithProperties(properties []Property) Builder
	Now() (Schema, error)
}

// Schema represents a table schema
type Schema interface {
	Resource() resources.Accessible
	Name() string
	Properties() Properties
}

// PropertiesBuilder represents a properties builder
type PropertiesBuilder interface {
	Create() PropertiesBuilder
	WithResource(res resources.Accessible) PropertiesBuilder
	WithProperties(properties []Property) PropertiesBuilder
	Now() (Properties, error)
}

// Properties represents a list of properties
type Properties interface {
	Resource() resources.Accessible
	All() []Property
	First() Property
	IsEmpty() bool
}

// PropertyBuilder represents a property builder
type PropertyBuilder interface {
	Create() PropertyBuilder
	WithResource(res resources.Accessible) PropertyBuilder
	WithName(name string) PropertyBuilder
	IsPrimaryKey() PropertyBuilder
	WithForeignKey(foreignKey Schema) PropertyBuilder
	WithType(typ Type) PropertyBuilder
	Now() (Property, error)
}

// Property represents a schema property
type Property interface {
	Resource() resources.Accessible
	Content() PropertyContent
	Name() string
}

// PropertyContent represents a property content
type PropertyContent interface {
	IsPrimaryKey() bool
	IsForeignKey() bool
	ForeignKey() Schema
	IsType() bool
	Type() Type
}

// TypeBuilder represents a type builder
type TypeBuilder interface {
	Create() TypeBuilder
	IsString() TypeBuilder
	IsInt() TypeBuilder
	IsFloat32() TypeBuilder
	IsFloat64() TypeBuilder
	IsData() TypeBuilder
	Now() (Type, error)
}

// Type represents a property type
type Type interface {
	IsString() bool
	IsInt() bool
	IsFloat32() bool
	IsFloat64() bool
	IsData() bool
}
