package rules

type rule struct {
	name     string
	code     string
	elements []Element
}

func createRule(
	name string,
	code string,
	elements []Element,
) Rule {
	out := rule{
		name:     name,
		code:     code,
		elements: elements,
	}

	return &out
}

// Name returns the name
func (obj *rule) Name() string {
	return obj.name
}

// Code returns the code
func (obj *rule) Code() string {
	return obj.code
}

// Elements returns the elements
func (obj *rule) Elements() []Element {
	return obj.elements
}
