package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

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

// ToHead converts a parsed head section to an Head instance
func (app *adapter) ToHead(parsed parsers.HeadSection) (Head, error) {
	name := parsed.Name()
	version := parsed.Version()
	builder := app.builder.Create().WithName(name).WithVersion(version)
	if parsed.HasImport() {
		parsedImports := parsed.Import()
		builder.WithImports(parsedImports)
	}

	return builder.Now()
}
