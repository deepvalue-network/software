package propositions

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
)

type sectionBuilder struct {
	gov     governments.Content
	holders shareholders.ShareHolders
	custom  *hash.Hash
}

func createSectionBuilder() SectionBuilder {
	out := sectionBuilder{
		gov:     nil,
		holders: nil,
		custom:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *sectionBuilder) Create() SectionBuilder {
	return createSectionBuilder()
}

// WithGovernment adds a government to the builder
func (app *sectionBuilder) WithGovernment(government governments.Content) SectionBuilder {
	app.gov = government
	return app
}

// WithShareHolders add shareHolders to the builder
func (app *sectionBuilder) WithShareHolders(shareHolders shareholders.ShareHolders) SectionBuilder {
	app.holders = shareHolders
	return app
}

// WithCustom adds a custom to the builder
func (app *sectionBuilder) WithCustom(custom hash.Hash) SectionBuilder {
	app.custom = &custom
	return app
}

// Now builds a new Section instance
func (app *sectionBuilder) Now() (Section, error) {
	if app.gov != nil {
		return createSectionWithGovernment(app.gov), nil
	}

	if app.holders != nil {
		return createSectionWithShareHolders(app.holders), nil
	}

	if app.custom != nil {
		return createSectionWithCustom(app.custom), nil
	}

	return nil, errors.New("the proposition Section is invalid")
}
