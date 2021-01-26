package labels

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/labels/label"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

type adapter struct {
	labelAdapter label.Adapter
	builder      Builder
}

func createAdapter(labelAdapter label.Adapter, builder Builder) Adapter {
	out := adapter{
		labelAdapter: labelAdapter,
		builder:      builder,
	}

	return &out
}

// ToLabels converts a parsed LabelSection to an optimized Labels
func (app *adapter) ToLabels(section parsers.LabelSection) (Labels, error) {
	lst := []label.Label{}
	declarations := section.Declarations()
	for _, oneDeclaration := range declarations {
		lbl, err := app.labelAdapter.ToLabel(oneDeclaration)
		if err != nil {
			return nil, err
		}

		lst = append(lst, lbl)
	}

	return app.builder.WithList(lst).Now()
}
