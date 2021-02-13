package opengl

import "github.com/go-gl/mathgl/mgl32"

type position struct {
	vec      mgl32.Vec3
	variable string
}

func createPosition(
	vec mgl32.Vec3,
	variable string,
) Position {
	out := position{
		vec:      vec,
		variable: variable,
	}

	return &out
}

// Vector returns the position vector
func (obj *position) Vector() mgl32.Vec3 {
	return obj.vec
}

// Variable returns the variable
func (obj *position) Variable() string {
	return obj.variable
}

// Add adds a position
func (obj *position) Add(pos Position) Position {
	vec := pos.Vector()
	updatedPosVec := obj.vec.Add(vec)
	return createPosition(updatedPosVec, obj.variable)
}
