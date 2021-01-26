package linkers

import (
	"github.com/steve-care-software/products/pangolin/domain/middle"
)

type language struct {
	app     Application
	matches []middle.PatternMatch
	paths   Paths
	root    string
}

func createLanguage(
	app Application,
	matches []middle.PatternMatch,
	paths Paths,
	root string,
) Language {
	out := language{
		app:     app,
		matches: matches,
		paths:   paths,
		root:    root,
	}

	return &out
}

// Application return the application
func (obj *language) Application() Application {
	return obj.app
}

// PatternMatches return the pattern matches
func (obj *language) PatternMatches() []middle.PatternMatch {
	return obj.matches
}

// Paths return the paths
func (obj *language) Paths() Paths {
	return obj.paths
}

// Root return the root pattern
func (obj *language) Root() string {
	return obj.root
}
