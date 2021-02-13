package opengl

import "github.com/go-gl/mathgl/mgl32"

type cameraLookAt struct {
	variable string
	eye      mgl32.Vec3
	center   mgl32.Vec3
	up       mgl32.Vec3
}

func createCameraLookAt(
	variable string,
	eye mgl32.Vec3,
	center mgl32.Vec3,
	up mgl32.Vec3,
) CameraLookAt {
	out := cameraLookAt{
		variable: variable,
		eye:      eye,
		center:   center,
		up:       up,
	}

	return &out
}

// Variable returns the variable
func (obj *cameraLookAt) Variable() string {
	return obj.variable
}

// Eye returns the eye
func (obj *cameraLookAt) Eye() mgl32.Vec3 {
	return obj.eye
}

// Center returns the center
func (obj *cameraLookAt) Center() mgl32.Vec3 {
	return obj.center
}

// Up returns the up
func (obj *cameraLookAt) Up() mgl32.Vec3 {
	return obj.up
}
