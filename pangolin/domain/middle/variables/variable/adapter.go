package variable

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	valueFactory value.Factory
	valueAdapter value.Adapter
	builder      Builder
	isGLobal     bool
}

func createAdapter(valueFactory value.Factory, valueAdapter value.Adapter, builder Builder, isGLobal bool) Adapter {
	out := adapter{
		valueFactory: valueFactory,
		valueAdapter: valueAdapter,
		builder:      builder,
		isGLobal:     isGLobal,
	}

	return &out
}

// FromConstant converts a constant declaration to a variable
func (app *adapter) FromConstant(declaration parsers.ConstantDeclaration) (Variable, error) {
	parsedValue := declaration.Value()
	val, err := app.valueAdapter.ToValue(parsedValue)
	if err != nil {
		return nil, err
	}

	// make sure the value fits the declaration type:

	name := declaration.Constant()
	builder := app.builder.Create().IsImmutable().WithName(name).WithValue(val)
	if app.isGLobal {
		builder.IsGlobal()
	}

	return builder.Now()
}

// FromVariable converts a variable declaration to a variable
func (app *adapter) FromVariable(declaration parsers.VariableDeclaration) (Variable, error) {
	typ := declaration.Type()
	val, err := app.valueFactory.Create(typ)
	if err != nil {
		return nil, err
	}

	name := declaration.Variable()
	builder := app.builder.Create().WithName(name).WithValue(val)
	if app.isGLobal {
		builder.IsGlobal()
	}

	if declaration.HasDirection() {
		direction := declaration.Direction()
		if direction.IsIncoming() {
			incoming := direction.Incoming()
			if incoming.IsMandatory() {
				builder.IsMandatory()
			}

			if incoming.IsOptional() {
				parsedDefaultValue := incoming.OptionalDefaultValue()
				defaultValue, err := app.valueAdapter.ToValueWithType(typ, parsedDefaultValue)
				if err != nil {
					return nil, err
				}

				builder.WithValue(defaultValue)
			}

			builder.IsIncoming()
		}

		if direction.IsOutgoing() {
			builder.IsOutgoing()
		}
	}

	return builder.Now()
}
