package opengl

type cameraProjection struct {
	variable    string
	fov         float32
	aspectRatio float32
	near        float32
	far         float32
}

func createCameraProjection(
	variable string,
	fov float32,
	aspectRatio float32,
	near float32,
	far float32,
) CameraProjection {
	out := cameraProjection{
		variable:    variable,
		fov:         fov,
		aspectRatio: aspectRatio,
		near:        near,
		far:         far,
	}

	return &out
}

// Variable returns the variable
func (obj *cameraProjection) Variable() string {
	return obj.variable
}

// FieldOfView returns the field of view
func (obj *cameraProjection) FieldOfView() float32 {
	return obj.fov
}

// AspectRation returns the aspect ratio
func (obj *cameraProjection) AspectRation() float32 {
	return obj.aspectRatio
}

// Near returns the near
func (obj *cameraProjection) Near() float32 {
	return obj.near
}

// Far returns the far
func (obj *cameraProjection) Far() float32 {
	return obj.far
}
