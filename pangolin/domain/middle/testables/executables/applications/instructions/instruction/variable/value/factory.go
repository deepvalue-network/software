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
	builder := app.builder.Create()
	if typ.IsStackFrame() {
		return builder.IsStackFrame().Now()
	}

	computableBuilder := app.computableBuilder.Create()
	if typ.IsBool() {
		computableBuilder.WithBool(defaultBool)
	}

	if typ.IsString() {
		computableBuilder.WithString(defaultString)
	}

	if typ.IsInt8() {
		computableBuilder.WithInt8(int8(defaultInt))
	}

	if typ.IsInt16() {
		computableBuilder.WithInt16(int16(defaultInt))
	}

	if typ.IsInt32() {
		computableBuilder.WithInt32(int32(defaultInt))
	}

	if typ.IsInt64() {
		computableBuilder.WithInt64(int64(defaultInt))
	}

	if typ.IsUint8() {
		computableBuilder.WithUint8(uint8(defaultUint))
	}

	if typ.IsUint16() {
		computableBuilder.WithUint16(uint16(defaultUint))
	}

	if typ.IsUint32() {
		computableBuilder.WithUint32(uint32(defaultUint))
	}

	if typ.IsUint64() {
		computableBuilder.WithUint64(uint64(defaultUint))
	}

	if typ.IsFloat32() {
		computableBuilder.WithFloat32(float32(defaultFloat))
	}

	if typ.IsFloat64() {
		computableBuilder.WithFloat64(float64(defaultFloat))
	}

	computable, err := computableBuilder.Now()
	if err != nil {
		return nil, err
	}

	return builder.WithComputable(computable).Now()
}
