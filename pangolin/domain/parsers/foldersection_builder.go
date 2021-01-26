package parsers

import "errors"

type folderSectionBuilder struct {
	isTail bool
	name   FolderName
}

func createFolderSectionBuilder() FolderSectionBuilder {
	out := folderSectionBuilder{
		isTail: false,
		name:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *folderSectionBuilder) Create() FolderSectionBuilder {
	return createFolderSectionBuilder()
}

// IsTail flags the builder as a tail
func (app *folderSectionBuilder) IsTail() FolderSectionBuilder {
	app.isTail = true
	return app
}

// WithName adds a name to the builder
func (app *folderSectionBuilder) WithName(name FolderName) FolderSectionBuilder {
	app.name = name
	return app
}

// Now builds a new FolderSection instance
func (app *folderSectionBuilder) Now() (FolderSection, error) {
	if app.name == nil {
		return nil, errors.New("the FolderName is mandatory in order to build a FolderSection instance")
	}

	if app.isTail {
		return createFolderSectionWithTail(app.name), nil
	}

	return createFolderSection(app.name), nil
}
