package propositions

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	gov         governments.Government
	section     Section
	activeOn    *time.Time
	deadline    *time.Time
	createdOn   *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		gov:       nil,
		section:   nil,
		activeOn:  nil,
		deadline:  nil,
		createdOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter)
}

// WithGovernment adds a government to the builder
func (app *contentBuilder) WithGovernment(gov governments.Government) ContentBuilder {
	app.gov = gov
	return app
}

// WithSection adds a section to the builder
func (app *contentBuilder) WithSection(section Section) ContentBuilder {
	app.section = section
	return app
}

// WithDeadline adds a deadline to the builder
func (app *contentBuilder) WithDeadline(deadline time.Time) ContentBuilder {
	app.deadline = &deadline
	return app
}

// ActiveOn adds an activation time to the builder
func (app *contentBuilder) ActiveOn(activeOn time.Time) ContentBuilder {
	app.activeOn = &activeOn
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.gov == nil {
		return nil, errors.New("the government is mandatory in order to build a Content instance")
	}

	if app.section == nil {
		return nil, errors.New("the section is mandatory in order to build a Content instance")
	}

	if app.deadline == nil {
		return nil, errors.New("the deadline is mandatory in order to build a Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.activeOn == nil {
		activatedOn := time.Now().UTC()
		app.activeOn = &activatedOn
	}

	if app.activeOn.Before(*app.createdOn) {
		str := fmt.Sprintf("the activation time (%s) cannot be before the creation time (%s)", app.activeOn.String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	if app.deadline.Before(*app.createdOn) {
		str := fmt.Sprintf("the deadline time (%s) cannot be before the creation time (%s)", app.deadline.String(), app.createdOn.String())
		return nil, errors.New(str)
	}

	if app.activeOn.Before(*app.deadline) {
		str := fmt.Sprintf("the activation time (%s) cannot be before the deadline time (%s)", app.activeOn.String(), app.deadline.String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.gov.Hash().Bytes(),
		app.section.Hash().Bytes(),
		[]byte(strconv.Itoa(app.activeOn.Second())),
		[]byte(strconv.Itoa(app.deadline.Second())),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.gov, app.section, *app.activeOn, *app.deadline, *app.createdOn), nil
}
