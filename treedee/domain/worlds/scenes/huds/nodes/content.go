package nodes

import (
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/displays"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models"
)

type content struct {
	model   models.Model
	display displays.Display
}

func createContentWithModel(
	model models.Model,
) Content {
	return createContentInternally(model, nil)
}

func createContentWithDisplay(
	display displays.Display,
) Content {
	return createContentInternally(nil, display)
}

func createContentInternally(
	model models.Model,
	display displays.Display,
) Content {
	out := content{
		model:   model,
		display: display,
	}

	return &out
}

// IsModel returns true if the content is a model, false otherwise
func (obj *content) IsModel() bool {
	return obj.model != nil
}

// Model returns the model, if any
func (obj *content) Model() models.Model {
	return obj.model
}

// IsDisplay returns true if the content is a display, false otherwise
func (obj *content) IsDisplay() bool {
	return obj.display != nil
}

// Display returns the display, if any
func (obj *content) Display() displays.Display {
	return obj.display
}
