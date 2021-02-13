package opengl

type vertexShaderVariables struct {
	tex    string
	vertex string
}

func createVertexShaderVariables(
	tex string,
	vertex string,
) VertexShaderVariables {
	out := vertexShaderVariables{
		tex:    tex,
		vertex: vertex,
	}

	return &out
}

// TextureCoordinates returns the texture coordinates variable
func (obj *vertexShaderVariables) TextureCoordinates() string {
	return obj.tex
}

// VertexCoordinates returns the vertex coordinates variable
func (obj *vertexShaderVariables) VertexCoordinates() string {
	return obj.vertex
}
