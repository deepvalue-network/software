package grammar

import "errors"

type retrieverCriteriaBuilder struct {
	name         string
	root         string
	baseDirPath  string
	tokensPath   string
	channelsPath string
	rulesPath    string
	extends      []RetrieverCriteria
}

func createRetrieverCriteriaBuilder() RetrieverCriteriaBuilder {
	out := retrieverCriteriaBuilder{
		name:         "",
		root:         "",
		baseDirPath:  "",
		tokensPath:   "",
		channelsPath: "",
		rulesPath:    "",
		extends:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *retrieverCriteriaBuilder) Create() RetrieverCriteriaBuilder {
	return createRetrieverCriteriaBuilder()
}

// WithName adds a name to the builder
func (app *retrieverCriteriaBuilder) WithName(name string) RetrieverCriteriaBuilder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *retrieverCriteriaBuilder) WithRoot(root string) RetrieverCriteriaBuilder {
	app.root = root
	return app
}

// WithBaseDirPath adds a baseDir path to the builder
func (app *retrieverCriteriaBuilder) WithBaseDirPath(baseDirPath string) RetrieverCriteriaBuilder {
	app.baseDirPath = baseDirPath
	return app
}

// WithTokensPath adds a tokens path to the builder
func (app *retrieverCriteriaBuilder) WithTokensPath(tokensPath string) RetrieverCriteriaBuilder {
	app.tokensPath = tokensPath
	return app
}

// WithChannelsPath adds a channels path to the builder
func (app *retrieverCriteriaBuilder) WithChannelsPath(channelsPath string) RetrieverCriteriaBuilder {
	app.channelsPath = channelsPath
	return app
}

// WithRulesPath adds a rules path to the builder
func (app *retrieverCriteriaBuilder) WithRulesPath(rulesPath string) RetrieverCriteriaBuilder {
	app.rulesPath = rulesPath
	return app
}

// WithExtends adds an extends path to the builder
func (app *retrieverCriteriaBuilder) WithExtends(extends []RetrieverCriteria) RetrieverCriteriaBuilder {
	app.extends = extends
	return app
}

// Now builds a new RetrieverCriteria instance
func (app *retrieverCriteriaBuilder) Now() (RetrieverCriteria, error) {
	if app.name == "" {
		app.name = defaultName
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build a RetrieverCriteria instance")
	}

	if app.baseDirPath == "" {
		return nil, errors.New("the base directory path is mandatory in order to build a RerieverCriteria instance")
	}

	if app.tokensPath == "" {
		return nil, errors.New("the tokens path is mandatory in order to build a RetrieverCriteria instance")
	}

	if app.extends != nil {
		extends := map[string]RetrieverCriteria{}
		for _, oneExtend := range extends {
			name := oneExtend.Name()
			extends[name] = oneExtend
		}

		if app.channelsPath != "" {
			return createRetrieverCriteriaWithChannelsPathAndExtends(
				app.name,
				app.root,
				app.baseDirPath,
				app.tokensPath,
				app.rulesPath,
				app.channelsPath,
				extends,
			), nil
		}

		return createRetrieverCriteriaWithExtends(
			app.name,
			app.root,
			app.baseDirPath,
			app.tokensPath,
			app.rulesPath,
			extends,
		), nil
	}

	if app.channelsPath != "" {
		return createRetrieverCriteriaWithChannelsPath(
			app.name,
			app.root,
			app.baseDirPath,
			app.tokensPath,
			app.rulesPath,
			app.channelsPath,
		), nil
	}

	return createRetrieverCriteria(
		app.name,
		app.root,
		app.baseDirPath,
		app.tokensPath,
		app.rulesPath,
	), nil
}
