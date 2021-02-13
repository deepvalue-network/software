package opengl

type alpha struct {
	value    float32
	variable string
}

func createAlpha(
	value float32,
	variable string,
) Alpha {
	out := alpha{
		value:    value,
		variable: variable,
	}

	return &out
}

// Value returns the value
func (obj *alpha) Value() float32 {
	return obj.value
}

// Variable returns the variable
func (obj *alpha) Variable() string {
	return obj.variable
}
