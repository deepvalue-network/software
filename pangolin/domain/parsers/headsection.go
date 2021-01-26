package parsers

type headSection struct {
	name    string
	version string
	imports []ImportSingle
}

func createHeadSection(
	name string,
	version string,
) HeadSection {
	return createHeadSectionInternally(name, version, nil)
}

func createHeadSectionWithImport(
	name string,
	version string,
	imports []ImportSingle,
) HeadSection {
	return createHeadSectionInternally(name, version, imports)
}

func createHeadSectionInternally(
	name string,
	version string,
	imports []ImportSingle,
) HeadSection {
	out := headSection{
		name:    name,
		version: version,
		imports: imports,
	}

	return &out
}

// Name returns the name
func (obj *headSection) Name() string {
	return obj.name
}

// Version returns the version
func (obj *headSection) Version() string {
	return obj.version
}

// HasImport returns true if there is an import, false otherwise
func (obj *headSection) HasImport() bool {
	return obj.imports != nil
}

// Import returns the import, if any
func (obj *headSection) Import() []ImportSingle {
	return obj.imports
}
