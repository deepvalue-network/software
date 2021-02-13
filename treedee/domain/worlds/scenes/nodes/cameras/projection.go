package cameras

type projection struct {
	variable string
	fov      float32
	aspect   float32
	near     float32
	far      float32
}

func createProjection(
	variable string,
	fov float32,
	aspect float32,
	near float32,
	far float32,
) Projection {
	out := projection{
		variable: variable,
		fov:      fov,
		aspect:   aspect,
		near:     near,
		far:      far,
	}

	return &out
}

// Variable returns the variable
func (obj *projection) Variable() string {
	return obj.variable
}

// FieldOfView returns the FOV
func (obj *projection) FieldOfView() float32 {
	return obj.fov
}

// AspectRatio returns the aspect ratio
func (obj *projection) AspectRatio() float32 {
	return obj.aspect
}

// Near returns the near distance
func (obj *projection) Near() float32 {
	return obj.near
}

// Far returns the far distance
func (obj *projection) Far() float32 {
	return obj.far
}
