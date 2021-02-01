package hydro

type hydrated struct {
	ptr   interface{}
	event EventFn
}

func createHydrated(
	ptr interface{},
) Hydrated {
	return createHydratedInternally(ptr, nil)
}

func createHydratedWithEvent(
	ptr interface{},
	event EventFn,
) Hydrated {
	return createHydratedInternally(ptr, event)
}

func createHydratedInternally(
	ptr interface{},
	event EventFn,
) Hydrated {
	out := hydrated{
		ptr:   ptr,
		event: event,
	}

	return &out
}

// Pointer returns the pointer
func (obj *hydrated) Pointer() interface{} {
	return obj.ptr
}

// HasEvent returns true if there is an event, false otherwise
func (obj *hydrated) HasEvent() bool {
	return obj.event != nil
}

// Event returns the event, if any
func (obj *hydrated) Event() EventFn {
	return obj.event
}
