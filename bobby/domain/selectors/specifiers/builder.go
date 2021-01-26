package specifiers

import "errors"

type builder struct {
	identifiersBuilder IdentifiersBuilder
	identifier         Identifier
	identifiers        []Identifier
}

func createBuilder(
	identifiersBuilder IdentifiersBuilder,
) Builder {
	out := builder{
		identifiersBuilder: identifiersBuilder,
		identifier:         nil,
		identifiers:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.identifiersBuilder)
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier Identifier) Builder {
	app.identifier = identifier
	return app
}

// WithIdentifiers add identifiers to the builder
func (app *builder) WithIdentifiers(identifiers []Identifier) Builder {
	app.identifiers = identifiers
	return app
}

// Now builds a new Specifier instance
func (app *builder) Now() (Specifier, error) {
	if app.identifier != nil {
		return createSpecifierWithIdentifier(app.identifier), nil
	}

	if app.identifiers != nil {
		list, err := app.identifiersBuilder.Create().WithIdentifiers(app.identifiers).Now()
		if err != nil {
			return nil, err
		}

		return createSpecifierWithIdentifiers(list), nil
	}

	return nil, errors.New("the Specifier is invalid")
}
