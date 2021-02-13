package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes"
)

type sceneBuilder struct {
	hudBuilder         HudBuilder
	nodeBuilder        NodeBuilder
	worldCameraBuilder WorldCameraBuilder
	scene              scenes.Scene
}

func createSceneBuilder(
	hudBuilder HudBuilder,
	nodeBuilder NodeBuilder,
	worldCameraBuilder WorldCameraBuilder,
) SceneBuilder {
	out := sceneBuilder{
		hudBuilder:         hudBuilder,
		nodeBuilder:        nodeBuilder,
		worldCameraBuilder: worldCameraBuilder,
	}

	return &out
}

// Create initializes the builder
func (app *sceneBuilder) Create() SceneBuilder {
	return createSceneBuilder(app.hudBuilder, app.nodeBuilder, app.worldCameraBuilder)
}

// WithScene adds a scene to the builder
func (app *sceneBuilder) WithScene(scene scenes.Scene) SceneBuilder {
	app.scene = scene
	return app
}

// Now builds a new Scene instance
func (app *sceneBuilder) Now() (Scene, error) {
	if app.scene == nil {
		return nil, errors.New("the scene is mandatory in order to build a Scene instance")
	}

	domainHud := app.scene.Hud()
	hud, err := app.hudBuilder.Create().WithHud(domainHud).Now()
	if err != nil {
		return nil, err
	}

	id := app.scene.ID()
	index := app.scene.Index()
	nodes := []Node{}
	domainNodes := app.scene.Nodes()
	for _, oneNode := range domainNodes {
		node, err := app.nodeBuilder.Create().WithNode(oneNode).Now()
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	worldCameras, err := app.retrieveWorldCamerasFromNodes(nodes)
	if err != nil {
		return nil, err
	}

	return createScene(worldCameras, id, index, hud, nodes), nil
}

func (app *sceneBuilder) retrieveWorldCamerasFromNodes(nodes []Node) (map[string]WorldCamera, error) {
	if nodes == nil {
		return map[string]WorldCamera{}, nil
	}

	out := map[string]WorldCamera{}
	for _, oneNode := range nodes {

		if oneNode.HasChildren() {
			children := oneNode.Children()
			subWorldCameras, err := app.retrieveWorldCamerasFromNodes(children)
			if err != nil {
				return nil, err
			}

			for keyname, oneCam := range subWorldCameras {
				out[keyname] = oneCam
			}
		}

		if !oneNode.HasContent() {
			continue
		}

		content := oneNode.Content()
		if !content.IsCamera() {
			continue
		}

		worldCamera, err := app.worldCameraBuilder.Create().WithNode(oneNode).Now()
		if err != nil {
			return nil, err
		}

		keyname := oneNode.ID().String()
		out[keyname] = worldCamera
	}

	return out, nil
}
