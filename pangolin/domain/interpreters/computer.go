package interpreters

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type computer struct {
	builder computable.Builder
}

func createComputer(builder computable.Builder) Computer {
	out := computer{
		builder: builder,
	}

	return &out
}

// Add adds two values and return the result
func (app *computer) Add(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'add' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := int8(*firstVal + *secondVal)
		return app.builder.Create().WithInt8(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := int16(*firstVal + *secondVal)
		return app.builder.Create().WithInt16(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := int32(*firstVal + *secondVal)
		return app.builder.Create().WithInt32(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := int64(*firstVal + *secondVal)
		return app.builder.Create().WithInt64(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := uint8(*firstVal + *secondVal)
		return app.builder.Create().WithUint8(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := uint16(*firstVal + *secondVal)
		return app.builder.Create().WithUint16(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := uint32(*firstVal + *secondVal)
		return app.builder.Create().WithUint32(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := uint64(*firstVal + *secondVal)
		return app.builder.Create().WithUint64(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := float32(*firstVal + *secondVal)
		return app.builder.Create().WithFloat32(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := float64(*firstVal + *secondVal)
		return app.builder.Create().WithFloat64(val).Now()
	}

	return nil, errors.New("the 'add' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64")
}

// Substract substracts two values and return the result
func (app *computer) Substract(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'substract' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := int8(*firstVal - *secondVal)
		return app.builder.Create().WithInt8(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := int16(*firstVal - *secondVal)
		return app.builder.Create().WithInt16(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := int32(*firstVal - *secondVal)
		return app.builder.Create().WithInt32(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := int64(*firstVal - *secondVal)
		return app.builder.Create().WithInt64(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := uint8(*firstVal - *secondVal)
		return app.builder.Create().WithUint8(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := uint16(*firstVal - *secondVal)
		return app.builder.Create().WithUint16(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := uint32(*firstVal - *secondVal)
		return app.builder.Create().WithUint32(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := uint64(*firstVal - *secondVal)
		return app.builder.Create().WithUint64(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := float32(*firstVal - *secondVal)
		return app.builder.Create().WithFloat32(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := float64(*firstVal - *secondVal)
		return app.builder.Create().WithFloat64(val).Now()
	}

	return nil, errors.New("the 'substract' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64")
}

// Multiply multiplies two values and return the result
func (app *computer) Multiply(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'multiply' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := int8(*firstVal * *secondVal)
		return app.builder.Create().WithInt8(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := int16(*firstVal * *secondVal)
		return app.builder.Create().WithInt16(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := int32(*firstVal * *secondVal)
		return app.builder.Create().WithInt32(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := int64(*firstVal * *secondVal)
		return app.builder.Create().WithInt64(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := uint8(*firstVal * *secondVal)
		return app.builder.Create().WithUint8(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := uint16(*firstVal * *secondVal)
		return app.builder.Create().WithUint16(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := uint32(*firstVal * *secondVal)
		return app.builder.Create().WithUint32(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := uint64(*firstVal * *secondVal)
		return app.builder.Create().WithUint64(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := float32(*firstVal * *secondVal)
		return app.builder.Create().WithFloat32(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := float64(*firstVal * *secondVal)
		return app.builder.Create().WithFloat64(val).Now()
	}

	return nil, errors.New("the 'multiply' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64")
}

// Divide divides two values and return the result and the remaining
func (app *computer) Divide(first computable.Value, second computable.Value) (computable.Value, computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, nil, errors.New("the 'divide' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		result := int8(*firstVal % *secondVal)
		remaining := int8(*firstVal - result)
		res, err := app.builder.Create().WithInt8(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithInt8(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		result := int16(*firstVal % *secondVal)
		remaining := int16(*firstVal - result)
		res, err := app.builder.Create().WithInt16(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithInt16(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		result := int32(*firstVal % *secondVal)
		remaining := int32(*firstVal - result)
		res, err := app.builder.Create().WithInt32(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithInt32(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		result := int64(*firstVal % *secondVal)
		remaining := int64(*firstVal - result)
		res, err := app.builder.Create().WithInt64(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithInt64(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		result := uint8(*firstVal % *secondVal)
		remaining := uint8(*firstVal - result)
		res, err := app.builder.Create().WithUint8(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithUint8(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		result := uint16(*firstVal % *secondVal)
		remaining := uint16(*firstVal - result)
		res, err := app.builder.Create().WithUint16(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithUint16(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		result := uint32(*firstVal % *secondVal)
		remaining := uint32(*firstVal - result)
		res, err := app.builder.Create().WithUint32(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithUint32(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		result := uint64(*firstVal % *secondVal)
		remaining := uint64(*firstVal - result)
		res, err := app.builder.Create().WithUint64(result).Now()
		if err != nil {
			return nil, nil, err
		}

		rem, err := app.builder.Create().WithUint64(remaining).Now()
		if err != nil {
			return nil, nil, err
		}

		return res, rem, nil
	}

	return nil, nil, errors.New("the 'divide' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64")
}

// IsLessThan compares the two values and return the result
func (app *computer) IsLessThan(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'lessThan' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := *firstVal < *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	return nil, errors.New("the 'less than' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64")
}

// IsEqual compares the two values and return the result
func (app *computer) IsEqual(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'equal' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsNil() {
		secondIsNil := second.IsNil()
		return app.builder.Create().WithBool(secondIsNil).Now()
	}

	if first.IsBool() {
		firstVal := first.Bool()
		secondVal := second.Bool()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsString() {
		firstVal := first.String()
		secondVal := second.String()
		val := *firstVal == *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	return nil, errors.New("the 'equal' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64, bool, string, nil")
}

// IsNotEqual compares the two values and return the result
func (app *computer) IsNotEqual(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'not equal' operation can't execute because the variable types mismatch")
	}

	if first.IsIntHeight() {
		firstVal := first.IntHeight()
		secondVal := second.IntHeight()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixteen() {
		firstVal := first.IntSixteen()
		secondVal := second.IntSixteen()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntThirtyTwo() {
		firstVal := first.IntThirtyTwo()
		secondVal := second.IntThirtyTwo()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsIntSixtyFour() {
		firstVal := first.IntSixtyFour()
		secondVal := second.IntSixtyFour()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintHeight() {
		firstVal := first.UintHeight()
		secondVal := second.UintHeight()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixteen() {
		firstVal := first.UintSixteen()
		secondVal := second.UintSixteen()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintThirtyTwo() {
		firstVal := first.UintThirtyTwo()
		secondVal := second.UintThirtyTwo()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsUintSixtyFour() {
		firstVal := first.UintSixtyFour()
		secondVal := second.UintSixtyFour()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatThirtyTwo() {
		firstVal := first.FloatThirtyTwo()
		secondVal := second.FloatThirtyTwo()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsFloatSixtyFour() {
		firstVal := first.FloatSixtyFour()
		secondVal := second.FloatSixtyFour()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsNil() {
		secondIsNil := second.IsNil()
		return app.builder.Create().WithBool(!secondIsNil).Now()
	}

	if first.IsBool() {
		firstVal := first.Bool()
		secondVal := second.Bool()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	if first.IsString() {
		firstVal := first.String()
		secondVal := second.String()
		val := *firstVal != *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	return nil, errors.New("the 'not equal' operation can't execute on these types: int8, in16, int32, int64, uint8, uint16, uint32, uint64, float32, float64, bool, string, nil")
}

// And compares the two values and return the result
func (app *computer) And(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'and' operation can't execute because the variable types mismatch")
	}

	if first.IsBool() {
		firstVal := first.Bool()
		secondVal := second.Bool()
		val := *firstVal && *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	return nil, errors.New("the 'and' operation can't execute on these types: bool")
}

// Or compares the two values and return the result
func (app *computer) Or(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'or' operation can't execute because the variable types mismatch")
	}

	if first.IsBool() {
		firstVal := first.Bool()
		secondVal := second.Bool()
		val := *firstVal || *secondVal
		return app.builder.Create().WithBool(val).Now()
	}

	return nil, errors.New("the 'or' operation can't execute on these types: bool")
}

// Concat concats the two values together and return the result
func (app *computer) Concat(first computable.Value, second computable.Value) (computable.Value, error) {
	if !app.typeMatch(first, second) {
		return nil, errors.New("the 'concat' operation can't execute because the variable types mismatch")
	}

	if first.IsString() {
		firstVal := first.String()
		secondVal := second.String()
		val := fmt.Sprintf("%s%s", *firstVal, *secondVal)
		return app.builder.Create().WithString(val).Now()
	}

	return nil, errors.New("the 'concat' operation can't execute on these types: string")
}

// Match takes a pattern and a value and returns the results
func (app *computer) Match(pattern computable.Value, value computable.Value) ([]computable.Value, computable.Value, error) {
	if !pattern.IsString() {
		return nil, nil, errors.New("the 'match' operation requires the pattern to be of type: string")
	}

	if !value.IsString() {
		return nil, nil, errors.New("the 'match' operation requires the value to be of type: string")
	}

	patternStr := pattern.String()
	valueStr := value.String()

	re, err := regexp.Compile(*patternStr)
	if err != nil {
		return nil, nil, err
	}

	out := []computable.Value{}
	results := re.FindAllString(*valueStr, -1)
	for _, oneResult := range results {
		val, err := app.builder.Create().WithString(oneResult).Now()
		if err != nil {
			return nil, nil, err
		}

		out = append(out, val)
	}

	am := uint32(len(out))
	amount, err := app.builder.Create().WithUint32(am).Now()
	if err != nil {
		return nil, nil, err
	}

	return out, amount, nil
}

func (app *computer) typeMatch(first computable.Value, second computable.Value) bool {
	if first.IsNil() {
		return second.IsNil()
	}

	if first.IsBool() {
		return second.IsBool()
	}

	if first.IsString() {
		return second.IsString()
	}

	if first.IsIntHeight() {
		return second.IsIntHeight()
	}

	if first.IsIntSixteen() {
		return second.IsIntSixteen()
	}

	if first.IsIntThirtyTwo() {
		return second.IsIntThirtyTwo()
	}

	if first.IsIntSixtyFour() {
		return second.IsIntSixtyFour()
	}

	if first.IsUintHeight() {
		return second.IsUintHeight()
	}

	if first.IsUintSixteen() {
		return second.IsUintSixteen()
	}

	if first.IsUintThirtyTwo() {
		return second.IsUintThirtyTwo()
	}

	if first.IsUintSixtyFour() {
		return second.IsUintSixtyFour()
	}

	if first.IsFloatThirtyTwo() {
		return second.IsFloatThirtyTwo()
	}

	if first.IsFloatSixtyFour() {
		return second.IsFloatSixtyFour()
	}

	if first.IsStackFrame() {
		return second.IsStackFrame()
	}

	// frame:
	return second.IsFrame()
}
