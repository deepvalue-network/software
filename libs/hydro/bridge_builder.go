package hydro

import (
	"errors"
)

type bridgeBuilder struct {
	dehydratedInterface   interface{}
	dehydratedConstructor interface{}
	dehydratedPointer     interface{}
	hydratedPointer       interface{}
	onHydrateFn           EventFn
	onDehydrateFn         EventFn
}

func createBridgeBuilder() BridgeBuilder {
	out := bridgeBuilder{
		dehydratedInterface:   "",
		dehydratedConstructor: nil,
		dehydratedPointer:     nil,
		hydratedPointer:       nil,
		onHydrateFn:           nil,
		onDehydrateFn:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *bridgeBuilder) Create() BridgeBuilder {
	return createBridgeBuilder()
}

// WithDehydratedInterface adds a dehydrated interface to the builder
func (app *bridgeBuilder) WithDehydratedInterface(dehydratedInterface interface{}) BridgeBuilder {
	app.dehydratedInterface = dehydratedInterface
	return app
}

// WithDehydratedConstructor adds a dehydrated constructor to the builder
func (app *bridgeBuilder) WithDehydratedConstructor(dehydratedConstructor interface{}) BridgeBuilder {
	app.dehydratedConstructor = dehydratedConstructor
	return app
}

// WithDehydratedPointer adds a dehydrated pointer to the builder
func (app *bridgeBuilder) WithDehydratedPointer(dehydratedPointer interface{}) BridgeBuilder {
	app.dehydratedPointer = dehydratedPointer
	return app
}

// WithHydratedPointer adds an hydrated pointer to the builder
func (app *bridgeBuilder) WithHydratedPointer(hydratedPointer interface{}) BridgeBuilder {
	app.hydratedPointer = hydratedPointer
	return app
}

// OnHydrate adds an onHydrate event to the builder
func (app *bridgeBuilder) OnHydrate(onHydrateFn EventFn) BridgeBuilder {
	app.onHydrateFn = onHydrateFn
	return app
}

// OnDehydrate adds an onDehydrate event to the builder
func (app *bridgeBuilder) OnDehydrate(onDehydrateFn EventFn) BridgeBuilder {
	app.onDehydrateFn = onDehydrateFn
	return app
}

// Now builds a new Bridge instance
func (app *bridgeBuilder) Now() (Bridge, error) {
	if app.dehydratedInterface == nil {
		return nil, errors.New("the dehydrated interface name is mandatory in order to build a Bridge instance")
	}

	if app.dehydratedConstructor == nil {
		return nil, errors.New("the dehydrated constructor is mandatory in order to build a Bridge instance")
	}

	if app.dehydratedPointer == nil {
		return nil, errors.New("the dehydrated pointer is mandatory in order to build a Bridge instance")
	}

	if app.hydratedPointer == nil {
		return nil, errors.New("the hydrated pointer is mandatory in order to build a Bridge instance")
	}

	var hydrated Hydrated
	if app.onHydrateFn != nil {
		hydrated = createHydratedWithEvent(app.hydratedPointer, app.onHydrateFn)
	}

	if hydrated == nil {
		hydrated = createHydrated(app.hydratedPointer)
	}

	var dehydrated Dehydrated
	if app.onDehydrateFn != nil {
		dehydrated = createDehydratedWithEvent(app.dehydratedInterface, app.dehydratedConstructor, app.dehydratedPointer, app.onDehydrateFn)
	}

	if dehydrated == nil {
		dehydrated = createDehydrated(app.dehydratedInterface, app.dehydratedConstructor, app.dehydratedPointer)
	}

	return createBridge(hydrated, dehydrated), nil
}
