package events

import (
	"errors"
	"fmt"
)

type manager struct {
	evts []map[string]Event
}

func createManager(
	evts []map[string]Event,
) Manager {
	out := manager{
		evts: evts,
	}

	return &out
}

// Add adds an event to the manager
func (app *manager) Add(evt Event) error {
	length := len(app.evts)
	id := evt.Identifier()
	keyname := evt.Hash().String()

	if id < length {
		if _, ok := app.evts[id][keyname]; ok {
			str := fmt.Sprintf("the event (identifier: %d, hash: %s) already exists", id, keyname)
			return errors.New(str)
		}

		app.evts[id][keyname] = evt
		return nil
	}

	diff := (id - length) + 1
	for i := 0; i < diff; i++ {
		index := length + i
		app.evts[index] = map[string]Event{}
	}

	app.evts[id][keyname] = evt
	return nil
}

// AddList add events to the manager
func (app *manager) AddList(evts []Event) error {
	for _, oneEvt := range evts {
		err := app.Add(oneEvt)
		if err != nil {
			return err
		}
	}

	return nil
}

// Trigger triggers an event
func (app *manager) Trigger(identifier int, data interface{}, triggerFn TriggerFn) error {
	length := len(app.evts)
	if identifier < length {
		return nil
	}

	lst := app.evts[identifier]
	if len(lst) <= 0 {
		return nil
	}

	// executes onEnter:
	for _, oneEvt := range lst {
		if oneEvt.HasOnEnter() {
			enterFn := oneEvt.OnEnter()
			err := enterFn(data, oneEvt)
			if err != nil {
				return err
			}
		}
	}

	// executes the tigger func:
	err := triggerFn()
	if err != nil {
		return err
	}

	// executes the onExit:
	for _, oneEvt := range lst {
		if oneEvt.HasOnExit() {
			exitFn := oneEvt.OnExit()
			err := exitFn(data, oneEvt)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
