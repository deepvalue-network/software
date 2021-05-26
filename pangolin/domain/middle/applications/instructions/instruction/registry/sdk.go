package registry

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRegisterBuilder creates a new register builder
func NewRegisterBuilder() RegisterBuilder {
	return createRegisterBuilder()
}

// NewFetchBuilder creates a new fetch builder
func NewFetchBuilder() FetchBuilder {
	return createFetchBuilder()
}

// NewIndexBuilder creates a new index builder instance
func NewIndexBuilder() IndexBuilder {
	return createIndexBuilder()
}

// Builder represents a registry builder
type Builder interface {
	Create() Builder
	WithFetch(fetch Fetch) Builder
	WithRegister(reg Register) Builder
	WithUnregister(unregister string) Builder
	Now() (Registry, error)
}

// Registry represents a registry
type Registry interface {
	IsFetch() bool
	Fetch() Fetch
	IsRegister() bool
	Register() Register
	IsUnregister() bool
	Unregister() string
}

// FetchBuilder represents a fetch builder
type FetchBuilder interface {
	Create() FetchBuilder
	From(from string) FetchBuilder
	To(to string) FetchBuilder
	WithIndex(index Index) FetchBuilder
	Now() (Fetch, error)
}

// Fetch represents a fetch registry
type Fetch interface {
	From() string
	To() string
	HasIndex() bool
	Index() Index
}

// RegisterBuilder represents a register builder
type RegisterBuilder interface {
	Create() RegisterBuilder
	WithVariable(variable string) RegisterBuilder
	WithIndex(index Index) RegisterBuilder
	Now() (Register, error)
}

// Register represents a register
type Register interface {
	Variable() string
	HasIndex() bool
	Index() Index
}

// IndexBuilder represents an index bulder
type IndexBuilder interface {
	Create() IndexBuilder
	WithInt(intVal int64) IndexBuilder
	WithVariable(variable string) IndexBuilder
	Now() (Index, error)
}

// Index represents an index
type Index interface {
	IsInt() bool
	Int() int64
	IsVariable() bool
	Variable() string
}
