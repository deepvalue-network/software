package disks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hydro"
)

type service struct {
	hydroAdapter hydro.Adapter
	basePath     string
	fileMode     os.FileMode
}

func creatService(
	hydroAdapter hydro.Adapter,
	basePath string,
	fileMode os.FileMode,
) files.Service {
	err := makeDirIfNotExists(basePath, fileMode)
	if err != nil {
		panic(err)
	}

	out := service{
		hydroAdapter: hydroAdapter,
		basePath:     basePath,
		fileMode:     fileMode,
	}

	return &out
}

// Insert inserts a file
func (app *service) Insert(name string, ins interface{}) error {
	path := filepath.Join(app.basePath, name)
	if fileExists(path) {
		str := fmt.Sprintf("the file (path: %s) already exists", path)
		return errors.New(str)
	}

	return app.save(path, ins)
}

// Update updates a file
func (app *service) Update(name string, ins interface{}) error {
	path := filepath.Join(app.basePath, name)
	if !fileExists(path) {
		str := fmt.Sprintf(fileDoesNotExistsPattern, path)
		return errors.New(str)
	}

	return app.save(path, ins)
}

// Delete deletes a file
func (app *service) Delete(name string) error {
	path := filepath.Join(app.basePath, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		str := fmt.Sprintf(fileDoesNotExistsPattern, path)
		return errors.New(str)
	}

	return os.Remove(path)
}

func (app *service) save(path string, ins interface{}) error {
	if str, ok := ins.(string); ok {
		return ioutil.WriteFile(path, []byte(str), app.fileMode)
	}

	if bytes, ok := ins.([]byte); ok {
		return ioutil.WriteFile(path, bytes, app.fileMode)
	}

	hydrated, err := app.hydroAdapter.Hydrate(ins)
	if err != nil {
		return err
	}

	js, err := json.Marshal(hydrated)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, js, app.fileMode)
}
