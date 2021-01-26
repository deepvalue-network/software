package parsers

type eventsAdapter struct {
	eventsBuilder EventsBuilder
	eventBuilder  EventBuilder
}

func createEventsAdapter(eventsBuilder EventsBuilder, eventBuilder EventBuilder) EventsAdapter {
	out := eventsAdapter{
		eventsBuilder: eventsBuilder,
		eventBuilder:  eventBuilder,
	}

	return &out
}

// ToEvents converts params to an Events instance
func (app *eventsAdapter) ToEvents(params []ToEventsParams) (Events, error) {
	events := []Event{}
	for _, oneParam := range params {
		builder := app.eventBuilder.Create()
		if oneParam.Token != "" {
			builder.WithToken(oneParam.Token)
		}

		if oneParam.Set != nil {
			builder.WithSet(oneParam.Set)
		}

		if oneParam.OnEnter != nil {
			builder.WithOnEnter(oneParam.OnEnter)
		}

		if oneParam.OnExit != nil {
			builder.WithOnExit(oneParam.OnExit)
		}

		if oneParam.RetrieveReplacement != nil {
			builder.WithRetrieveReplacement(oneParam.RetrieveReplacement)
		}

		event, err := builder.Now()
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return app.eventsBuilder.Create().WithEvents(events).Now()
}
