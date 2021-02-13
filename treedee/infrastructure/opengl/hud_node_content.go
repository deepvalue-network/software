package opengl

type hudNodeContent struct {
	model   Model
	display HudDisplay
}

func createHudNodeContentWithModel(
	model Model,
) HudNodeContent {
	return createHudNodeContentInternally(model, nil)
}

func createHudNodeContentWithDisplay(
	display HudDisplay,
) HudNodeContent {
	return createHudNodeContentInternally(nil, display)
}

func createHudNodeContentInternally(
	model Model,
	display HudDisplay,
) HudNodeContent {
	out := hudNodeContent{
		model:   model,
		display: display,
	}

	return &out
}

// IsModel returns true if there is a model, false otherwise
func (obj *hudNodeContent) IsModel() bool {
	return obj.model != nil
}

// Model returns the model, if any
func (obj *hudNodeContent) Model() Model {
	return obj.model
}

// IsDisplay returns true if there is a display, false otherwise
func (obj *hudNodeContent) IsDisplay() bool {
	return obj.display != nil
}

// Display returns the display, if any
func (obj *hudNodeContent) Display() HudDisplay {
	return obj.display
}
