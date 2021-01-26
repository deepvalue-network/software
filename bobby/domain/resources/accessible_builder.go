package resources

import "errors"

type accessibleBuilder struct {
	immutable ImmutableAccessible
	mutable   MutableAccessible
}

func createAccessibleBuilder() AccessibleBuilder {
	out := accessibleBuilder{
		immutable: nil,
		mutable:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *accessibleBuilder) Create() AccessibleBuilder {
	return createAccessibleBuilder()
}

// WithImmutable adds an immutable to the builder
func (app *accessibleBuilder) WithImmutable(immutable ImmutableAccessible) AccessibleBuilder {
	app.immutable = immutable
	return app
}

// WithMutable adds a mutable to the builder
func (app *accessibleBuilder) WithMutable(mutable MutableAccessible) AccessibleBuilder {
	app.mutable = mutable
	return app
}

// Now builds a new Accessible instance
func (app *accessibleBuilder) Now() (Accessible, error) {
	if app.immutable != nil {
		return createAccessibleWithImmutable(app.immutable), nil
	}

	if app.mutable != nil {
		return createAccessibleWithMutable(app.mutable), nil
	}

	return nil, errors.New("the Accessible is invalid")
}
