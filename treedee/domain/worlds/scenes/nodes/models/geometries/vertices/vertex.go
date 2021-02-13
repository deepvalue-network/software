package vertices

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
)

type vertex struct {
	space fl32.Vec3
	tex   fl32.Vec2
}

func createVertex(
	space fl32.Vec3,
	tex fl32.Vec2,
) Vertex {
	out := vertex{
		space: space,
		tex:   tex,
	}

	return &out
}

// Space returns the space
func (obj *vertex) Space() fl32.Vec3 {
	return obj.space
}

// Texture returns the texture
func (obj *vertex) Texture() fl32.Vec2 {
	return obj.tex
}
