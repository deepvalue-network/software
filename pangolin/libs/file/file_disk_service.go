package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileDiskService struct {
	basePath string
}

func createFileDiskService(basePath string) Service {
	out := fileDiskService{
		basePath: basePath,
	}

	return &out
}

// Save saves content on disk
func (app *fileDiskService) Save(relativePath string, content []byte) error {
	path := filepath.Join(app.basePath, relativePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dir := filepath.Dir(path)
		os.MkdirAll(dir, 0777)
	}

	return ioutil.WriteFile(path, []byte(content), 0777)
}

// Delete deletes content from disk
func (app *fileDiskService) Delete(relativePath string) error {
	path := filepath.Join(app.basePath, relativePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		str := fmt.Sprintf("the given relative path (%s) does not exists", relativePath)
		return errors.New(str)
	}

	return os.Remove(path)
}

// DeleteAll deletes all content from disk
func (app *fileDiskService) DeleteAll(relativePath []string) error {
	for _, onePath := range relativePath {
		err := app.Delete(onePath)
		if err != nil {
			return err
		}
	}

	return nil
}
