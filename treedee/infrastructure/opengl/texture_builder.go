package opengl

import (
	"errors"
	"fmt"
	"image"
	image_color "image/color"
	"image/draw"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures"
)

type textureBuilder struct {
	programBuilder ProgramBuilder
	cameraBuilder  CameraBuilder
	tex            textures.Texture
}

func createBuilder(
	programBuilder ProgramBuilder,
	cameraBuilder CameraBuilder,
) TextureBuilder {
	out := textureBuilder{
		programBuilder: programBuilder,
		cameraBuilder:  cameraBuilder,
		tex:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *textureBuilder) Create() TextureBuilder {
	return createBuilder(
		app.programBuilder,
		app.cameraBuilder,
	)
}

// WithTexture adds a texture to the builder
func (app *textureBuilder) WithTexture(tex textures.Texture) TextureBuilder {
	app.tex = tex
	return app
}

// Now builds a new Texture instance
func (app *textureBuilder) Now() (Texture, error) {
	if app.tex == nil {
		return nil, errors.New("the texture is mandatory in order to build a Texture instance")
	}

	id := app.tex.ID()
	dim := app.tex.Dimension()
	variable := app.tex.Variable()

	if app.tex.IsCamera() {
		domainCam := app.tex.Camera()
		cam, err := app.cameraBuilder.Create().WithCamera(domainCam).Now()
		if err != nil {
			return nil, err
		}

		return createTextureWithCamera(id, dim, variable, cam), nil
	}

	if app.tex.IsShader() {
		domainShader := app.tex.Shader()
		shaderProg, err := app.programBuilder.Create().WithTextureShader(domainShader).Now()
		if err != nil {
			return nil, err
		}

		shaderID := domainShader.ID()
		isDynamic := domainShader.IsDynamic()
		shader := createTextureShader(shaderID, shaderProg, isDynamic)
		return createTextureWithShader(id, dim, variable, shader), nil
	}

	if app.tex.IsPixels() {
		pixels := app.tex.Pixels()
		width := int(dim.X())
		height := int(dim.Y())
		srcRGBA := image.NewRGBA(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: width,
				Y: height,
			},
		})

		y := -1
		for index, onePixel := range pixels {
			color := onePixel.Color()
			alpha := onePixel.Alpha()
			rgbaColor := image_color.RGBA{
				R: color.Red(),
				G: color.Green(),
				B: color.Blue(),
				A: alpha,
			}

			x := index % width
			srcRGBA.Set(x, y, rgbaColor)

			if x == 0 {
				y++
			}
		}

		dstRGBA := image.NewRGBA(
			image.Rectangle{
				Min: image.Point{
					X: 0,
					Y: 0,
				},
				Max: image.Point{
					X: width,
					Y: height,
				},
			},
		)

		if dstRGBA.Stride != dstRGBA.Rect.Size().X*4 {
			return nil, fmt.Errorf("the destination RGBA image has an unsupported stride")
		}

		draw.Draw(dstRGBA, dstRGBA.Bounds(), srcRGBA, image.Point{0, 0}, draw.Src)

		var identifier uint32
		gl.GenTextures(1, &identifier)
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, identifier)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(dstRGBA.Rect.Size().X),
			int32(dstRGBA.Rect.Size().Y),
			0,
			gl.RGBA,
			gl.UNSIGNED_BYTE,
			gl.Ptr(dstRGBA.Pix),
		)

		return createTextureWithResource(id, dim, variable, identifier), nil
	}

	return nil, errors.New("the Texture is invalid")
}
