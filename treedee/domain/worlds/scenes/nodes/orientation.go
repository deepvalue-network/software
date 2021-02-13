package nodes

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
)

type orientation struct {
	angle     float32
	direction fl32.Vec3
	variable  string
}

func createOrientation(
	angle float32,
	direction fl32.Vec3,
	variable string,
) Orientation {
	out := orientation{
		angle:     angle,
		direction: direction,
		variable:  variable,
	}

	return &out
}

// Angle returns the angle
func (obj *orientation) Angle() float32 {
	return obj.angle
}

// Direction returns the direction
func (obj *orientation) Direction() fl32.Vec3 {
	return obj.direction
}

// Variable returns the variable
func (obj *orientation) Variable() string {
	return obj.variable
}
