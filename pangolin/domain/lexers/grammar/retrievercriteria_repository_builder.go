package grammar

import "errors"

type retrieverCriteriaRepositoryBuilder struct {
	builder     RetrieverCriteriaBuilder
	fileFetcher FileFetcher
	root        string
	name        string
}

func createRetrieverCriteriaRepositoryBuilder(
	builder RetrieverCriteriaBuilder,
) RetrieverCriteriaRepositoryBuilder {
	out := retrieverCriteriaRepositoryBuilder{
		builder:     builder,
		fileFetcher: nil,
		root:        "",
		name:        "",
	}

	return &out
}

// Create initializes the builder
func (app *retrieverCriteriaRepositoryBuilder) Create() RetrieverCriteriaRepositoryBuilder {
	return createRetrieverCriteriaRepositoryBuilder(
		app.builder,
	)
}

// WithName adds a name to the builder
func (app *retrieverCriteriaRepositoryBuilder) WithName(name string) RetrieverCriteriaRepositoryBuilder {
	app.name = name
	return app
}

// WithRoot adds a root to the builder
func (app *retrieverCriteriaRepositoryBuilder) WithRoot(root string) RetrieverCriteriaRepositoryBuilder {
	app.root = root
	return app
}

// WithFileFetcher adds a fileFetcher to the builder
func (app *retrieverCriteriaRepositoryBuilder) WithFileFetcher(fileFetcher FileFetcher) RetrieverCriteriaRepositoryBuilder {
	app.fileFetcher = fileFetcher
	return app
}

// Now buiilds a new RetrieverCriteriaRepository instance
func (app *retrieverCriteriaRepositoryBuilder) Now() (RetrieverCriteriaRepository, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a RetrieverCriteriaRepository instance")
	}

	if app.fileFetcher == nil {
		app.fileFetcher = defaultFetch
	}

	if app.root != "" {
		return createRetrieverCriteriaRepositoryWithRoot(app.fileFetcher, app.builder, app.name, app.root), nil
	}

	return createRetrieverCriteriaRepository(app.fileFetcher, app.builder, app.name), nil
}
