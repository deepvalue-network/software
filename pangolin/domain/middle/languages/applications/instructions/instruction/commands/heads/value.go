package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
)

type value struct {
	name    string
	version string
	imports []heads.External
}

func createValueWithName(name string) Value {
	return createValueInternally(name, "", nil)
}

func createValueWithVersion(version string) Value {
	return createValueInternally("", version, nil)
}

func createValueWithImports(imports []heads.External) Value {
	return createValueInternally("", "", imports)
}

func createValueInternally(
	name string,
	version string,
	imports []heads.External,
) Value {
	out := value{
		name:    name,
		version: version,
		imports: imports,
	}

	return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *value) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *value) Name() string {
	return obj.name
}

// IsVersion returns true if there is a version, false otherwise
func (obj *value) IsVersion() bool {
	return obj.version != ""
}

// Version returns the version, if any
func (obj *value) Version() string {
	return obj.version
}

// IsImports returns true if there is imports, false otherwise
func (obj *value) IsImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *value) Imports() []heads.External {
	return obj.imports
}
