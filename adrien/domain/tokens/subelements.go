package tokens

import (
	"errors"
	"fmt"
)

type subElements struct {
	list []SubElement
	mp   map[string]SubElement
}

func createSubElements(
	list []SubElement,
	mp map[string]SubElement,
) SubElements {
	out := subElements{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the sub elements
func (obj *subElements) All() []SubElement {
	return obj.list
}

// Find returns the sub element by name
func (obj *subElements) Find(name string) (SubElement, error) {
	if el, ok := obj.mp[name]; ok {
		return el, nil
	}

	str := fmt.Sprintf("the subElement (name: %s) is not declared", name)
	return nil, errors.New(str)
}
