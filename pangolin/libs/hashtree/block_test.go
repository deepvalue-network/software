package hashtree

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBlock_Success(t *testing.T) {
	blk, err := createBlockFromData([][]byte{
		[]byte("this is some data"),
		[]byte("this is other data"),
		[]byte("blah blah blah"),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil: %s", err.Error())
		return
	}

	js, err := json.Marshal(blk)
	if err != nil {
		t.Errorf("the error was expected to be nil: %s", err.Error())
		return
	}

	ins := new(block)
	err = json.Unmarshal(js, ins)
	if err != nil {
		t.Errorf("the error was expected to be nil: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(blk, ins) {
		t.Errorf("the json conversion is invalid")
		return
	}
}
