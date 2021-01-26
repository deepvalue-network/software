package hydro

// EventFn represents an event func
type EventFn func(ins interface{}, fieldName string, structName string) (interface{}, error)

// NewAdapterBuilder creates a new adapter builder instance
func NewAdapterBuilder() AdapterBuilder {
	return createAdapterBuilder()
}

// NewManagerFactory creates a new manager factory instance
func NewManagerFactory() ManagerFactory {
	return createManagerFactory()
}

// NewBridgeBuilder creates a new bridge builder instance
func NewBridgeBuilder() BridgeBuilder {
	return createBridgeBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithManager(manager Manager) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents an adapter
type Adapter interface {
	Hydrate(dehydrate interface{}) (interface{}, error)
	Dehydrate(hydrate interface{}) (interface{}, error)
}

// ManagerFactory represents a manager factory
type ManagerFactory interface {
	Create() Manager
}

// Manager represents a manager
type Manager interface {
	Fetch(pkg string, name string) (Bridge, error)
	Register(bridge Bridge)
}

// BridgeBuilder represents a bridge builder
type BridgeBuilder interface {
	Create() BridgeBuilder
	WithPointer(pointer interface{}) BridgeBuilder
	WithConstructor(constructorFn interface{}) BridgeBuilder
	WithInterfaceName(interfaceName string) BridgeBuilder
	WithStructName(structName string) BridgeBuilder
	OnHydrate(onHydrateFn EventFn) BridgeBuilder
	OnDehydrate(onDehydrateFn EventFn) BridgeBuilder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	ConstructorFn() interface{}
	Pointer() interface{}
	Interface() string
	Struct() string
	HasEvents() bool
	Events() Events
}

// Events represents bridge events
type Events interface {
	HasOnHydrate() bool
	OnHydrate() EventFn
	HasOnDehydrate() bool
	OnDehydrate() EventFn
}
