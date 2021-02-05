package accesses

import "github.com/deepvalue-network/software/identity/domain/accesses/access"

type factory struct {
	builder Builder
}

func createFactory(
	builder Builder,
) Factory {
	out := factory{
		builder: builder,
	}

	return &out
}

// Create creates an accesses instance
func (app *factory) Create() (Accesses, error) {
	mp := map[string]access.Access{}
	return app.builder.Create().WithMap(mp).Now()
}
