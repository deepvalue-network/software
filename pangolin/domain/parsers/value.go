package parsers

type value struct {
	isNil    bool
	numeric  NumericValue
	bl       *bool
	strValue string
}

func createValueWithNil() Value {
	return createValueInternally(true, nil, nil, "")
}

func createValueWithNumeric(numeric NumericValue) Value {
	return createValueInternally(false, numeric, nil, "")
}

func createValueWithBool(bl *bool) Value {
	return createValueInternally(false, nil, bl, "")
}

func createValueWithString(strValue string) Value {
	return createValueInternally(false, nil, nil, strValue)
}

func createValueInternally(
	isNil bool,
	numeric NumericValue,
	bl *bool,
	strValue string,
) Value {
	out := value{
		isNil:    isNil,
		numeric:  numeric,
		bl:       bl,
		strValue: strValue,
	}

	return &out
}

// IsNil returns true if the value is nil, false otherwise
func (obj *value) IsNil() bool {
	return obj.isNil
}

// IsNumeric returns true if the value is numeric, false otherwise
func (obj *value) IsNumeric() bool {
	return obj.numeric != nil
}

// Numeric returns the numeric value, if any
func (obj *value) Numeric() NumericValue {
	return obj.numeric
}

// IsBool returns true if the value is bool, false otherwise
func (obj *value) IsBool() bool {
	return obj.bl != nil
}

// Bool returns the bool value, if any
func (obj *value) Bool() *bool {
	return obj.bl
}

// IsString returns true if the value is a string, false otherwise
func (obj *value) IsString() bool {
	return obj.strValue != ""
}

// String returns the string value, if any
func (obj *value) String() string {
	return obj.strValue
}
