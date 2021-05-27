package labels

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	labelAdapter label.Adapter
	builder      Builder
}

func createAdapter(
	labelAdapter label.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		labelAdapter: labelAdapter,
		builder:      builder,
	}

	return &out
}

// ToLabels converts a parsed language label section to labels
func (app *adapter) ToLabels(parsed parsers.LanguageLabelSection) (Labels, error) {
	out := []label.Label{}
	declarations := parsed.Declarations()
	for _, oneDeclaration := range declarations {
		label, err := app.labelAdapter.ToLabel(oneDeclaration)
		if err != nil {
			return nil, err
		}

		out = append(out, label)
	}

	return app.builder.Create().WithList(out).Now()
}
