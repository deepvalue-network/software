package huds

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/nodes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

type hud struct {
	id    *uuid.UUID
	nodes []nodes.Node
	mat   materials.Material
}

func createHudWithNodes(
	id *uuid.UUID,
	nodes []nodes.Node,
) Hud {
	return createHudInternally(id, nodes, nil)
}

func createHudWithMaterial(
	id *uuid.UUID,
	mat materials.Material,
) Hud {
	return createHudInternally(id, nil, mat)
}

func createHudWithNodesAndMaterial(
	id *uuid.UUID,
	nodes []nodes.Node,
	mat materials.Material,
) Hud {
	return createHudInternally(id, nodes, mat)
}

func createHudInternally(
	id *uuid.UUID,
	nodes []nodes.Node,
	mat materials.Material,
) Hud {
	out := hud{
		id:    id,
		nodes: nodes,
		mat:   mat,
	}

	return &out
}

// ID returns the id
func (obj *hud) ID() *uuid.UUID {
	return obj.id
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *hud) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes
func (obj *hud) Nodes() []nodes.Node {
	return obj.nodes
}

// HasMaterial returns true if there is material, false otherwise
func (obj *hud) HasMaterial() bool {
	return obj.mat != nil
}

// Material returns the material, if any
func (obj *hud) Material() materials.Material {
	return obj.mat
}
