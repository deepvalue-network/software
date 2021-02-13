package nodes

import "github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"

type position struct {
	vec      fl32.Vec3
	variable string
}

func createPosition(
	vec fl32.Vec3,
	variable string,
) Position {
	out := position{
		vec:      vec,
		variable: variable,
	}

	return &out
}

// Vector returns the vector
func (obj *position) Vector() fl32.Vec3 {
	return obj.vec
}

// Variable returns the variable
func (obj *position) Variable() string {
	return obj.variable
}
