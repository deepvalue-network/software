package variable

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
)

type adapterBuilder struct {
	valueFactory value.Factory
	valueAdapter value.Adapter
	builder      Builder
	isGlobal     bool
}

func createAdapterBuilder(
	valueFactory value.Factory,
	valueAdapter value.Adapter,
	builder Builder,
) AdapterBuilder {
	out := adapterBuilder{
		valueFactory: valueFactory,
		valueAdapter: valueAdapter,
		builder:      builder,
		isGlobal:     false,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(app.valueFactory, app.valueAdapter, app.builder)
}

// IsGlobal sets the builder as global
func (app *adapterBuilder) IsGlobal() AdapterBuilder {
	app.isGlobal = true
	return app
}

// Now builds a new adapter instance
func (app *adapterBuilder) Now() Adapter {
	return createAdapter(app.valueFactory, app.valueAdapter, app.builder, app.isGlobal)
}
