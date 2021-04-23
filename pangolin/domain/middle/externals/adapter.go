package externals

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

// ToExternals converts parsed import singles, to externals
func (app *adapter) ToExternals(parsed []parsers.ImportSingle) ([]External, error) {
	out := []External{}
	for _, oneParsed := range parsed {
		ext, err := app.ToExternal(oneParsed)
		if err != nil {
			return nil, err
		}

		out = append(out, ext)
	}

	return out, nil
}

// ToExternal converts parsed import single, to external
func (app *adapter) ToExternal(parsed parsers.ImportSingle) (External, error) {
	name := parsed.Name()
	path := parsed.Path().String()
	return app.builder.Create().
		WithPath(path).
		WithName(name).
		Now()
}
