package parsers

import (
	"errors"
	"fmt"
)

type events struct {
	evts []Event
}

func createEvents(evts []Event) Events {
	out := events{
		evts: evts,
	}

	return &out
}

// Events return the []Event instances
func (obj *events) Events() []Event {
	return obj.evts
}

// GetByToken return the Event by token, if any
func (obj *events) GetByToken(token string) (Event, error) {
	for _, oneEvent := range obj.evts {
		if oneEvent.Token() == token {
			return oneEvent, nil
		}
	}

	str := fmt.Sprintf("the Token (Name: %s) is invalid", token)
	return nil, errors.New(str)
}
