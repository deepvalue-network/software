package disks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/products/libs/files/domain/files"
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hydro"
)

type repository struct {
	hydroAdapter hydro.Adapter
	basePath     string
	ptr          interface{}
}

func createRepository(
	hydroAdapter hydro.Adapter,
	basePath string,
	ptr interface{},
) files.Repository {
	out := repository{
		hydroAdapter: hydroAdapter,
		basePath:     basePath,
		ptr:          ptr,
	}

	return &out
}

// List lists the hashes of head blocks
func (app *repository) List() ([]hash.Hash, error) {
	return listFiles(app.basePath)
}

// Retrieve retrieves a file by name
func (app *repository) Retrieve(name string) (interface{}, error) {
	path := filepath.Join(app.basePath, name)
	if !fileExists(path) {
		str := fmt.Sprintf(fileDoesNotExistsPattern, path)
		return nil, errors.New(str)
	}

	js, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hydrated := app.ptr
	err = json.Unmarshal(js, hydrated)
	if err != nil {
		return nil, err
	}

	return app.hydroAdapter.Dehydrate(hydrated)
}
