package glfw

import (
	"errors"
	"fmt"
	"runtime"

	application_windows "github.com/deepvalue-network/software/treedee/application/windows"
	"github.com/deepvalue-network/software/treedee/domain/windows"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type builder struct {
	win windows.Window
}

func createBuilder() application_windows.Builder {
	out := builder{
		win: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() application_windows.Builder {
	return createBuilder()
}

// WithWindow adds a window to the builder
func (app *builder) WithWindow(win windows.Window) application_windows.Builder {
	app.win = win
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (application_windows.Application, error) {
	if app.win == nil {
		return nil, errors.New("the window is mandatory in order to build a GLFW Application")
	}

	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		str := fmt.Sprintf("failed to initialize glfw: %s", err.Error())
		return nil, errors.New(str)
	}

	resizable := glfw.False
	if app.win.IsResizable() {
		resizable = glfw.False
	}

	glfw.WindowHint(glfw.Resizable, resizable)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	if app.win.IsFullscreen() {
		monitor := glfw.GetPrimaryMonitor()
		if monitor == nil {
			return nil, errors.New("there was an error while fetching the primary monitor")
		}

		mode := monitor.GetVideoMode()
		if mode == nil {
			return nil, errors.New("there was an error while fetching the video mode")
		}

		win, err := glfw.CreateWindow(mode.Width, mode.Height, app.win.Title(), monitor, nil)
		if err != nil {
			return nil, err
		}

		win.MakeContextCurrent()
		return createApplication(win), nil
	}

	win, err := glfw.CreateWindow(int(app.win.Width()), int(app.win.Height()), app.win.Title(), nil, nil)
	if err != nil {
		return nil, err
	}

	win.MakeContextCurrent()
	return createApplication(win), nil
}
