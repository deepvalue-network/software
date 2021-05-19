package stackframes

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

type stackFrame struct {
	frameBuilder FrameBuilder
	variables    map[string]computable.Value
	frames       []Frame
	index        int
}

func createStackFrame(
	frameBuilder FrameBuilder,
	variables map[string]computable.Value,
) StackFrame {
	out := stackFrame{
		frameBuilder: frameBuilder,
		variables:    variables,
		index:        0,
		frames: []Frame{
			frameBuilder.Create().WithVariables(variables).Now(),
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
	newFrame := app.frameBuilder.Create().WithVariables(app.variables).Now()
	app.frames = append(app.frames, newFrame)
	app.index = len(app.frames) - 1
}

// Pop pops the last stacked frame as the current frame
func (app *stackFrame) Pop() error {
	if app.index <= 0 {
		return errors.New("the pop instruction cannot be executed because the frame stack is empty")
	}

	app.frames = app.frames[:app.index]
	app.index = len(app.frames) - 1
	return nil
}

// Index returns the stackframe index
func (app *stackFrame) Index() int {
	return app.index
}

// Skip skips the stackFrame to the specified index
func (app *stackFrame) Skip(index int) error {
	if index >= len(app.frames) {
		str := fmt.Sprintf("the stackFrame could not skip to index: %d because it only contains %d frames", index, len(app.frames))
		return errors.New(str)
	}

	if app.index < 0 {
		str := fmt.Sprintf("the smallest index to skip to is 0, %d provided", index)
		return errors.New(str)
	}

	app.index = index
	return nil
}

// Current returns the current frame
func (app *stackFrame) Current() Frame {
	return app.frames[app.index]
}
