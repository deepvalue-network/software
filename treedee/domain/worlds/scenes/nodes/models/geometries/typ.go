package geometries

type typ struct {
	isTriangle bool
}

func createTypeWithTriangle() Type {
	return createTypeInternally(true)
}

func createTypeInternally(
	isTriangle bool,
) Type {
	out := typ{
		isTriangle: isTriangle,
	}

	return &out
}

// IsTriangle returns true if the type is triangle, false otherwise
func (obj *typ) IsTriangle() bool {
	return obj.isTriangle
}
