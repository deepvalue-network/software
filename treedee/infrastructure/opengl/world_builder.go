package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/application/windows"
	domain_window "github.com/deepvalue-network/software/treedee/domain/windows"
	"github.com/deepvalue-network/software/treedee/domain/worlds"
)

type worldBuilder struct {
	winAppBuilder windows.Builder
	sceneBuilder  SceneBuilder
	win           domain_window.Window
	logic         WorldLogicFn
	world         worlds.World
}

func createWorldBuilder(
	winAppBuilder windows.Builder,
	sceneBuilder SceneBuilder,
) WorldBuilder {
	out := worldBuilder{
		winAppBuilder: winAppBuilder,
		sceneBuilder:  sceneBuilder,
		logic:         nil,
		world:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *worldBuilder) Create() WorldBuilder {
	return createWorldBuilder(app.winAppBuilder, app.sceneBuilder)
}

// WithWorld adds a world to the builder
func (app *worldBuilder) WithWorld(world worlds.World) WorldBuilder {
	app.world = world
	return app
}

// WithWindow adds a window to the builder
func (app *worldBuilder) WithWindow(win domain_window.Window) WorldBuilder {
	app.win = win
	return app
}

// WithLogic adds a logic to the builder
func (app *worldBuilder) WithLogic(logic WorldLogicFn) WorldBuilder {
	app.logic = logic
	return app
}

// Now builds a new World instance
func (app *worldBuilder) Now() (World, error) {
	if app.world == nil {
		return nil, errors.New("the world is mandatory in order to build a World instance")
	}

	if app.win == nil {
		return nil, errors.New("the window is mandatory in order to build a World instance")
	}

	if app.logic == nil {
		return nil, errors.New("the logic is mandatory in order to build a World instance")
	}

	win, err := app.winAppBuilder.Create().WithWindow(app.win).Now()
	if err != nil {
		return nil, err
	}

	scenes := []Scene{}
	domainScenes := app.world.Scenes()
	for _, oneScene := range domainScenes {
		scene, err := app.sceneBuilder.Create().WithScene(oneScene).Now()
		if err != nil {
			return nil, err
		}

		scenes = append(scenes, scene)
	}

	id := app.world.ID()
	currentSceneIndex := app.world.CurrentSceneIndex()
	return createWorld(id, win, app.logic, currentSceneIndex, scenes), nil
}
