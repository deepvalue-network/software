package value

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value/computable"
)

type builder struct {
	comp           computable.Value
	globalVariable string
	localVariable  string
}

func createBuilder() Builder {
	out := builder{
		comp:           nil,
		globalVariable: "",
		localVariable:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithComputable adds a computable to the builder
func (app *builder) WithComputable(computable computable.Value) Builder {
	app.comp = computable
	return app
}

// WithGlobalVariable adds a globalVariable to the builder
func (app *builder) WithGlobalVariable(globalVariable string) Builder {
	app.globalVariable = globalVariable
	return app
}

// WithGlobalVariable adds a globalVariable to the builder
func (app *builder) WithLocalVariable(localVariable string) Builder {
	app.localVariable = localVariable
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.comp != nil {
		return createValueWithComputable(app.comp), nil
	}

	if app.globalVariable != "" {
		return createValueWithGlobalVariabe(app.globalVariable), nil
	}

	if app.localVariable != "" {
		return createValueWithLocalVariabe(app.localVariable), nil
	}

	return nil, errors.New("the Value is invalid")
}
