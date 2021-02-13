package alphas

type alpha struct {
	alpha    uint8
	variable string
}

func createAlpha(
	al uint8,
	variable string,
) Alpha {
	out := alpha{
		alpha:    al,
		variable: variable,
	}

	return &out
}

// Alpha returns the alpha value
func (obj *alpha) Alpha() uint8 {
	return obj.alpha
}

// Variable returns the variable name
func (obj *alpha) Variable() string {
	return obj.variable
}
