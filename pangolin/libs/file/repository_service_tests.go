package file

import (
	"errors"
	"fmt"
	"path/filepath"
)

// RepositoryServiceForTests represents a repository and service for tests
type RepositoryServiceForTests struct {
	data map[string][]byte
}

// CreateRepositoryServiceForTests creates a repository and service for tests
func CreateRepositoryServiceForTests() *RepositoryServiceForTests {
	out := RepositoryServiceForTests{
		data: map[string][]byte{},
	}

	return &out
}

// Retrieve retrieves data for path
func (app *RepositoryServiceForTests) Retrieve(relativePath string) ([]byte, error) {
	if dat, ok := app.data[relativePath]; ok {
		return dat, nil
	}

	str := fmt.Sprintf("the relativePath (%s) does not exists", relativePath)
	return nil, errors.New(str)
}

// RetrieveAll retrieves all files in a given directory
func (app *RepositoryServiceForTests) RetrieveAll(relativePath string) ([]string, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, err
	}

	out := []string{}
	for oneRelFilePath := range app.data {
		absDir, err := filepath.Abs(filepath.Dir(oneRelFilePath))
		if err != nil {
			return nil, err
		}

		if absDir == absPath {
			out = append(out, oneRelFilePath)
		}
	}

	return out, nil
}

// Save saves data on path
func (app *RepositoryServiceForTests) Save(relativePath string, content []byte) error {
	app.data[relativePath] = content
	return nil
}

// Delete deletes content
func (app *RepositoryServiceForTests) Delete(relativePath string) error {
	if _, ok := app.data[relativePath]; ok {
		delete(app.data, relativePath)
		return nil
	}

	str := fmt.Sprintf("the given relativePath (%s) is invalid", relativePath)
	return errors.New(str)
}

// DeleteAll deletes all content
func (app *RepositoryServiceForTests) DeleteAll(relativePath []string) error {
	for _, onePath := range relativePath {
		err := app.Delete(onePath)
		if err != nil {
			return err
		}
	}

	return nil
}
