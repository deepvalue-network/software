package pixels

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/colors"
)

type pixel struct {
	color colors.Color
	alpha uint8
}

func createPixel(
	color colors.Color,
	alpha uint8,
) Pixel {
	out := pixel{
		color: color,
		alpha: alpha,
	}

	return &out
}

// Color returns the color
func (obj *pixel) Color() colors.Color {
	return obj.color
}

// Alpha returns the alpha
func (obj *pixel) Alpha() uint8 {
	return obj.alpha
}
