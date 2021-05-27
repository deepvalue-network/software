package value

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type factory struct {
	computableBuilder computable.Builder
	builder           Builder
}

func createFactory(computableBuilder computable.Builder, builder Builder) Factory {
	out := factory{
		computableBuilder: computableBuilder,
		builder:           builder,
	}

	return &out
}

// Create creates a value based on a type
func (app *factory) Create(typ parsers.Type) (Value, error) {
	builder := app.computableBuilder.Create()
	if typ.IsBool() {
		builder.WithBool(defaultBool)
	}

	if typ.IsString() {
		builder.WithString(defaultString)
	}

	if typ.IsInt8() {
		builder.WithInt8(int8(defaultInt))
	}

	if typ.IsInt16() {
		builder.WithInt16(int16(defaultInt))
	}

	if typ.IsInt32() {
		builder.WithInt32(int32(defaultInt))
	}

	if typ.IsInt64() {
		builder.WithInt64(int64(defaultInt))
	}

	if typ.IsUint8() {
		builder.WithUint8(uint8(defaultUint))
	}

	if typ.IsUint16() {
		builder.WithUint16(uint16(defaultUint))
	}

	if typ.IsUint32() {
		builder.WithUint32(uint32(defaultUint))
	}

	if typ.IsUint64() {
		builder.WithUint64(uint64(defaultUint))
	}

	if typ.IsFloat32() {
		builder.WithFloat32(float32(defaultFloat))
	}

	if typ.IsFloat64() {
		builder.WithFloat64(float64(defaultFloat))
	}

	computable, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithComputable(computable).Now()
}
