package hydro

type bridge struct {
	consFn interface{}
	ptr    interface{}
	in     string
	strct  string
	evts   Events
}

func createBridge(
	consFn interface{},
	ptr interface{},
	in string,
	strct string,
) Bridge {
	return createBridgeinternally(consFn, ptr, in, strct, nil)
}

func createBridgeWithEvents(
	consFn interface{},
	ptr interface{},
	in string,
	strct string,
	evts Events,
) Bridge {
	return createBridgeinternally(consFn, ptr, in, strct, evts)
}

func createBridgeinternally(
	consFn interface{},
	ptr interface{},
	in string,
	strct string,
	evts Events,
) Bridge {
	out := bridge{
		consFn: consFn,
		ptr:    ptr,
		in:     in,
		strct:  strct,
		evts:   evts,
	}

	return &out
}

// ConstructorFn returns the constructor func encapsulated in a value reflection
func (obj *bridge) ConstructorFn() interface{} {
	return obj.consFn
}

// Pointer returns the struct pointer
func (obj *bridge) Pointer() interface{} {
	return obj.ptr
}

// Interface returns the interface name
func (obj *bridge) Interface() string {
	return obj.in
}

// Struct returns the struct name
func (obj *bridge) Struct() string {
	return obj.strct
}

// HasEvents returns true if there is events, false otherwise
func (obj *bridge) HasEvents() bool {
	return obj.evts != nil
}

// Events returns the events, if any
func (obj *bridge) Events() Events {
	return obj.evts
}
