package viewports

import "github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"

type viewport struct {
	rectangle ints.Rectangle
	variable  string
}

func createViewport(
	rectangle ints.Rectangle,
	variable string,
) Viewport {
	out := viewport{
		rectangle: rectangle,
		variable:  variable,
	}

	return &out
}

// Rectangle returns the rectangle
func (obj *viewport) Rectangle() ints.Rectangle {
	return obj.rectangle
}

// Variable returns the variable
func (obj *viewport) Variable() string {
	return obj.variable
}

// IsContained returns true if the viewport is contained within the given dimension, false otherwise
func (obj *viewport) IsContained(dim ints.Vec2) bool {
	vpos := obj.rectangle.Position()
	vdim := obj.rectangle.Dimension()
	width := vpos.X() + vdim.X()
	height := vpos.Y() + vdim.Y()
	return width < dim.X() && height < dim.Y()
}
