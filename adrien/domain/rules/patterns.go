package rules

import (
	"errors"
	"fmt"
)

type patterns struct {
	list []Pattern
	mp   map[string]Pattern
}

func createPatterns(
	list []Pattern,
	mp map[string]Pattern,
) Patterns {
	out := patterns{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns all patterns
func (obj *patterns) All() []Pattern {
	return obj.list
}

// Find finds a rule by name
func (obj *patterns) Find(name string) (Pattern, error) {
	if rule, ok := obj.mp[name]; ok {
		return rule, nil
	}

	str := fmt.Sprintf("the rule (%s) does not exists", name)
	return nil, errors.New(str)
}
