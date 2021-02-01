package disks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/steve-care-software/products/identity/domain/users"
	"github.com/steve-care-software/products/libs/cryptography/encryption"
)

type userRepository struct {
	basePath string
}

func createUserRepository(basePath string) users.Repository {
	out := userRepository{
		basePath: basePath,
	}

	return &out
}

// List lists the user names
func (app *userRepository) List() ([]string, error) {
	files, err := ioutil.ReadDir(app.basePath)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		names = append(names, oneFile.Name())
	}

	return names, nil
}

// Retrieve retrieves a user
func (app *userRepository) Retrieve(name string, seed string, password string) (users.User, error) {
	filePath := filepath.Join(app.basePath, name)
	encData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	enc := encryption.NewEncryption(password)
	decryptedData, err := enc.Decrypt(string(encData))
	if err != nil {
		return nil, err
	}

	hydrated := new(hydratedUser)
	err = json.Unmarshal(decryptedData, hydrated)
	if err != nil {
		return nil, err
	}

	dehydrated, err := hydroAdapter.Dehydrate(hydrated)
	if err != nil {
		return nil, err
	}

	if user, ok := dehydrated.(users.User); ok {
		return user, nil
	}

	str := fmt.Sprintf("the user (name: %s) saved on disk is of invalid type", name)
	return nil, errors.New(str)
}
