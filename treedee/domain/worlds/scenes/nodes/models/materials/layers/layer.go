package layers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/alphas"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures"
)

type layer struct {
	id    *uuid.UUID
	index uint
	alpha alphas.Alpha
	tex   textures.Texture
}

func createLayer(
	id *uuid.UUID,
	index uint,
	alpha alphas.Alpha,
	tex textures.Texture,
) Layer {
	out := layer{
		id:    id,
		index: index,
		alpha: alpha,
		tex:   tex,
	}

	return &out
}

// ID returns the id
func (obj *layer) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *layer) Index() uint {
	return obj.index
}

// Alpha returns the alpha
func (obj *layer) Alpha() alphas.Alpha {
	return obj.alpha
}

// Texture returns the texture
func (obj *layer) Texture() textures.Texture {
	return obj.tex
}
