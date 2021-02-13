package ints

import "fmt"

type rectangle struct {
	pos Vec2
	dim Vec2
}

func createRectangle(
	pos Vec2,
	dim Vec2,
) Rectangle {
	out := rectangle{
		pos: pos,
		dim: dim,
	}

	return &out
}

// Position returns the position
func (obj *rectangle) Position() Vec2 {
	return obj.pos
}

// Dimension returns the dimension
func (obj *rectangle) Dimension() Vec2 {
	return obj.dim
}

// String returns the string representation of a rectangle
func (obj *rectangle) String() string {
	return fmt.Sprintf("[pos: %s, dim: %s]", obj.pos.String(), obj.dim.String())
}
