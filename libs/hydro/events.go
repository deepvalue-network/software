package hydro

type events struct {
	onHydrate   EventFn
	onDehydrate EventFn
}

func createEventsWithHydrate(
	onHydrate EventFn,
) Events {
	return createEventsInternally(onHydrate, nil)
}

func createEventsWithDehydrate(
	onDehydrate EventFn,
) Events {
	return createEventsInternally(nil, onDehydrate)
}

func createEventsWithHydrateAndDehydrate(
	onHydrate EventFn,
	onDehydrate EventFn,
) Events {
	return createEventsInternally(onHydrate, onDehydrate)
}

func createEventsInternally(
	onHydrate EventFn,
	onDehydrate EventFn,
) Events {
	out := events{
		onHydrate:   onHydrate,
		onDehydrate: onDehydrate,
	}

	return &out
}

// HasOnHydrate returns true if there is an onHydrate event, false otherwise
func (obj *events) HasOnHydrate() bool {
	return obj.onHydrate != nil
}

// OnHydrate returns the onHydrate event, if any
func (obj *events) OnHydrate() EventFn {
	return obj.onHydrate
}

// HasOnDehydrate returns true if there is an onDehydrate event, false otherwise
func (obj *events) HasOnDehydrate() bool {
	return obj.onDehydrate != nil
}

// OnDehydrate returns the onDehydrate event, if any
func (obj *events) OnDehydrate() EventFn {
	return obj.onDehydrate
}
