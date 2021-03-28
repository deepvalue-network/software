package parsers

import "errors"

type jumpBuilder struct {
	label     string
	condition string
}

func createJumpBuilder() JumpBuilder {
	out := jumpBuilder{
		label:     "",
		condition: "",
	}

	return &out
}

// Create initializes the builder
func (obj *jumpBuilder) Create() JumpBuilder {
	return createJumpBuilder()
}

// WithLabel adds a label to the builder
func (obj *jumpBuilder) WithLabel(label string) JumpBuilder {
	obj.label = label
	return obj
}

// WithCondition adds a condition to the builder
func (obj *jumpBuilder) WithCondition(condition string) JumpBuilder {
	obj.condition = condition
	return obj
}

// Now builds a new Jump instance
func (obj *jumpBuilder) Now() (Jump, error) {
	if obj.label == "" {
		return nil, errors.New("the Label is mandatory in order to build a Jump instance")
	}

	if obj.condition != "" {
		return createJumpWithCondition(obj.label, obj.condition), nil
	}

	return createJump(obj.label), nil
}
