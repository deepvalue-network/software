package file

import (
	"github.com/steve-care-software/products/pangolin/libs/cryptography/encryption"
)

// NewEncryptedFileDiskRepository creates a new encrypted repository that reads from encrypted files on disk
func NewEncryptedFileDiskRepository(password string, basePath string) Repository {
	encryption := encryption.NewEncryption(password)
	repository := NewFileDiskRepository(basePath)
	return createEncryptedFileDiskRepository(encryption, repository)
}

// NewEncryptedFileDiskService creates a new encrypted service that writes encrypted data on disk
func NewEncryptedFileDiskService(password string, basePath string) Service {
	encryption := encryption.NewEncryption(password)
	service := NewFileDiskService(basePath)
	return createEncryptedFileDiskService(encryption, service)
}

// NewFileDiskRepository creates a new repository that reads from files on disk
func NewFileDiskRepository(basePath string) Repository {
	return createFileDiskRepository(basePath)
}

// NewFileDiskService creates a new service that writes data on disk
func NewFileDiskService(basePath string) Service {
	return createFileDiskService(basePath)
}

// Repository represents a file repository
type Repository interface {
	Retrieve(relativePath string) ([]byte, error)
	RetrieveAll(relativePath string) ([]string, error)
}

// Service represents the file service
type Service interface {
	Save(relativePath string, content []byte) error
	Delete(relativePath string) error
	DeleteAll(relativePath []string) error
}
