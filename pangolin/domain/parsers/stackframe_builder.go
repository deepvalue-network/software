package parsers

import "errors"

type stackFrameBuilder struct {
	isPush bool
	isPop  bool
}

func createStackFrameBuilder() StackFrameBuilder {
	out := stackFrameBuilder{
		isPush: false,
		isPop:  false,
	}

	return &out
}

// Create initializes the builder
func (app *stackFrameBuilder) Create() StackFrameBuilder {
	return createStackFrameBuilder()
}

// IsPush flags the builder as push
func (app *stackFrameBuilder) IsPush() StackFrameBuilder {
	app.isPush = true
	return app
}

// IsPop flags the builder as pop
func (app *stackFrameBuilder) IsPop() StackFrameBuilder {
	app.isPop = true
	return app
}

// Now builds a new StackFrame instance
func (app *stackFrameBuilder) Now() (StackFrame, error) {
	if app.isPush {
		return createStackFrameWithPush(), nil
	}

	if app.isPop {
		return createStackFrameWithPop(), nil
	}

	return nil, errors.New("the StackFrame is invalid")
}
