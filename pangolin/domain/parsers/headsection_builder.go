package parsers

import "errors"

type headSectionBuilder struct {
	values []HeadValue
}

func createHeadSectionBuilder() HeadSectionBuilder {
	out := headSectionBuilder{
		values: nil,
	}

	return &out
}

// Create initializes the builder
func (app *headSectionBuilder) Create() HeadSectionBuilder {
	return createHeadSectionBuilder()
}

// WithValues add values to the builder
func (app *headSectionBuilder) WithValues(values []HeadValue) HeadSectionBuilder {
	app.values = values
	return app
}

// Now builds a new HeadSection instance
func (app *headSectionBuilder) Now() (HeadSection, error) {
	if app.values == nil {
		app.values = []HeadValue{}
	}

	name := ""
	version := ""
	imports := []ImportSingle{}
	for _, oneValue := range app.values {
		if oneValue.IsName() {
			name = oneValue.Name()
			continue
		}

		if oneValue.IsVersion() {
			version = oneValue.Version()
			continue
		}

		if oneValue.IsImport() {
			imports = oneValue.Import()
			continue
		}
	}

	if name == "" {
		return nil, errors.New("the name is mandatory in order to build an HeadSection instance")
	}

	if version == "" {
		return nil, errors.New("the version is mandatory in order to build an HeadSection instance")
	}

	if len(imports) > 0 {
		return createHeadSectionWithImport(name, version, imports), nil
	}

	return createHeadSection(name, version), nil
}
