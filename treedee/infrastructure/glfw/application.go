package glfw

import (
	"time"

	"github.com/deepvalue-network/software/treedee/application/windows"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type application struct {
	win *glfw.Window
}

func createApplication(
	win *glfw.Window,
) windows.Application {
	out := application{
		win: win,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(fn windows.UpdateFn) error {
	defer glfw.Terminate()
	prev := time.Now().UTC()
	for !app.win.ShouldClose() {
		// update:
		current := time.Now().UTC()
		err := fn(prev, current)
		if err != nil {
			return err
		}

		// Maintenance
		app.win.SwapBuffers()
		glfw.PollEvents()

		// update the time:
		prev = current
	}

	return nil
}
