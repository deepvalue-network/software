package views

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter      hash.Adapter
	minHashesInOwner uint
	section          Section
	newOwner         []hash.Hash
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	minHashesInOwner uint,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:      hashAdapter,
		minHashesInOwner: minHashesInOwner,
		section:          nil,
		newOwner:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.minHashesInOwner)
}

// WithSection adds a section to the builder
func (app *contentBuilder) WithSection(section Section) ContentBuilder {
	app.section = section
	return app
}

// WithNewOwner adds a new owner to the builder
func (app *contentBuilder) WithNewOwner(newOwner []hash.Hash) ContentBuilder {
	app.newOwner = newOwner
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.section == nil {
		return nil, errors.New("the section is mandatory in order to build a view transaction Content instance")
	}

	if app.newOwner == nil {
		app.newOwner = []hash.Hash{}
	}

	amount := len(app.newOwner)
	if amount < int(app.minHashesInOwner) {
		str := fmt.Sprintf("there must be at least %d public key hashes in the new owner in order to build a view transaction Content instance, %d provided", app.minHashesInOwner, amount)
		return nil, errors.New(str)
	}

	data := [][]byte{
		app.section.Hash().Bytes(),
	}

	for _, oneHash := range app.newOwner {
		data = append(data, oneHash.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createContent(*hash, app.section, app.newOwner), nil
}
