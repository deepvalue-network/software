package targets

import (
	"errors"
	"fmt"
)

type targets struct {
	list []Target
	mp   map[string]Target
}

func createTargets(
	list []Target,
	mp map[string]Target,
) Targets {
	out := targets{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the targets
func (obj *targets) All() []Target {
	return obj.list
}

// Fetch fetches a target by name
func (obj *targets) Fetch(name string) (Target, error) {
	if target, ok := obj.mp[name]; ok {
		return target, nil
	}

	str := fmt.Sprintf("the target (name: %s) does not exists", name)
	return nil, errors.New(str)
}
