package grammar

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
)

type retrieverCriteriaRepository struct {
	fileFetcher FileFetcher
	builder     RetrieverCriteriaBuilder
	name        string
	root        string
}

func createRetrieverCriteriaRepository(
	fileFetcher FileFetcher,
	builder RetrieverCriteriaBuilder,
	name string,
) RetrieverCriteriaRepository {
	return createRetrieverCriteriaRepositoryInternally(fileFetcher, builder, name, "")
}

func createRetrieverCriteriaRepositoryWithRoot(
	fileFetcher FileFetcher,
	builder RetrieverCriteriaBuilder,
	name string,
	root string,
) RetrieverCriteriaRepository {
	return createRetrieverCriteriaRepositoryInternally(fileFetcher, builder, name, root)
}

func createRetrieverCriteriaRepositoryInternally(
	fileFetcher FileFetcher,
	builder RetrieverCriteriaBuilder,
	name string,
	root string,
) RetrieverCriteriaRepository {
	out := retrieverCriteriaRepository{
		fileFetcher: fileFetcher,
		builder:     builder,
		name:        name,
		root:        root,
	}

	return &out
}

// Retrieve retrieves a retrieverCriteria instance from path
func (app *retrieverCriteriaRepository) Retrieve(filePath string) (RetrieverCriteria, error) {
	return app.retrieveChild(app.root, app.name, filePath)
}

func (app *retrieverCriteriaRepository) retrieveChild(root string, name string, filePath string) (RetrieverCriteria, error) {
	// read the file:
	js, err := app.fileFetcher(filePath)
	if err != nil {
		return nil, err
	}

	// convert to json:
	lex := new(grammarJSON)
	err = json.Unmarshal(js, lex)
	if err != nil {
		str := fmt.Sprintf("there was an error converting the file's data (path: %s) to json: %s", filePath, err.Error())
		return nil, errors.New(str)
	}

	dirPath := filepath.Dir(filePath)
	builder := app.builder.Create().
		WithName(name).
		WithBaseDirPath(dirPath).
		WithRoot(lex.Root).
		WithTokensPath(lex.Tokens).
		WithRulesPath(lex.Rules)

	if root != "" {
		builder.WithRoot(root)
	}

	if lex.Channels != "" {
		builder.WithChannelsPath(lex.Channels)
	}

	if lex.Extends != nil {
		criterias := []RetrieverCriteria{}
		for name, path := range lex.Extends {
			criteria, err := app.retrieveChild("", name, path)
			if err != nil {
				return nil, err
			}

			criterias = append(criterias, criteria)
		}

		builder.WithExtends(criterias)
	}

	return builder.Now()
}
