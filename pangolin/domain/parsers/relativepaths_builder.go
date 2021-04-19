package parsers

import "errors"

type relativePathsBuilder struct {
	list []RelativePath
}

func createRelativePathsBuilder() RelativePathsBuilder {
	out := relativePathsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *relativePathsBuilder) Create() RelativePathsBuilder {
	return createRelativePathsBuilder()
}

// WithRelativePaths add paths to the builder
func (app *relativePathsBuilder) WithRelativePaths(relPaths []RelativePath) RelativePathsBuilder {
	app.list = relPaths
	return app
}

// Now builds a new RelativePaths instance
func (app *relativePathsBuilder) Now() (RelativePaths, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 RelativePath in order to build a RelativePaths instance")
	}

	return createRelativePaths(app.list), nil
}
