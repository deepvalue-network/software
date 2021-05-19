package heads

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type head struct {
	name    string
	version string
	imports []parsers.ImportSingle
}

func createHead(
	name string,
	version string,
) Head {
	return createHeadInternally(name, version, nil)
}

func createHeadWithImports(
	name string,
	version string,
	imports []parsers.ImportSingle,
) Head {
	return createHeadInternally(name, version, imports)
}

func createHeadInternally(
	name string,
	version string,
	imports []parsers.ImportSingle,
) Head {
	out := head{
		name:    name,
		version: version,
		imports: imports,
	}

	return &out
}

// Name returns the name
func (obj *head) Name() string {
	return obj.name
}

// Version returns the version
func (obj *head) Version() string {
	return obj.version
}

// HasImports returns true if there is imports, false otherwise
func (obj *head) HasImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *head) Imports() []parsers.ImportSingle {
	return obj.imports
}
