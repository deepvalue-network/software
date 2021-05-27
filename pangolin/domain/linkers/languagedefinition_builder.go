package linkers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
)

type languageDefinitionBuilder struct {
	app     LanguageApplication
	matches []definitions.PatternMatch
	paths   Paths
	root    string
}

func createLanguageDefinitionBuilder() LanguageDefinitionBuilder {
	out := languageDefinitionBuilder{
		app:     nil,
		matches: nil,
		paths:   nil,
		root:    "",
	}

	return &out
}

// Create initializes the builder
func (app *languageDefinitionBuilder) Create() LanguageDefinitionBuilder {
	return createLanguageDefinitionBuilder()
}

// WithApplication adds an application to the builder
func (app *languageDefinitionBuilder) WithApplication(appli LanguageApplication) LanguageDefinitionBuilder {
	app.app = appli
	return app
}

// WithPatternMatches add pattern matches to the builder
func (app *languageDefinitionBuilder) WithPatternMatches(matches []definitions.PatternMatch) LanguageDefinitionBuilder {
	app.matches = matches
	return app
}

// WithPaths adds a paths to the builder
func (app *languageDefinitionBuilder) WithPaths(paths Paths) LanguageDefinitionBuilder {
	app.paths = paths
	return app
}

// WithRoot adds a root to the builder
func (app *languageDefinitionBuilder) WithRoot(root string) LanguageDefinitionBuilder {
	app.root = root
	return app
}

// Now builds a new LanguageDefinition instance
func (app *languageDefinitionBuilder) Now() (LanguageDefinition, error) {
	if app.app == nil {
		return nil, errors.New("the application is mandatory in order to build a LanguageDefinition instance")
	}

	if app.paths == nil {
		return nil, errors.New("the Paths instance is mandatory in order to build a LanguageDefinition instance")
	}

	if app.root == "" {
		return nil, errors.New("the root pattern is mandatory in order to build a LanguageDefinition instance")
	}

	if app.matches != nil && len(app.matches) <= 0 {
		app.matches = nil
	}

	if app.matches == nil {
		return nil, errors.New("the []PatternMatch are mandatory in order to build a LanguageDefinition instance")
	}

	return createLanguageDefinition(app.app, app.matches, app.paths, app.root), nil
}
