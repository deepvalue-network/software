package disks

import (
	"io/ioutil"
	"os"

	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

func makeDirIfNotExists(path string, fileMode os.FileMode) error {
	if !fileExists(path) {
		err := os.MkdirAll(path, fileMode)
		if err != nil {
			return err
		}
	}

	return nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func listFilesForIDs(dirPath string) ([]*uuid.UUID, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	out := []*uuid.UUID{}
	for _, oneFile := range files {
		if oneFile.IsDir() {
			continue
		}

		name := oneFile.Name()
		id, err := uuid.FromString(name)
		if err != nil {
			continue
		}

		out = append(out, &id)
	}

	return out, nil
}

func listFilesForHashes(dirPath string) ([]hash.Hash, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	out := []hash.Hash{}
	hashAdapter := hash.NewAdapter()
	for _, oneFile := range files {
		if oneFile.IsDir() {
			continue
		}

		name := oneFile.Name()
		hsh, err := hashAdapter.FromString(name)
		if err != nil {
			continue
		}

		out = append(out, *hsh)
	}

	return out, nil
}
