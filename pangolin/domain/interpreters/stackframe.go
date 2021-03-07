package interpreters

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable/value/computable"
)

type stackFrame struct {
	frameBuilder FrameBuilder
	variables    map[string]computable.Value
	constants    map[string]computable.Value
	frames       []Frame
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
		frames: []Frame{
			frameBuilder.Create().WithVariables(variables).WithConstants(constants).Now(),
		},
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
	newFrame := app.frameBuilder.Create().WithConstants(app.constants).WithVariables(app.variables).Now()
	app.frames = append(app.frames, newFrame)
}

// Pop pops the last stacked frame as the current frame
func (app *stackFrame) Pop() error {
	if len(app.frames) <= 1 {
		return errors.New("the pop instruction cannot be executed because the frame stack is empty")
	}

	index := len(app.frames) - 1
	app.frames = app.frames[:index]
	return nil
}

// Current returns the current frame
func (app *stackFrame) Current() Frame {
	index := len(app.frames) - 1
	return app.frames[index]
}
