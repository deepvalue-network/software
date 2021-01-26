package parsers

type numericValue struct {
	isNegative bool
	intValue   *int
	floatValue *float64
}

func createNumericValueWithInt(isNegative bool, intValue *int) NumericValue {
	return createNumericValueInternally(isNegative, intValue, nil)
}

func createNumericValueWithFloat(isNegative bool, floatValue *float64) NumericValue {
	return createNumericValueInternally(isNegative, nil, floatValue)
}

func createNumericValueInternally(
	isNegative bool,
	intValue *int,
	floatValue *float64,
) NumericValue {
	out := numericValue{
		isNegative: isNegative,
		intValue:   intValue,
		floatValue: floatValue,
	}

	return &out
}

// IsNegative returns true if the value is negative, false otherwise
func (obj *numericValue) IsNegative() bool {
	return obj.isNegative
}

// IsInt returns true if the value is an int, false otherwise
func (obj *numericValue) IsInt() bool {
	return obj.intValue != nil
}

// Int returns the intValue, if any
func (obj *numericValue) Int() *int {
	return obj.intValue
}

// IsFloat returns true if the value is a float, false otherwise
func (obj *numericValue) IsFloat() bool {
	return obj.floatValue != nil
}

// Float returns the floatValue, if any
func (obj *numericValue) Float() *float64 {
	return obj.floatValue
}
