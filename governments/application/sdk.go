package application

import (
	"github.com/deepvalue-network/software/governments/application/authenticated"
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/resolutions"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Application represents the government application
type Application interface {
	Identity() Identity
	Government() Government
}

// Government represents the government application
type Government interface {
	List() ([]*uuid.UUID, error)
	Retrieve(id *uuid.UUID) (governments.Government, error)
	Proposition(id *uuid.UUID) Proposition
	Resolution(id *uuid.UUID) Resolution
}

// Proposition represents the proposition application
type Proposition interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) propositions.Proposition
}

// Resolution represents a resolution
type Resolution interface {
	List() ([]hash.Hash, error)
	Pending() ([]hash.Hash, error)
	Passed() ([]hash.Hash, error)
	Denied() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (resolutions.Resolution, error)
}

// Identity represents the identity application
type Identity interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (shareholders.ShareHolder, error)
	Authenticate(name string, seed string, password string) authenticated.Application
}
