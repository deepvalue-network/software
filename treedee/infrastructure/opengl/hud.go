package opengl

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
)

type hud struct {
	id    *uuid.UUID
	prog  uint32
	nodes []HudNode
	mat   Material
}

func createHudWithNodes(
	id *uuid.UUID,
	prog uint32,
	nodes []HudNode,
) Hud {
	return createHudInternally(id, prog, nodes, nil)
}

func createHudWithMaterial(
	id *uuid.UUID,
	prog uint32,
	mat Material,
) Hud {
	return createHudInternally(id, prog, nil, mat)
}

func createHudWithNodesAndMaterial(
	id *uuid.UUID,
	prog uint32,
	nodes []HudNode,
	mat Material,
) Hud {
	return createHudInternally(id, prog, nodes, mat)
}

func createHudInternally(
	id *uuid.UUID,
	prog uint32,
	nodes []HudNode,
	mat Material,
) Hud {
	out := hud{
		id:    id,
		prog:  prog,
		nodes: nodes,
		mat:   mat,
	}

	return &out
}

// ID returns the id
func (obj *hud) ID() *uuid.UUID {
	return obj.id
}

// Program returns the program
func (obj *hud) Program() uint32 {
	return obj.prog
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *hud) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes, if any
func (obj *hud) Nodes() []HudNode {
	return obj.nodes
}

// HasMaterial returns true if there is material, false otherwise
func (obj *hud) HasMaterial() bool {
	return obj.mat != nil
}

// Material returns the material, if any
func (obj *hud) Material() Material {
	return obj.mat
}

// Render renders the head-up display
func (obj *hud) Render(
	delta time.Duration,
	activeScene Scene,
) error {

	// render the nodes, if any:
	if obj.HasNodes() {
		for _, oneNode := range obj.nodes {
			err := oneNode.Render(delta, activeScene)
			if err != nil {
				return err
			}
		}
	}

	// render the material, if any:
	if obj.HasMaterial() {
		// use the program:
		gl.UseProgram(obj.prog)

		// render the material:
		mat := obj.Material()
		err := mat.Render(delta, nil, nil, activeScene, obj.prog)
		if err != nil {
			return err
		}
	}

	return nil
}
