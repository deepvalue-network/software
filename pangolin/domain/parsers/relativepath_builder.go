package parsers

import (
	"errors"
	"path/filepath"
)

type relativePathBuilder struct {
	folderNameBuilder    FolderNameBuilder
	folderSectionBuilder FolderSectionBuilder
	sections             []FolderSection
	path                 string
}

func createRelativePathBuilder(
	folderNameBuilder FolderNameBuilder,
	folderSectionBuilder FolderSectionBuilder,
) RelativePathBuilder {
	out := relativePathBuilder{
		folderNameBuilder:    folderNameBuilder,
		folderSectionBuilder: folderSectionBuilder,
		sections:             nil,
		path:                 "",
	}

	return &out
}

// Create initializes the builder
func (app *relativePathBuilder) Create() RelativePathBuilder {
	return createRelativePathBuilder(app.folderNameBuilder, app.folderSectionBuilder)
}

// WithSections add sections to the builder
func (app *relativePathBuilder) WithSections(sections []FolderSection) RelativePathBuilder {
	app.sections = sections
	return app
}

func (app *relativePathBuilder) WithPath(path string) RelativePathBuilder {
	app.path = path
	return app
}

// Now builds a new RelativePath instance
func (app *relativePathBuilder) Now() (RelativePath, error) {
	if app.path != "" {
		folderSections := []FolderSection{}
		folders := filepath.SplitList(app.path)
		amount := len(folders)
		for index, oneName := range folders {
			folderNameBuilder := app.folderNameBuilder.Create()

			hasDots := false
			if oneName == "." {
				folderNameBuilder.IsCurrent()
				hasDots = true
			}

			if oneName == ".." {
				folderNameBuilder.IsPrevious()
				hasDots = true
			}

			if !hasDots {
				folderNameBuilder.WithName(oneName)
			}

			folderName, err := folderNameBuilder.Now()
			if err != nil {
				return nil, err
			}

			folderSectionBuilder := app.folderSectionBuilder.Create().WithName(folderName)
			if (index + 1) >= amount {
				folderSectionBuilder.IsTail()
			}

			folderSection, err := folderSectionBuilder.Now()
			if err != nil {
				return nil, err
			}

			folderSections = append(folderSections, folderSection)
		}

		app.sections = folderSections
	}

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
