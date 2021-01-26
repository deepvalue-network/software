package schemas

type typ struct {
	isString  bool
	isInt     bool
	isFloat32 bool
	isFloat64 bool
	isData    bool
}

func createTypeWithString() Type {
	return createTypeInternally(true, false, false, false, false)
}

func createTypeWithInt() Type {
	return createTypeInternally(false, true, false, false, false)
}

func createTypeWithFloat32() Type {
	return createTypeInternally(false, false, true, false, false)
}

func createTypeWithFloat64() Type {
	return createTypeInternally(false, false, false, true, false)
}

func createTypeWithData() Type {
	return createTypeInternally(false, false, false, false, true)
}

func createTypeInternally(
	isString bool,
	isInt bool,
	isFloat32 bool,
	isFloat64 bool,
	isData bool,
) Type {
	out := typ{
		isString:  isString,
		isInt:     isInt,
		isFloat32: isFloat32,
		isFloat64: isFloat64,
		isData:    isData,
	}

	return &out
}

// IsString returns true if there is a string, false otherwise
func (obj *typ) IsString() bool {
	return obj.isString
}

// IsInt returns true if there is an int, false otherwise
func (obj *typ) IsInt() bool {
	return obj.isInt
}

// IsFloat32 returns true if there is a float32, false otherwise
func (obj *typ) IsFloat32() bool {
	return obj.isFloat32
}

// IsFloat64 returns true if there is a float64, false otherwise
func (obj *typ) IsFloat64() bool {
	return obj.isFloat64
}

// IsData returns true if there is data, false otherwise
func (obj *typ) IsData() bool {
	return obj.isData
}
