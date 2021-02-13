package colors

type color struct {
	red   uint8
	green uint8
	blue  uint8
}

func createColor(
	red uint8,
	green uint8,
	blue uint8,
) Color {
	out := color{
		red:   red,
		green: green,
		blue:  blue,
	}

	return &out
}

// Red returns the red value
func (obj *color) Red() uint8 {
	return obj.red
}

// Green returns the green value
func (obj *color) Green() uint8 {
	return obj.green
}

// Blue returns the blue value
func (obj *color) Blue() uint8 {
	return obj.blue
}
