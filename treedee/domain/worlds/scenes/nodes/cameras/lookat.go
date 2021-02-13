package cameras

import "github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"

type lookAt struct {
	variable string
	eye      fl32.Vec3
	center   fl32.Vec3
	up       fl32.Vec3
}

func createLookAt(
	variable string,
	eye fl32.Vec3,
	center fl32.Vec3,
	up fl32.Vec3,
) LookAt {
	out := lookAt{
		variable: variable,
		eye:      eye,
		center:   center,
		up:       up,
	}

	return &out
}

// Variable returns the variable
func (obj *lookAt) Variable() string {
	return obj.variable
}

// Eye returns the eye of the camera
func (obj *lookAt) Eye() fl32.Vec3 {
	return obj.eye
}

// Center returns the center of the camera
func (obj *lookAt) Center() fl32.Vec3 {
	return obj.center
}

// Up returns the up of the camera
func (obj *lookAt) Up() fl32.Vec3 {
	return obj.up
}
