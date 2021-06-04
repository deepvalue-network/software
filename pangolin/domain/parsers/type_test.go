package parsers

import (
	"testing"
)

func Test_type_bool_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/bool.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsBool() {
		t.Errorf("the type was expected to be bool")
		return
	}
}

func Test_type_int8_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/int8.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsInt8() {
		t.Errorf("the type was expected to be int8")
		return
	}
}

func Test_type_int16_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/int16.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsInt16() {
		t.Errorf("the type was expected to be int16")
		return
	}
}

func Test_type_int32_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/int32.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsInt32() {
		t.Errorf("the type was expected to be int32")
		return
	}
}

func Test_type_int64_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/int64.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsInt64() {
		t.Errorf("the type was expected to be int64")
		return
	}
}

func Test_type_float32_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/float32.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsFloat32() {
		t.Errorf("the type was expected to be float32")
		return
	}
}

func Test_type_float64_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/float64.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsFloat64() {
		t.Errorf("the type was expected to be float64")
		return
	}
}

func Test_type_uint8_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/uint8.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsUint8() {
		t.Errorf("the type was expected to be uint8")
		return
	}
}

func Test_type_uint16_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/uint16.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsUint16() {
		t.Errorf("the type was expected to be uint16")
		return
	}
}

func Test_type_uint32_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/uint32.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsUint32() {
		t.Errorf("the type was expected to be uint32")
		return
	}
}

func Test_type_uint64_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/uint64.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsUint64() {
		t.Errorf("the type was expected to be uint64")
		return
	}
}

func Test_type_string_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/string.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsString() {
		t.Errorf("the type was expected to be a string")
		return
	}
}

func Test_type_stackframe_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("type", grammarFile)

	file := "./tests/codes/type/stackframe.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	typ := ins.(Type)
	if !typ.IsStackFrame() {
		t.Errorf("the type was expected to be a stackframe")
		return
	}
}
