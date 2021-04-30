package linkers

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"

type languageDefinition struct {
	app     LanguageApplication
	matches []definitions.PatternMatch
	paths   Paths
	root    string
}

func createLanguageDefinition(
	app LanguageApplication,
	matches []definitions.PatternMatch,
	paths Paths,
	root string,
) LanguageDefinition {
	out := languageDefinition{
		app:     app,
		matches: matches,
		paths:   paths,
		root:    root,
	}

	return &out
}

// Application return the application
func (obj *languageDefinition) Application() LanguageApplication {
	return obj.app
}

// PatternMatches return the pattern matches
func (obj *languageDefinition) PatternMatches() []definitions.PatternMatch {
	return obj.matches
}

// Paths return the paths
func (obj *languageDefinition) Paths() Paths {
	return obj.paths
}

// Root return the root pattern
func (obj *languageDefinition) Root() string {
	return obj.root
}
