package heads

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/heads"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type value struct {
	name    string
	version string
	imports []parsers.ImportSingle
	loads   []heads.LoadSingle
}

func createValueWithName(name string) Value {
	return createValueInternally(name, "", nil, nil)
}

func createValueWithVersion(version string) Value {
	return createValueInternally("", version, nil, nil)
}

func createValueWithImports(imports []parsers.ImportSingle) Value {
	return createValueInternally("", "", imports, nil)
}

func createValueWithLoads(loads []heads.LoadSingle) Value {
	return createValueInternally("", "", nil, loads)
}

func createValueInternally(
	name string,
	version string,
	imports []parsers.ImportSingle,
	loads []heads.LoadSingle,
) Value {
	out := value{
		name:    name,
		version: version,
		imports: imports,
		loads:   loads,
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
func (obj *value) Imports() []parsers.ImportSingle {
	return obj.imports
}

// IsLoads returns true if there is loads, false otherwise
func (obj *value) IsLoads() bool {
	return obj.loads != nil
}

// Loads returns the loads, if any
func (obj *value) Loads() []heads.LoadSingle {
	return obj.loads
}
