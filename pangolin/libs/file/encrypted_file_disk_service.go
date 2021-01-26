package file

import (
	"github.com/steve-care-software/products/pangolin/libs/cryptography/encryption"
)

type encryptedFileDiskService struct {
	encryption encryption.Encryption
	service    Service
}

func createEncryptedFileDiskService(encryption encryption.Encryption, service Service) Service {
	out := encryptedFileDiskService{
		encryption: encryption,
		service:    service,
	}

	return &out
}

// Save saves content on disk
func (app *encryptedFileDiskService) Save(relativePath string, content []byte) error {
	encrypted, err := app.encryption.Encrypt(content)
	if err != nil {
		return err
	}

	return app.service.Save(relativePath, []byte(encrypted))
}

// Delete deletes content from disk
func (app *encryptedFileDiskService) Delete(relativePath string) error {
	return app.service.Delete(relativePath)
}

// DeleteAll deletes all content from disk
func (app *encryptedFileDiskService) DeleteAll(relativePath []string) error {
	return app.service.DeleteAll(relativePath)
}
