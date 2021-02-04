package files

import "github.com/steve-care-software/products/libs/hash"

// Repository represents a file repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(name string) (interface{}, error)
}

// Service represents a file service
type Service interface {
	Insert(name string, data interface{}) error
	Update(name string, data interface{}) error
	Delete(name string) error
}
