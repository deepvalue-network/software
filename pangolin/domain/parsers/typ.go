package parsers

type typ struct {
	isNil     bool
	isBool    bool
	isInt8    bool
	isInt16   bool
	isInt32   bool
	isInt64   bool
	isFloat32 bool
	isFloat64 bool
	isUint8   bool
	isUint16  bool
	isUint32  bool
	isUint64  bool
	isString  bool
}

func createTypeWithNil() Type {
	return createTypeInternally(true, false, false, false, false, false, false, false, false, false, false, false, false)
}

func createTypeWithBool() Type {
	return createTypeInternally(false, true, false, false, false, false, false, false, false, false, false, false, false)
}

func createTypeWithInt8() Type {
	return createTypeInternally(false, false, true, false, false, false, false, false, false, false, false, false, false)
}

func createTypeWithInt16() Type {
	return createTypeInternally(false, false, false, true, false, false, false, false, false, false, false, false, false)
}

func createTypeWithInt32() Type {
	return createTypeInternally(false, false, false, false, true, false, false, false, false, false, false, false, false)
}

func createTypeWithInt64() Type {
	return createTypeInternally(false, false, false, false, false, true, false, false, false, false, false, false, false)
}

func createTypeWithFloat32() Type {
	return createTypeInternally(false, false, false, false, false, false, true, false, false, false, false, false, false)
}

func createTypeWithFloat64() Type {
	return createTypeInternally(false, false, false, false, false, false, false, true, false, false, false, false, false)
}

func createTypeWithUint8() Type {
	return createTypeInternally(false, false, false, false, false, false, false, false, true, false, false, false, false)
}

func createTypeWithUint16() Type {
	return createTypeInternally(false, false, false, false, false, false, false, false, false, true, false, false, false)
}

func createTypeWithUint32() Type {
	return createTypeInternally(false, false, false, false, false, false, false, false, false, false, true, false, false)
}

func createTypeWithUint64() Type {
	return createTypeInternally(false, false, false, false, false, false, false, false, false, false, false, true, false)
}

func createTypeWithString() Type {
	return createTypeInternally(false, false, false, false, false, false, false, false, false, false, false, false, true)
}

func createTypeInternally(
	isNil bool,
	isBool bool,
	isInt8 bool,
	isInt16 bool,
	isInt32 bool,
	isInt64 bool,
	isFloat32 bool,
	isFloat64 bool,
	isUint8 bool,
	isUint16 bool,
	isUint32 bool,
	isUint64 bool,
	isString bool,
) Type {
	out := typ{
		isNil:     isNil,
		isBool:    isBool,
		isInt8:    isInt8,
		isInt16:   isInt16,
		isInt32:   isInt32,
		isInt64:   isInt64,
		isFloat32: isFloat32,
		isFloat64: isFloat64,
		isUint8:   isUint8,
		isUint16:  isUint16,
		isUint32:  isUint32,
		isUint64:  isUint64,
		isString:  isString,
	}

	return &out
}

// IsNil returns true if the type is nil
func (obj *typ) IsNil() bool {
	return obj.isNil
}

// IsBool returns true if the type is bool
func (obj *typ) IsBool() bool {
	return obj.isBool
}

// IsInt8 returns true if the type is int8
func (obj *typ) IsInt8() bool {
	return obj.isInt8
}

// IsInt16 returns true if the type is int16
func (obj *typ) IsInt16() bool {
	return obj.isInt16
}

// IsInt32 returns true if the type is int32
func (obj *typ) IsInt32() bool {
	return obj.isInt32
}

// IsInt64 returns true if the type is int64
func (obj *typ) IsInt64() bool {
	return obj.isInt64
}

// IsFloat32 returns true if the type is float32
func (obj *typ) IsFloat32() bool {
	return obj.isFloat32
}

// IsFloat64 returns true if the type is float64
func (obj *typ) IsFloat64() bool {
	return obj.isFloat64
}

// IsUint8 returns true if the type is uint8
func (obj *typ) IsUint8() bool {
	return obj.isUint8
}

// IsUint16 returns true if the type is uint16
func (obj *typ) IsUint16() bool {
	return obj.isUint16
}

// IsUint32 returns true if the type is uint32
func (obj *typ) IsUint32() bool {
	return obj.isUint32
}

// IsUint64 returns true if the type is uint64
func (obj *typ) IsUint64() bool {
	return obj.isUint64
}

// IsString returns true if the type is char
func (obj *typ) IsString() bool {
	return obj.isString
}

// String returns the type as string
func (obj *typ) String() string {
	if obj.IsNil() {
		return "nil"
	}

	if obj.IsBool() {
		return "bool"
	}

	if obj.IsInt8() {
		return "int8"
	}

	if obj.IsInt16() {
		return "int16"
	}

	if obj.IsInt32() {
		return "int32"
	}

	if obj.IsInt64() {
		return "int64"
	}

	if obj.IsUint8() {
		return "uint8"
	}

	if obj.IsUint16() {
		return "uint16"
	}

	if obj.IsUint32() {
		return "uint32"
	}

	if obj.IsUint64() {
		return "uint64"
	}

	if obj.IsFloat32() {
		return "float32"
	}

	if obj.IsFloat64() {
		return "float64"
	}

	return "string"
}
