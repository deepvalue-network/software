package scenes

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes"
)

type scene struct {
	id    *uuid.UUID
	index uint
	hud   huds.Hud
	nodes []nodes.Node
}

func createScene(
	id *uuid.UUID,
	index uint,
	hud huds.Hud,
	nodes []nodes.Node,
) Scene {
	return createSceneInternally(id, index, hud, nodes)
}

func createSceneInternally(
	id *uuid.UUID,
	index uint,
	hud huds.Hud,
	nodes []nodes.Node,
) Scene {
	out := scene{
		id:    id,
		index: index,
		hud:   hud,
		nodes: nodes,
	}

	return &out
}

// ID returns the id
func (obj *scene) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *scene) Index() uint {
	return obj.index
}

// Hud returns the heads-up display
func (obj *scene) Hud() huds.Hud {
	return obj.hud
}

// Nodes return the nodes, if any
func (obj *scene) Nodes() []nodes.Node {
	return obj.nodes
}
