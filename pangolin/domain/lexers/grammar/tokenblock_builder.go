package grammar

import "errors"

type tokenBlockBuilder struct {
	section           TokenSection
	optional          RawToken
	multipleOptional  RawToken
	multipleMandatory RawToken
}

func createTokenBlockBuilder() TokenBlockBuilder {
	out := tokenBlockBuilder{
		section:           nil,
		optional:          nil,
		multipleOptional:  nil,
		multipleMandatory: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBlockBuilder) Create() TokenBlockBuilder {
	return createTokenBlockBuilder()
}

// WithSection adds a section to the builder
func (app *tokenBlockBuilder) WithSection(section TokenSection) TokenBlockBuilder {
	app.section = section
	return app
}

// WithOptional adds an optional RawToken to the builder
func (app *tokenBlockBuilder) WithOptional(optional RawToken) TokenBlockBuilder {
	app.optional = optional
	return app
}

// WithMultipleOptional adds a multipleOptional RawToken to the builder
func (app *tokenBlockBuilder) WithMultipleOptional(multipleOptional RawToken) TokenBlockBuilder {
	app.multipleOptional = multipleOptional
	return app
}

// WithMultipleMandatory adds a multipleMandatory RawToken to the builder
func (app *tokenBlockBuilder) WithMultipleMandatory(multipleMandatory RawToken) TokenBlockBuilder {
	app.multipleMandatory = multipleMandatory
	return app
}

// Now builds a new TokenBlock instance
func (app *tokenBlockBuilder) Now() (TokenBlock, error) {
	if app.section == nil {
		return nil, errors.New("the TokenSection is mandatory in order to build a TokenBlock instance")
	}

	if app.optional != nil {
		return createTokenBlockWithOptional(app.section, app.optional), nil
	}

	if app.multipleOptional != nil {
		return createTokenBlockWithMultipleOptional(app.section, app.multipleOptional), nil
	}

	if app.multipleMandatory != nil {
		return createTokenBlockWithMultipleMandatory(app.section, app.multipleMandatory), nil
	}

	return createTokenBlock(app.section), nil
}
