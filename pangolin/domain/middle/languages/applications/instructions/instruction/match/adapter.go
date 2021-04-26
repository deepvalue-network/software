package match

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type adapter struct {
	builder Builder
}

func createAdapter(
	builder Builder,
) Adapter {
	out := adapter{
		builder: builder,
	}

	return &out
}

// ToMatch converts a parsed match to match instance
func (app *adapter) ToMatch(parsed parsers.Match) (Match, error) {
	input := parsed.Input()
	builder := app.builder.Create().WithInput(input)
	if parsed.HasPattern() {
		pattern := parsed.Pattern()
		builder.WithPattern(pattern)
	}

	return builder.Now()
}
