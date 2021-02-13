package opengl

type nodeContent struct {
	model Model
	cam   Camera
}

func createNodeContentWithModel(
	model Model,
) NodeContent {
	return createNodeContentInternally(model, nil)
}

func createNodeContentWithCamera(
	cam Camera,
) NodeContent {
	return createNodeContentInternally(nil, cam)
}

func createNodeContentInternally(
	model Model,
	cam Camera,
) NodeContent {
	out := nodeContent{
		model: model,
		cam:   cam,
	}

	return &out
}

// IsModel returns true if there is a model, false otherwise
func (obj *nodeContent) IsModel() bool {
	return obj.model != nil
}

// Model returns the model, if any
func (obj *nodeContent) Model() Model {
	return obj.model
}

// IsCamera returns true if there is a camera, false otherwise
func (obj *nodeContent) IsCamera() bool {
	return obj.cam != nil
}

// Camera returns the camera, if any
func (obj *nodeContent) Camera() Camera {
	return obj.cam
}
