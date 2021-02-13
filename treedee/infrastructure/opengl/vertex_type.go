package opengl

type vertexType struct {
	isTriangle bool
}

func createVertexTypeWithTriangle() VertexType {
	return createVertexTypeInternally(true)
}

func createVertexTypeInternally(
	isTriangle bool,
) VertexType {
	out := vertexType{
		isTriangle: isTriangle,
	}

	return &out
}

// IsTriangle returns true if the vertex type is triangle, false otherwise
func (obj *vertexType) IsTriangle() bool {
	return obj.isTriangle
}
