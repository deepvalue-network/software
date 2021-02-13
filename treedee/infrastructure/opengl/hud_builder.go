package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds"
)

type hudBuilder struct {
	hudNodeBuilder  HudNodeBuilder
	materialBuilder MaterialBuilder
	programBuilder  ProgramBuilder
	hud             huds.Hud
}

func createHudBuilder(
	hudNodeBuilder HudNodeBuilder,
	materialBuilder MaterialBuilder,
	programBuilder ProgramBuilder,
) HudBuilder {
	out := hudBuilder{
		hudNodeBuilder:  hudNodeBuilder,
		materialBuilder: materialBuilder,
		programBuilder:  programBuilder,
	}

	return &out
}

// Create initializes the buildeer
func (app *hudBuilder) Create() HudBuilder {
	return createHudBuilder(app.hudNodeBuilder, app.materialBuilder, app.programBuilder)
}

// WithHud adds an hud to the builder
func (app *hudBuilder) WithHud(hud huds.Hud) HudBuilder {
	app.hud = hud
	return app
}

// Now builds a new Hud instance
func (app *hudBuilder) Now() (Hud, error) {
	if app.hud == nil {
		return nil, errors.New("the HUD is mandatory in order to build an HUD instance")
	}

	nodes := []HudNode{}
	if app.hud.HasNodes() {
		domainNodes := app.hud.Nodes()
		for _, oneDomainNode := range domainNodes {
			node, err := app.hudNodeBuilder.Create().WithNode(oneDomainNode).Now()
			if err != nil {
				return nil, err
			}

			nodes = append(nodes, node)
		}
	}

	if len(nodes) <= 0 {
		nodes = nil
	}

	prog, err := app.programBuilder.Create().Now()
	if err != nil {
		return nil, err
	}

	id := app.hud.ID()
	if app.hud.HasMaterial() {
		domainMaterial := app.hud.Material()
		mat, err := app.materialBuilder.Create().WithMaterial(domainMaterial).Now()
		if err != nil {
			return nil, err
		}

		if nodes != nil {
			return createHudWithNodesAndMaterial(id, prog, nodes, mat), nil
		}

		return createHudWithMaterial(id, prog, mat), nil
	}

	if nodes != nil {
		return createHudWithNodes(id, prog, nodes), nil
	}

	return nil, errors.New("the HUD is invalid")
}
