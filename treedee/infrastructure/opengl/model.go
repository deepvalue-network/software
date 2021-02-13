package opengl

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	uuid "github.com/satori/go.uuid"
)

type model struct {
	id   *uuid.UUID
	prog uint32
	geo  Geometry
	mat  Material
}

func createModel(
	id *uuid.UUID,
	prog uint32,
	geo Geometry,
	mat Material,
) Model {
	out := model{
		id:   id,
		prog: prog,
		geo:  geo,
		mat:  mat,
	}

	return &out
}

// ID returns the id
func (obj *model) ID() *uuid.UUID {
	return obj.id
}

// Program returns the program
func (obj *model) Program() uint32 {
	return obj.prog
}

// Geometry returns the geometry
func (obj *model) Geometry() Geometry {
	return obj.geo
}

// Material returns the material
func (obj *model) Material() Material {
	return obj.mat
}

// Render renders the model
func (obj *model) Render(
	delta time.Duration,
	activeCamera WorldCamera,
	activeScene Scene,
) error {

	// use the right program:
	gl.UseProgram(obj.prog)

	// fetch the camera:
	camera := activeCamera.Camera()

	// projection:
	projection := camera.Projection()
	projVariable := fmt.Sprintf(glStrPattern, projection.Variable())
	fov := projection.FieldOfView()
	aspectRatio := projection.AspectRation()
	near := projection.Near()
	far := projection.Far()

	projMat := mgl32.Perspective(
		mgl32.DegToRad(fov),
		aspectRatio,
		near,
		far,
	)

	projectionUniform := gl.GetUniformLocation(obj.prog, gl.Str(projVariable))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projMat[0])

	// lookAt:
	lookAt := camera.LookAt()
	lookAtVariable := fmt.Sprintf(glStrPattern, lookAt.Variable())
	eye := lookAt.Eye()
	center := lookAt.Center()
	up := lookAt.Up()

	lookAtMat := mgl32.LookAtV(
		mgl32.Vec3{eye[0], eye[1], eye[2]},
		mgl32.Vec3{center[0], center[1], center[2]},
		mgl32.Vec3{up[0], up[1], up[2]},
	)

	lookAtUniform := gl.GetUniformLocation(obj.prog, gl.Str(lookAtVariable))
	gl.UniformMatrix4fv(lookAtUniform, 1, false, &lookAtMat[0])

	// position:
	pos := activeCamera.Position()
	posVec := pos.Vector()
	posVariable := pos.Variable()
	posVarUniform := gl.GetUniformLocation(obj.prog, gl.Str(posVariable))
	gl.Uniform3f(posVarUniform, posVec.X(), posVec.Y(), posVec.Z())

	// orientation:
	orientation := activeCamera.Orientation()
	angle := orientation.Angle()
	dirVec := orientation.Direction()
	orientationVar := orientation.Variable()
	orientationVarUniform := gl.GetUniformLocation(obj.prog, gl.Str(orientationVar))
	gl.Uniform4f(orientationVarUniform, dirVec.X(), dirVec.Y(), dirVec.Z(), angle)

	// prepare the geometry:
	err := obj.geo.Prepare()
	if err != nil {
		return err
	}

	// render the material:
	err = obj.mat.Render(delta, pos, orientation, activeScene, obj.prog)
	if err != nil {
		return err
	}

	// renders the geometry:
	return obj.geo.Render()
	return nil
}
