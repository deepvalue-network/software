package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/displays"
)

type hudDisplayBuilder struct {
	programBuilder  ProgramBuilder
	cameraBuilder   CameraBuilder
	materialBuilder MaterialBuilder
	display         displays.Display
}

func createHudDisplayBuilder(
	programBuilder ProgramBuilder,
	cameraBuilder CameraBuilder,
	materialBuilder MaterialBuilder,
) HudDisplayBuilder {
	out := hudDisplayBuilder{
		programBuilder:  programBuilder,
		cameraBuilder:   cameraBuilder,
		materialBuilder: materialBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *hudDisplayBuilder) Create() HudDisplayBuilder {
	return createHudDisplayBuilder(app.programBuilder, app.cameraBuilder, app.materialBuilder)
}

// WithDisplay adds a display to the builder
func (app *hudDisplayBuilder) WithDisplay(display displays.Display) HudDisplayBuilder {
	app.display = display
	return app
}

// Now builds a new HudDisplay instance
func (app *hudDisplayBuilder) Now() (HudDisplay, error) {
	if app.display == nil {
		return nil, errors.New("the display is mandatory in order to build a HudDisplay instance")
	}

	domainCamera := app.display.Camera()
	cam, err := app.cameraBuilder.Create().WithCamera(domainCamera).Now()
	if err != nil {
		return nil, err
	}

	id := app.display.ID()
	index := app.display.Index()
	viewport := app.display.Viewport()
	if app.display.HasMaterial() {
		domainMat := app.display.Material()
		mat, err := app.materialBuilder.Create().WithMaterial(domainMat).Now()
		if err != nil {
			return nil, err
		}

		return createHudDisplayWithMaterial(id, index, viewport, cam, mat), nil
	}

	return createHudDisplay(id, index, viewport, cam), nil
}
