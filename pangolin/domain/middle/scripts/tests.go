package scripts

import (
	"errors"
	"fmt"
)

type tests struct {
	list []Test
	mp   map[string]Test
}

func createTests(
	list []Test,
	mp map[string]Test,
) Tests {
	out := tests{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the tests
func (obj *tests) All() []Test {
	return obj.list
}

// FetchByName fetches a test by name
func (obj *tests) FetchByName(name string) (Test, error) {
	if test, ok := obj.mp[name]; ok {
		return test, nil
	}

	str := fmt.Sprintf("the test (name: %s) does not exists", name)
	return nil, errors.New(str)
}
