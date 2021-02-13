package opengl

import (
	"errors"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
)

type cameraBuilder struct {
	cam cameras.Camera
}

func createCameraBuilder() CameraBuilder {
	out := cameraBuilder{
		cam: nil,
	}

	return &out
}

// Create initializes the builder
func (app *cameraBuilder) Create() CameraBuilder {
	return createCameraBuilder()
}

// WithCamera adds a camera to the builder
func (app *cameraBuilder) WithCamera(cam cameras.Camera) CameraBuilder {
	app.cam = cam
	return app
}

// Now builds a new Camera instance
func (app *cameraBuilder) Now() (Camera, error) {
	if app.cam == nil {
		return nil, errors.New("the camera is mandatory in order to build a Camera instance")
	}

	domainProjection := app.cam.Projection()
	projVariable := domainProjection.Variable()
	fov := domainProjection.FieldOfView()
	aspectRatio := domainProjection.AspectRatio()
	near := domainProjection.Near()
	far := domainProjection.Far()
	projection := createCameraProjection(projVariable, fov, aspectRatio, near, far)

	domainLookAt := app.cam.LookAt()
	lookAtVariable := domainLookAt.Variable()
	eye := domainLookAt.Eye()
	center := domainLookAt.Center()
	up := domainLookAt.Up()
	lookAt := createCameraLookAt(
		lookAtVariable,
		mgl32.Vec3{
			eye.X(),
			eye.Y(),
			eye.Z(),
		},
		mgl32.Vec3{
			center.X(),
			center.Y(),
			center.Z(),
		},
		mgl32.Vec3{
			up.X(),
			up.Y(),
			up.Z(),
		},
	)

	id := app.cam.ID()
	index := app.cam.Index()
	return createCamera(id, index, projection, lookAt), nil
}
