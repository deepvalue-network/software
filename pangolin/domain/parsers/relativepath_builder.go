package parsers

import "errors"

type relativePathBuilder struct {
	sections []FolderSection
}

func createRelativePathBuilder() RelativePathBuilder {
	out := relativePathBuilder{
		sections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *relativePathBuilder) Create() RelativePathBuilder {
	return createRelativePathBuilder()
}

// WithSections add sections to the builder
func (app *relativePathBuilder) WithSections(sections []FolderSection) RelativePathBuilder {
	app.sections = sections
	return app
}

// Now builds a new RelativePath instance
func (app *relativePathBuilder) Now() (RelativePath, error) {
	if app.sections == nil {
		app.sections = []FolderSection{}
	}

	if len(app.sections) <= 0 {
		return nil, errors.New("there must be at least 1 FolderSection in order to build a RelativePath")
	}

	head := []FolderSection{}
	var tail FolderSection
	for _, oneSection := range app.sections {
		if oneSection.IsTail() {
			tail = oneSection
			continue
		}

		head = append(head, oneSection)
	}

	if tail != nil {
		return createRelativePathWithTail(head, tail), nil
	}

	return createRelativePath(head), nil
}
