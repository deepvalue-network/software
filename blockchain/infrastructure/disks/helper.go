package disks

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

func makeDirIfNotExists(path string, fileMode os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, fileMode)
		if err != nil {
			return err
		}
	}

	return nil
}

func blockFilePath(basePath string, block blocks.Block) string {
	return filePath(basePath, block.Tree().Head())
}

func filePath(basePath string, hash hash.Hash) string {
	return filepath.Join(basePath, hash.String())
}

func listFiles(dirPath string) ([]hash.Hash, error) {
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
			return nil, err
		}

		out = append(out, *hsh)
	}

	return out, nil
}
