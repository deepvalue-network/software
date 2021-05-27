package value

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	computableBuilder computable.Builder
	builder           Builder
}

func createAdapter(computableBuilder computable.Builder, builder Builder) Adapter {
	out := adapter{
		computableBuilder: computableBuilder,
		builder:           builder,
	}

	return &out
}

// ToValue converts a parsed value to an optimized value
func (app *adapter) ToValue(parsed parsers.ValueRepresentation) (Value, error) {
	builder := app.builder.Create()
	if parsed.IsValue() {
		parsedValue := parsed.Value()
		fl, in, bl, str := app.toValue(parsedValue)
		computableBuilder := app.computableBuilder.Create()
		if fl != nil {
			computableBuilder.WithFloat64(*fl)
		}

		if in != nil {
			computableBuilder.WithInt64(int64(*in))
		}

		if bl != nil {
			computableBuilder.WithBool(*bl)
		}

		if str != nil {
			computableBuilder.WithString(*str)
		}

		computable, err := computableBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithComputable(computable)
	}

	if parsed.IsVariable() {
		parsedVariable := parsed.Variable()
		builder.WithVariable(parsedVariable)
	}

	return builder.Now()
}

// ToValueWithType converts a parsed value to an optimized value and optimize using the type
func (app *adapter) ToValueWithType(typ parsers.Type, parsed parsers.ValueRepresentation) (Value, error) {
	builder := app.builder.Create()
	if parsed.IsValue() {
		parsedValue := parsed.Value()
		fl, in, bl, str := app.toValue(parsedValue)
		computableBuilder := app.computableBuilder.Create()
		if fl != nil {
			val := *fl
			if typ.IsFloat32() {
				computableBuilder.WithFloat32(float32(val))
			}

			if typ.IsFloat64() {
				computableBuilder.WithFloat64(float64(val))
			}
		}

		if in != nil {
			val := *in

			if typ.IsInt8() {
				computableBuilder.WithInt8(int8(val))
			}

			if typ.IsInt16() {
				computableBuilder.WithInt16(int16(val))
			}

			if typ.IsInt32() {
				computableBuilder.WithInt32(int32(val))
			}

			if typ.IsInt64() {
				computableBuilder.WithInt64(int64(val))
			}

			if typ.IsUint8() || typ.IsUint16() || typ.IsUint32() || typ.IsUint32() {
				if val < 0 {
					str := fmt.Sprintf("the uint type (%s) cannot contain a negative value (%d)", typ.String(), val)
					return nil, errors.New(str)
				}
			}

			if typ.IsUint8() {
				computableBuilder.WithUint8(uint8(val))
			}

			if typ.IsUint16() {
				computableBuilder.WithUint16(uint16(val))
			}

			if typ.IsUint32() {
				computableBuilder.WithUint32(uint32(val))
			}

			if typ.IsUint64() {
				computableBuilder.WithUint64(uint64(val))
			}
		}

		if bl != nil {
			computableBuilder.WithBool(*bl)
		}

		if str != nil {
			computableBuilder.WithString(*str)
		}

		computable, err := computableBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithComputable(computable)
	}

	if parsed.IsVariable() {
		parsedVariable := parsed.Variable()
		builder.WithVariable(parsedVariable)
	}

	return builder.Now()
}

func (app *adapter) toValue(parsed parsers.Value) (*float64, *int, *bool, *string) {
	if parsed.IsNumeric() {
		numeric := parsed.Numeric()
		isNegative := numeric.IsNegative()
		if numeric.IsFloat() {
			fl := numeric.Float()
			val := *fl
			if isNegative {
				val *= -1.0
			}

			return &val, nil, nil, nil
		}

		in := numeric.Int()
		val := *in
		if isNegative {
			val *= -1
		}

		return nil, &val, nil, nil
	}

	if parsed.IsBool() {
		bl := parsed.Bool()
		return nil, nil, bl, nil
	}

	if parsed.IsString() {
		str := parsed.String()
		return nil, nil, nil, &str
	}

	return nil, nil, nil, nil
}
