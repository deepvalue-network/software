package opengl

import (
	"errors"

	"github.com/go-gl/mathgl/mgl32"
	hud_nodes "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/nodes"
)

type hudNodeBuilder struct {
	modelBuilder      ModelBuilder
	hudDisplayBuilder HudDisplayBuilder
	node              hud_nodes.Node
}

func createHudNodeBuilder(
	modelBuilder ModelBuilder,
	hudDisplayBuilder HudDisplayBuilder,
) HudNodeBuilder {
	out := hudNodeBuilder{
		modelBuilder:      modelBuilder,
		hudDisplayBuilder: hudDisplayBuilder,
		node:              nil,
	}

	return &out
}

// Create initializes the builder
func (app *hudNodeBuilder) Create() HudNodeBuilder {
	return createHudNodeBuilder(
		app.modelBuilder,
		app.hudDisplayBuilder,
	)
}

// WithNode adds a node to the builder
func (app *hudNodeBuilder) WithNode(node hud_nodes.Node) HudNodeBuilder {
	app.node = node
	return app
}

// Now builds a new HudNode instance
func (app *hudNodeBuilder) Now() (HudNode, error) {
	if app.node == nil {
		return nil, errors.New("the hudNode is mandatory in order to build a HudNode instance")
	}

	return app.convert(app.node)
}

func (app *hudNodeBuilder) convertAll(nodes []hud_nodes.Node) ([]HudNode, error) {
	out := []HudNode{}
	for _, oneNode := range nodes {
		node, err := app.convert(oneNode)
		if err != nil {
			return nil, err
		}

		out = append(out, node)
	}

	return out, nil
}

func (app *hudNodeBuilder) convert(node hud_nodes.Node) (HudNode, error) {
	domainPos := node.Position()
	posVec := domainPos.Vector()
	posVariable := domainPos.Variable()
	pos := createHudPosition(
		mgl32.Vec2{
			posVec.X(),
			posVec.Y(),
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
		return createHudNode(id, pos, orientation), nil
	}

	children := []HudNode{}
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

		content := createHudNodeContentWithModel(model)
		if len(children) > 0 {
			return createHudNodeWithContentAndChildren(id, pos, orientation, content, children), nil
		}

		return createHudNodeWithContent(id, pos, orientation, content), nil
	}

	if len(children) > 0 {
		return createHudNodeWithChildren(id, pos, orientation, children), nil
	}

	return createHudNode(id, pos, orientation), nil
}
