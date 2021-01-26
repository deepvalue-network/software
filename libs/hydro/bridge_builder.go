package hydro

import "errors"

type bridgeBuilder struct {
	pointer        interface{}
	constructorFn  interface{}
	interfaceName  string
	structName     string
	onHydrateFn   EventFn
	onDehydrateFn EventFn
}

func createBridgeBuilder() BridgeBuilder {
	out := bridgeBuilder{
		pointer:        nil,
		constructorFn:  nil,
		interfaceName:  "",
		structName:     "",
		onHydrateFn:   nil,
		onDehydrateFn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *bridgeBuilder) Create() BridgeBuilder {
	return createBridgeBuilder()
}

// WithPointer adds a pointer to the builder
func (app *bridgeBuilder) WithPointer(pointer interface{}) BridgeBuilder {
	app.pointer = pointer
	return app
}

// WithConstructor adds a constructor func to the builder
func (app *bridgeBuilder) WithConstructor(constructorFn interface{}) BridgeBuilder {
	app.constructorFn = constructorFn
	return app
}

// WithInterfaceName adds an interfaceName to the builder
func (app *bridgeBuilder) WithInterfaceName(interfaceName string) BridgeBuilder {
	app.interfaceName = interfaceName
	return app
}

// WithStructName adds a structName to the builder
func (app *bridgeBuilder) WithStructName(structName string) BridgeBuilder {
	app.structName = structName
	return app
}

// OnHydrate adds an OnHydrateOn before event func to the builder
func (app *bridgeBuilder) OnHydrate(onHydrateFn EventFn) BridgeBuilder {
	app.onHydrateFn = onHydrateFn
	return app
}

// OnDehydrate adds an OnDehydrateOn before event func to the builder
func (app *bridgeBuilder) OnDehydrate(onDehydrateFn EventFn) BridgeBuilder {
	app.onDehydrateFn = onDehydrateFn
	return app
}

// Now builds a new Bridge instance
func (app *bridgeBuilder) Now() (Bridge, error) {
	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Bridge instance")
	}

	if app.constructorFn == nil {
		return nil, errors.New("the constructor func is mandatory in order to build a Bridge instance")
	}

	if app.interfaceName == "" {
		return nil, errors.New("the interface name is mandatory in order ro build a Bridge instance")
	}

	if app.structName == "" {
		return nil, errors.New("the struct name is mandatory in order to build a Bridge instance")
	}

	if app.onHydrateFn != nil && app.onDehydrateFn != nil {
		evts := createEventsWithHydrateAndDehydrate(app.onHydrateFn, app.onDehydrateFn)
		return createBridgeWithEvents(app.constructorFn, app.pointer, app.interfaceName, app.structName, evts), nil
	}

	if app.onHydrateFn != nil {
		evts := createEventsWithHydrate(app.onHydrateFn)
		return createBridgeWithEvents(app.constructorFn, app.pointer, app.interfaceName, app.structName, evts), nil
	}

	if app.onDehydrateFn != nil {
		evts := createEventsWithDehydrate(app.onDehydrateFn)
		return createBridgeWithEvents(app.constructorFn, app.pointer, app.interfaceName, app.structName, evts), nil
	}

	return createBridge(app.constructorFn, app.pointer, app.interfaceName, app.structName), nil
}
