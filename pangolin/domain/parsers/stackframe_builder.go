package parsers

import "errors"

type stackFrameBuilder struct {
	assignment FrameAssignment
	push       Push
	pop        Pop
}

func createStackFrameBuilder() StackFrameBuilder {
	out := stackFrameBuilder{
		assignment: nil,
		push:       nil,
		pop:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *stackFrameBuilder) Create() StackFrameBuilder {
	return createStackFrameBuilder()
}

// WithAssignment adds an assignment to the builder
func (app *stackFrameBuilder) WithAssignment(assignment FrameAssignment) StackFrameBuilder {
	app.assignment = assignment
	return app
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
	if app.assignment != nil {
		return createStackFrameWithAssignment(app.assignment), nil
	}

	if app.push != nil {
		return createStackFrameWithPush(app.push), nil
	}

	if app.pop != nil {
		return createStackFrameWithPop(app.pop), nil
	}

	return nil, errors.New("the StackFrame is invalid")
}
