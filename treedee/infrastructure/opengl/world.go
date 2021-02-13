package opengl

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/application/windows"
)

type world struct {
	id                *uuid.UUID
	winApp            windows.Application
	logic             WorldLogicFn
	currentSceneIndex uint
	scenes            []Scene
}

func createWorld(
	id *uuid.UUID,
	winApp windows.Application,
	logic WorldLogicFn,
	currentSceneIndex uint,
	scenes []Scene,
) World {
	out := world{
		id:                id,
		winApp:            winApp,
		logic:             logic,
		currentSceneIndex: currentSceneIndex,
		scenes:            scenes,
	}

	return &out
}

// ID returns the id
func (obj *world) ID() *uuid.UUID {
	return obj.id
}

// CurrentSceneIndex returns the currentSceneIndex
func (obj *world) CurrentSceneIndex() uint {
	return obj.currentSceneIndex
}

// Scenes returns the scenes
func (obj *world) Scenes() []Scene {
	return obj.scenes
}

// WindowApplication returns the window application
func (obj *world) WindowApplication() windows.Application {
	return obj.winApp
}

// Logic returns the logic func
func (obj *world) Logic() WorldLogicFn {
	return obj.logic
}

// Executes the world
func (obj *world) Execute() error {
	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.5, 0.2, 0.4, 1.0)

	// execute the window app:
	return obj.winApp.Execute(obj.updateFn)
}

func (obj *world) updateFn(prev time.Time, current time.Time) error {
	// update the logic:
	err := obj.logic(obj)
	if err != nil {
		return err
	}

	// sleep some time:
	time.Sleep(time.Duration(float64(time.Second) * 0.001))
	return nil
}
