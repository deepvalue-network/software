package files

import (
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Repository represents a file repository
type Repository interface {
	List() ([]hash.Hash, error)
	ListIDs() ([]*uuid.UUID, error)
	Retrieve(name string) (interface{}, error)
}

// Service represents a file service
type Service interface {
	Insert(name string, data interface{}) error
	Update(name string, data interface{}) error
	Delete(name string) error
}
