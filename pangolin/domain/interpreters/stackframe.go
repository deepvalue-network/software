package interpreters

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value/computable"
)

type stackFrame struct {
	frameBuilder FrameBuilder
	variables    map[string]computable.Value
	constants    map[string]computable.Value
	frames       []Frame
	current      Frame
}

func createStackFrame(
	frameBuilder FrameBuilder,
	variables map[string]computable.Value,
	constants map[string]computable.Value,
) StackFrame {
	out := stackFrame{
		frameBuilder: frameBuilder,
		variables:    variables,
		constants:    constants,
		frames:       []Frame{},
		current:      frameBuilder.Create().WithVariables(variables).WithConstants(constants).Now(),
	}

	return &out
}

// PushTo pushes the current frame to a variableName's frame
func (app *stackFrame) PushTo(name string) error {
	fmt.Println("finish PushTo inside pangolin interpreter stackframe PushTo method")
	return nil
}

// Push pushes the current frame to the stack
func (app *stackFrame) Push() {
	app.frames = append(app.frames, app.current)
	app.current = app.frameBuilder.Create().WithConstants(app.constants).WithVariables(app.variables).Now()
}

// Pop pops the last stacked frame as the current frame
func (app *stackFrame) Pop() error {
	if len(app.frames) <= 0 {
		return errors.New("the pop instruction cannot be executed because the frame stack is empty")
	}

	index := len(app.frames) - 1
	app.current = app.frames[index]
	app.frames = app.frames[:index]
	return nil
}

// Current returns the current frame
func (app *stackFrame) Current() Frame {
	return app.current
}
