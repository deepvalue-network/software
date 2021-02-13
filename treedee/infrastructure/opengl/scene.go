package opengl

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type scene struct {
	worldCameras map[string]WorldCamera
	id           *uuid.UUID
	index        uint
	hud          Hud
	nodes        []Node
}

func createScene(
	worldCameras map[string]WorldCamera,
	id *uuid.UUID,
	index uint,
	hud Hud,
	nodes []Node,
) Scene {
	return createSceneInternally(worldCameras, id, index, hud, nodes)
}

func createSceneInternally(
	worldCameras map[string]WorldCamera,
	id *uuid.UUID,
	index uint,
	hud Hud,
	nodes []Node,
) Scene {
	out := scene{
		worldCameras: worldCameras,
		id:           id,
		index:        index,
		hud:          hud,
		nodes:        nodes,
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

// Hud returns the hud
func (obj *scene) Hud() Hud {
	return obj.hud
}

// HasNodes returns true if there is nodes, false otherwise
func (obj *scene) HasNodes() bool {
	return obj.nodes != nil
}

// Nodes returns the nodes, if any
func (obj *scene) Nodes() []Node {
	return obj.nodes
}

// Render renders a scene
func (obj *scene) Render(
	delta time.Duration,
	activeCameraID *uuid.UUID,
) error {
	// retrieve the active camera:
	keyname := activeCameraID.String()
	if activeCamera, ok := obj.worldCameras[keyname]; ok {
		// render the nodes, if any:
		if obj.HasNodes() {
			for _, oneNode := range obj.nodes {
				err := oneNode.Render(delta, activeCamera, obj)
				if err != nil {
					return err
				}
			}
		}

		// render the heads-up display:
		//pos := activeCamera.Position()
		//orientation := activeCamera.Orientation()
		return obj.hud.Render(delta, obj)
	}

	str := fmt.Sprintf("the given active Camera (ID: %s) does not exists in the current Scene (ID: %s)", keyname, obj.ID().String())
	return errors.New(str)
}
