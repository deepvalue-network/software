package parsers

type event struct {
	token               string
	set                 SetFn
	enter               EventFn
	exit                EventFn
	retrieveReplacement RetrieveReplacementsFn
}

func createEventWithOnEnter(token string, enter EventFn) Event {
	return createEventInternally(token, enter, nil, nil, nil)
}

func createEventWithOnEnterAndSet(token string, enter EventFn, set SetFn) Event {
	return createEventInternally(token, enter, nil, nil, set)
}

func createEventWithOnEnterAndRetrieveReplacement(token string, enter EventFn, retrieveReplacement RetrieveReplacementsFn) Event {
	return createEventInternally(token, enter, nil, retrieveReplacement, nil)
}

func createEventWithOnEnterAndRetrieveReplacementAndSet(token string, enter EventFn, retrieveReplacement RetrieveReplacementsFn, set SetFn) Event {
	return createEventInternally(token, enter, nil, retrieveReplacement, set)
}

func createEventWithOnExit(token string, exit EventFn) Event {
	return createEventInternally(token, nil, exit, nil, nil)
}

func createEventWithOnExitAndSet(token string, set SetFn, exit EventFn) Event {
	return createEventInternally(token, nil, exit, nil, set)
}

func createEventWithOnExitAndRetrieveReplacement(token string, exit EventFn, retrieveReplacement RetrieveReplacementsFn) Event {
	return createEventInternally(token, nil, exit, retrieveReplacement, nil)
}

func createEventWithOnExitAndRetrieveReplacementAndSet(token string, exit EventFn, retrieveReplacement RetrieveReplacementsFn, set SetFn) Event {
	return createEventInternally(token, nil, exit, retrieveReplacement, set)
}

func createEventInternally(token string, enter EventFn, exit EventFn, retrieveReplacement RetrieveReplacementsFn, set SetFn) Event {
	out := event{
		token:               token,
		enter:               enter,
		exit:                exit,
		retrieveReplacement: retrieveReplacement,
		set:                 set,
	}

	return &out
}

// Token returns the token
func (obj *event) Token() string {
	return obj.token
}

// HasSet returns true if there is a setFn, false otherwise
func (obj *event) HasSet() bool {
	return obj.set != nil
}

// Set returns the setFn
func (obj *event) Set() SetFn {
	return obj.set
}

// HasOnEnter returns true if the event is an onEnter event
func (obj *event) HasOnEnter() bool {
	return obj.enter != nil
}

// OnEnter returns the onEnter event func
func (obj *event) OnEnter() EventFn {
	return obj.enter
}

// HasOnExit returns true if the event is an onExit event
func (obj *event) HasOnExit() bool {
	return obj.exit != nil
}

// OnExit returns the onExit event func
func (obj *event) OnExit() EventFn {
	return obj.exit
}

// Fn returns the right func based on the input flag, returns nil if nothing is found
func (obj *event) Fn(isEnter bool) EventFn {
	if isEnter {
		return obj.enter
	}

	return obj.exit
}

// HasRetrieveReplacement returns true if there is a retrieveReplacementFn, false otherwise
func (obj *event) HasRetrieveReplacement() bool {
	return obj.retrieveReplacement != nil
}

// RetrieveReplacement returns the master event, if any
func (obj *event) RetrieveReplacement() RetrieveReplacementsFn {
	return obj.retrieveReplacement
}
