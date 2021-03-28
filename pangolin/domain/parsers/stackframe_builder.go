package parsers

import "errors"

type stackFrameBuilder struct {
	push Push
	pop  Pop
}

func createStackFrameBuilder() StackFrameBuilder {
	out := stackFrameBuilder{
		push: nil,
		pop:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *stackFrameBuilder) Create() StackFrameBuilder {
	return createStackFrameBuilder()
}

// WithPush adds a push to the builder
func (app *stackFrameBuilder) WithPush(push Push) StackFrameBuilder {
	app.push = push
	return app
}

// WithPop adds a pop to the builder
func (app *stackFrameBuilder) WithPop(pop Pop) StackFrameBuilder {
	app.pop = pop
	return app
}

// Now builds a new StackFrame instance
func (app *stackFrameBuilder) Now() (StackFrame, error) {
	if app.push != nil {
		return createStackFrameWithPush(app.push), nil
	}

	if app.pop != nil {
		return createStackFrameWithPop(app.pop), nil
	}

	return nil, errors.New("the StackFrame is invalid")
}
