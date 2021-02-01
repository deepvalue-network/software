package hydro

const doubleStringPattern = "%s/%s"

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
	WithDehydratedInterface(dehydratedInterface interface{}) BridgeBuilder
	WithDehydratedConstructor(dehydratedConstructor interface{}) BridgeBuilder
	WithDehydratedPointer(dehydratedPointer interface{}) BridgeBuilder
	WithHydratedPointer(hydratedPointer interface{}) BridgeBuilder
	OnHydrate(onHydrateFn EventFn) BridgeBuilder
	OnDehydrate(onDehydrateFn EventFn) BridgeBuilder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Hydrated() Hydrated
	Dehydrated() Dehydrated
}

// Hydrated represents an hydrated bridge
type Hydrated interface {
	Pointer() interface{}
	HasEvent() bool
	Event() EventFn
}

// Dehydrated represents a dehydrated bridge
type Dehydrated interface {
	Interface() interface{}
	ConstructorFn() interface{}
	Pointer() interface{}
	HasEvent() bool
	Event() EventFn
}
