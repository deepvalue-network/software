package opengl

import "github.com/go-gl/mathgl/mgl32"

type orientation struct {
	angle    float32
	dir      mgl32.Vec3
	variable string
}

func createOrientation(
	angle float32,
	dir mgl32.Vec3,
	variable string,
) Orientation {
	out := orientation{
		angle:    angle,
		dir:      dir,
		variable: variable,
	}

	return &out
}

// Angle returns the angle
func (obj *orientation) Angle() float32 {
	return obj.angle
}

// Direction returns the direction
func (obj *orientation) Direction() mgl32.Vec3 {
	return obj.dir
}

// Variable returns the variable
func (obj *orientation) Variable() string {
	return obj.variable
}

// Add adds an orientation
func (obj *orientation) Add(orientation Orientation) Orientation {
	vec := orientation.Direction()
	updatedDirection := obj.dir.Add(vec)
	updatedAngle := obj.angle + orientation.Angle()
	return createOrientation(updatedAngle, updatedDirection, obj.variable)
}
