package hydro

type dehydrated struct {
	in          interface{}
	constructor interface{}
	ptr         interface{}
	event       EventFn
}

func createDehydrated(
	in interface{},
	constructor interface{},
	ptr interface{},
) Dehydrated {
	return createDehydratedInternally(in, constructor, ptr, nil)
}

func createDehydratedWithEvent(
	in interface{},
	constructor interface{},
	ptr interface{},
	event EventFn,
) Dehydrated {
	return createDehydratedInternally(in, constructor, ptr, event)
}

func createDehydratedInternally(
	in interface{},
	constructor interface{},
	ptr interface{},
	event EventFn,
) Dehydrated {
	out := dehydrated{
		in:          in,
		constructor: constructor,
		ptr:         ptr,
		event:       event,
	}

	return &out
}

// Interface returns the interface
func (obj *dehydrated) Interface() interface{} {
	return obj.in
}

// ConstructorFn returns the constructor func
func (obj *dehydrated) ConstructorFn() interface{} {
	return obj.constructor
}

// Pointer returns the pointer
func (obj *dehydrated) Pointer() interface{} {
	return obj.ptr
}

// HasEvent returns true if there is an event, false otherwise
func (obj *dehydrated) HasEvent() bool {
	return obj.event != nil
}

// Event returns the event, if any
func (obj *dehydrated) Event() EventFn {
	return obj.event
}
