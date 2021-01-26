package file

import (
	"io/ioutil"
	"path/filepath"
)

type fileDiskRepository struct {
	basePath string
}

func createFileDiskRepository(basePath string) Repository {
	out := fileDiskRepository{
		basePath: basePath,
	}

	return &out
}

// Retrieve retrieves data from file using its name
func (app *fileDiskRepository) Retrieve(relativePath string) ([]byte, error) {
	path := filepath.Join(app.basePath, relativePath)
	encrypted, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

// RetrieveAll retrieves all files in a given directory
func (app *fileDiskRepository) RetrieveAll(relativePath string) ([]string, error) {
	path := filepath.Join(app.basePath, relativePath)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	out := []string{}
	for _, oneFile := range files {
		if oneFile.IsDir() {
			continue
		}

		out = append(out, oneFile.Name())
	}

	return out, nil
}
