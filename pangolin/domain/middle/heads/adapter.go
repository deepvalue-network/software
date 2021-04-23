package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/externals"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	externalAdapter externals.Adapter
	builder         Builder
}

func createAdapter(
	externalAdapter externals.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		externalAdapter: externalAdapter,
		builder:         builder,
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
		extends, err := app.externalAdapter.ToExternals(parsedImports)
		if err != nil {
			return nil, err
		}

		builder.WithImports(extends)
	}

	return builder.Now()
}
