package opengl

import (
	"errors"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes"
)

type nodeBuilder struct {
	modelBuilder  ModelBuilder
	cameraBuilder CameraBuilder
	node          nodes.Node
}

func createNodeBuilder(
	modelBuilder ModelBuilder,
	cameraBuilder CameraBuilder,
) NodeBuilder {
	out := nodeBuilder{
		modelBuilder:  modelBuilder,
		cameraBuilder: cameraBuilder,
		node:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder(app.modelBuilder, app.cameraBuilder)
}

// WithNode adds a node to the builder
func (app *nodeBuilder) WithNode(node nodes.Node) NodeBuilder {
	app.node = node
	return app
}

// Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.node == nil {
		return nil, errors.New("the node is mandatory in order to build a Node instance")
	}

	return app.convert(app.node)
}

func (app *nodeBuilder) convertAll(nodes []nodes.Node) ([]Node, error) {
	out := []Node{}
	for _, oneNode := range nodes {
		node, err := app.convert(oneNode)
		if err != nil {
			return nil, err
		}

		out = append(out, node)
	}

	return out, nil
}

func (app *nodeBuilder) convert(node nodes.Node) (Node, error) {
	domainPos := node.Position()
	posVec := domainPos.Vector()
	posVariable := domainPos.Variable()
	pos := createPosition(
		mgl32.Vec3{
			posVec.X(),
			posVec.Y(),
			posVec.Z(),
		},
		posVariable,
	)

	domainOrientation := node.Orientation()
	angle := domainOrientation.Angle()
	domainDirection := domainOrientation.Direction()
	orientationVariable := domainOrientation.Variable()
	orientation := createOrientation(
		angle,
		mgl32.Vec3{
			domainDirection.X(),
			domainDirection.Y(),
			domainDirection.Z(),
		},
		orientationVariable,
	)

	id := node.ID()
	if !node.HasContent() {
		return createNode(id, pos, orientation), nil
	}

	children := []Node{}
	if node.HasChildren() {
		domainChildren := node.Children()
		list, err := app.convertAll(domainChildren)
		if err != nil {
			return nil, err
		}

		children = list
	}

	domainContent := node.Content()
	if domainContent.IsModel() {
		domainModel := domainContent.Model()
		model, err := app.modelBuilder.Create().WithModel(domainModel).Now()
		if err != nil {
			return nil, err
		}

		content := createNodeContentWithModel(model)
		if len(children) > 0 {
			return createNodeWithContentAndChildren(id, pos, orientation, content, children), nil
		}

		return createNodeWithContent(id, pos, orientation, content), nil
	}

	if len(children) > 0 {
		return createNodeWithChildren(id, pos, orientation, children), nil
	}

	return createNode(id, pos, orientation), nil
}
