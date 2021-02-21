package identities

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewShareHoldersBuilder creates a new shareHolders builder instance
func NewShareHoldersBuilder() ShareHoldersBuilder {
	return createShareHoldersBuilder()
}

// NewShareHolderBuilder creates a new shareHolder builder instance
func NewShareHolderBuilder() ShareHolderBuilder {
	hashAdapter := hash.NewAdapter()
	return createShareHolderBuilder(hashAdapter)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithShareHolders(shareHolders ShareHolders) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Seed() string
	ShareHolders() ShareHolders
}

// ShareHoldersBuilder represents a shareholders builder
type ShareHoldersBuilder interface {
	Create() ShareHoldersBuilder
	WithShareHolders(shareHolders []ShareHolder) ShareHoldersBuilder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	All() []ShareHolder
	Fetch(gov governments.Government) (ShareHolder, error)
}

// ShareHolderBuilder represents a shareholder builder
type ShareHolderBuilder interface {
	Create() ShareHolderBuilder
	WithGovernment(gov governments.Government) ShareHolderBuilder
	WithPublic(public shareholders.ShareHolder) ShareHolderBuilder
	WithSigPK(sigPK signature.PrivateKey) ShareHolderBuilder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Government() governments.Government
	Public() shareholders.ShareHolder
	SigPK() signature.PrivateKey
}

// Repository represents an identity repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (Identity, error)
}

// Service represents a shareholders service
type Service interface {
	Insert(ins Identity, password string) error
	Update(origin Identity, updated Identity, password string) error
	UpdateWithPassword(origin Identity, updated Identity, originalPassword string, updatedPassword string) error
	Delete(ins Identity, password string) error
}
