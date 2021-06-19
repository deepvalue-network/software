package grammars

import "errors"

type fileBuilder struct {
	root         string
	tokensPath   string
	rulesPath    string
	channelsPath string
}

func createFileBuilder() FileBuilder {
	out := fileBuilder{
		root:         "",
		tokensPath:   "",
		rulesPath:    "",
		channelsPath: "",
	}

	return &out
}

// Create initializes the builder
func (app *fileBuilder) Create() FileBuilder {
	return createFileBuilder()
}

// WithRoot adds a root token to the builder
func (app *fileBuilder) WithRoot(root string) FileBuilder {
	app.root = root
	return app
}

// WithTokensPath adds a tokens path to the builder
func (app *fileBuilder) WithTokensPath(tokensPath string) FileBuilder {
	app.tokensPath = tokensPath
	return app
}

// WithRulesPath adds a rules path to the builder
func (app *fileBuilder) WithRulesPath(rulesPath string) FileBuilder {
	app.rulesPath = rulesPath
	return app
}

// WithChannelsPath adds a channels path to the builder
func (app *fileBuilder) WithChannelsPath(channelsPath string) FileBuilder {
	app.channelsPath = channelsPath
	return app
}

// Now builds a new File instance
func (app *fileBuilder) Now() (File, error) {
	if app.root == "" {
		return nil, errors.New("the root token is mandatory in order to build a File instance")
	}

	if app.tokensPath == "" {
		return nil, errors.New("the tokens path is mandatory in order to build a File instance")
	}

	if app.rulesPath == "" {
		return nil, errors.New("the rules path is mandatory in order to build a File instance")
	}

	if app.channelsPath != "" {
		return createFileWithChannelsPath(app.root, app.tokensPath, app.rulesPath, app.channelsPath), nil
	}

	return createFile(app.root, app.tokensPath, app.rulesPath), nil
}
