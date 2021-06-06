package parsers

type headSection struct {
	name    string
	version string
	imports []ImportSingle
	load    []LoadSingle
}

func createHeadSection(
	name string,
	version string,
) HeadSection {
	return createHeadSectionInternally(name, version, nil, nil)
}

func createHeadSectionWithImport(
	name string,
	version string,
	imports []ImportSingle,
) HeadSection {
	return createHeadSectionInternally(name, version, imports, nil)
}

func createHeadSectionWithLoad(
	name string,
	version string,
	load []LoadSingle,
) HeadSection {
	return createHeadSectionInternally(name, version, nil, load)
}

func createHeadSectionInternally(
	name string,
	version string,
	imports []ImportSingle,
	load []LoadSingle,
) HeadSection {
	out := headSection{
		name:    name,
		version: version,
		imports: imports,
		load:    load,
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

// HasLoad returns true if there is a load, false otherwise
func (obj *headSection) HasLoad() bool {
	return obj.load != nil
}

// Load returns the load, if any
func (obj *headSection) Load() []LoadSingle {
	return obj.load
}
