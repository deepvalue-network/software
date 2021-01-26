package parsers

import "errors"

type folderNameBuilder struct {
	isCurrent  bool
	isPrevious bool
	name       string
}

func createFolderNameBuilder() FolderNameBuilder {
	out := folderNameBuilder{
		isCurrent:  false,
		isPrevious: false,
		name:       "",
	}

	return &out
}

// Create initializes the builder
func (app *folderNameBuilder) Create() FolderNameBuilder {
	return createFolderNameBuilder()
}

// IsCurrent flags the builder as the current folder
func (app *folderNameBuilder) IsCurrent() FolderNameBuilder {
	app.isCurrent = true
	return app
}

// IsPrevious flags the builder as the previous folder
func (app *folderNameBuilder) IsPrevious() FolderNameBuilder {
	app.isPrevious = true
	return app
}

// WithName adds a name to the builder
func (app *folderNameBuilder) WithName(name string) FolderNameBuilder {
	app.name = name
	return app
}

// Now builds a new FolderName instance
func (app *folderNameBuilder) Now() (FolderName, error) {
	if app.isCurrent {
		return createFolderNameWithCurrent(), nil
	}

	if app.isPrevious {
		return createFolderNameWithPrevious(), nil
	}

	if app.name != "" {
		return createFolderNameWithName(app.name), nil
	}

	return nil, errors.New("the FolderName is invalid")
}
