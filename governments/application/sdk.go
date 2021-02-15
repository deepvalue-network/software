package application

import (
	"github.com/deepvalue-network/software/governments/application/authenticated"
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/governments/domain/resolutions"
	"github.com/deepvalue-network/software/governments/domain/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Application represents the government application
type Application interface {
	Retrieve(id *uuid.UUID) (governments.Government, error)
	Proposition(governmentID *uuid.UUID) Proposition
	Resolution(governmentID *uuid.UUID) Resolution
	ShareHolder(governmentID *uuid.UUID) ShareHolder
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

// ShareHolder represents the shareholder application
type ShareHolder interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (shareholders.ShareHolder, error)
	Authenticate(hash hash.Hash, pk signature.PrivateKey) authenticated.Application
}
