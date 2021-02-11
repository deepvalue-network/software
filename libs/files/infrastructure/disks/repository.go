package disks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hydro"
	uuid "github.com/satori/go.uuid"
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

// List lists the hashes
func (app *repository) List() ([]hash.Hash, error) {
	return listFilesForHashes(app.basePath)
}

// ListIDs list the ids
func (app *repository) ListIDs() ([]*uuid.UUID, error) {
	return listFilesForIDs(app.basePath)
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

	fmt.Printf("\n%s\n", js)

	return app.hydroAdapter.Dehydrate(hydrated)
}
