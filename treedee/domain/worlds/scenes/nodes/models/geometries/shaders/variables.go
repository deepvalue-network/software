package shaders

type variables struct {
	tex    string
	vertex string
}

func createVariables(
	tex string,
	vertex string,
) Variables {
	out := variables{
		tex:    tex,
		vertex: vertex,
	}

	return &out
}

// TextureCoordinates returns the texture coordinates
func (obj *variables) TextureCoordinates() string {
	return obj.tex
}

// VertexCoordinates returns the vertex coordinates
func (obj *variables) VertexCoordinates() string {
	return obj.vertex
}
