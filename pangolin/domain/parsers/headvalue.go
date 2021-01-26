package parsers

type headValue struct {
	name    string
	version string
	imports []ImportSingle
}

func createHeadValueWithName(
	name string,
) HeadValue {
	return createHeadValueInternally(name, "", nil)
}

func createHeadValueWithVersion(
	version string,
) HeadValue {
	return createHeadValueInternally("", version, nil)
}

func createHeadValueWithImport(
	imports []ImportSingle,
) HeadValue {
	return createHeadValueInternally("", "", imports)
}

func createHeadValueInternally(
	name string,
	version string,
	imports []ImportSingle,
) HeadValue {
	out := headValue{
		name:    name,
		version: version,
		imports: imports,
	}

	return &out
}

// IsName returns true if there is a name, false otherwise
func (obj *headValue) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *headValue) Name() string {
	return obj.name
}

// IsVersion returns true if there is a version, false otherwise
func (obj *headValue) IsVersion() bool {
	return obj.version != ""
}

// Version returns the version, if any
func (obj *headValue) Version() string {
	return obj.version
}

// IsImport returns true if there is an import, false otherwise
func (obj *headValue) IsImport() bool {
	return obj.imports != nil
}

// Import returns the import, if any
func (obj *headValue) Import() []ImportSingle {
	return obj.imports
}
