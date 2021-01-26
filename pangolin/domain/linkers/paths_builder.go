package linkers

import "errors"

type pathsBuilder struct {
	baseDir  string
	tokens   string
	rules    string
	logics   string
	channels string
}

func createPathsBuilder() PathsBuilder {
	out := pathsBuilder{
		baseDir:  "",
		tokens:   "",
		rules:    "",
		logics:   "",
		channels: "",
	}

	return &out
}

// Create initializes the builder
func (app *pathsBuilder) Create() PathsBuilder {
	return createPathsBuilder()
}

// WithBaseDir adds a baseDir to the builder
func (app *pathsBuilder) WithBaseDir(baseDir string) PathsBuilder {
	app.baseDir = baseDir
	return app
}

// WithTokens add tokens to the builder
func (app *pathsBuilder) WithTokens(tokens string) PathsBuilder {
	app.tokens = tokens
	return app
}

// WithRules add rules to the builder
func (app *pathsBuilder) WithRules(rules string) PathsBuilder {
	app.rules = rules
	return app
}

// WithLogics add logics to the builder
func (app *pathsBuilder) WithLogics(logics string) PathsBuilder {
	app.logics = logics
	return app
}

// WithChannels add channels to the builder
func (app *pathsBuilder) WithChannels(channels string) PathsBuilder {
	app.channels = channels
	return app
}

// Now builds a new Paths instance
func (app *pathsBuilder) Now() (Paths, error) {
	if app.baseDir == "" {
		return nil, errors.New("the base directory is mandatory in roder to build a Paths instance")
	}
	if app.tokens == "" {
		return nil, errors.New("the tokens path is mandatory in order to build a Paths instance")
	}

	if app.rules == "" {
		return nil, errors.New("the rules path is mandatory in order to build a Paths instance")
	}

	if app.logics == "" {
		return nil, errors.New("the logics path is mandatory in order to build a Paths instance")
	}

	if app.channels != "" {
		return createPathsWithChannels(app.baseDir, app.tokens, app.rules, app.logics, app.channels), nil
	}

	return createPaths(app.baseDir, app.tokens, app.rules, app.logics), nil
}
