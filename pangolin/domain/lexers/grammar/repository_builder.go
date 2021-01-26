package grammar

type repositoryBuilder struct {
	retrieverCriteriaRepositoryBuilder RetrieverCriteriaRepositoryBuilder
	builder                            Builder
	ruleAdapter                        RuleAdapter
	tokensAdapter                      TokensAdapter
	fileFetcher                        FileFetcher
}

func createRepositoryBuilder(
	retrieverCriteriaRepositoryBuilder RetrieverCriteriaRepositoryBuilder,
	builder Builder,
	ruleAdapter RuleAdapter,
	tokensAdapter TokensAdapter,
) RepositoryBuilder {
	out := repositoryBuilder{
		retrieverCriteriaRepositoryBuilder: retrieverCriteriaRepositoryBuilder,
		builder:                            builder,
		ruleAdapter:                        ruleAdapter,
		tokensAdapter:                      tokensAdapter,
		fileFetcher:                        nil,
	}

	return &out
}

// Create initializes the builder
func (app *repositoryBuilder) Create() RepositoryBuilder {
	return createRepositoryBuilder(
		app.retrieverCriteriaRepositoryBuilder,
		app.builder,
		app.ruleAdapter,
		app.tokensAdapter,
	)
}

// WithFileFetcher adds a fileFetcher to the builder
func (app *repositoryBuilder) WithFileFetcher(fileFetcher FileFetcher) RepositoryBuilder {
	app.fileFetcher = fileFetcher
	return app
}

// Now builds a new repository instance
func (app *repositoryBuilder) Now() (Repository, error) {
	if app.fileFetcher == nil {
		app.fileFetcher = defaultFetch
	}

	return createRepository(
		app.retrieverCriteriaRepositoryBuilder,
		app.fileFetcher,
		app.builder,
		app.ruleAdapter,
		app.tokensAdapter,
	), nil
}
