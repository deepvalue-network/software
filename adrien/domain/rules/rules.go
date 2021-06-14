package rules

import (
	"errors"
	"fmt"
)

type rules struct {
	list []Rule
	mp   map[string]Rule
}

func createRules(
	list []Rule,
	mp map[string]Rule,
) Rules {
	out := rules{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns all rules
func (obj *rules) All() []Rule {
	return obj.list
}

// Find finds a rule by name
func (obj *rules) Find(name string) (Rule, error) {
	if rule, ok := obj.mp[name]; ok {
		return rule, nil
	}

	str := fmt.Sprintf("the rule (%s) does not exists", name)
	return nil, errors.New(str)
}
