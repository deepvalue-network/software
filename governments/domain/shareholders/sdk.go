package shareholders

import (
	"hash"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transactions"
)

// Builder represents a shareholders builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithShareHolders(shareHolders []ShareHolder) Builder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	Name() string
	Seed() string
	ShareHolders() []ShareHolder
}

// ShareHolderBuilder represents a shareholder builder
type ShareHolderBuilder interface {
	Create() ShareHolderBuilder
	WithGovernment(gov governments.Government) ShareHolderBuilder
	WithPublic(public shareholders.ShareHolder) ShareHolderBuilder
	WithTransactions(transactions transactions.Transactions) ShareHolderBuilder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Hash() hash.Hash
	Government() governments.Government
	Public() shareholders.ShareHolder
	Transactions() transactions.Transactions
}

// Repository represents a shareholders repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, seed string, password string) (ShareHolders, error)
}

// Service represents a shareholders service
type Service interface {
	Insert(ins ShareHolders, password string) error
	Update(origin ShareHolders, updated ShareHolders, password string) error
	UpdateWithPassword(origin ShareHolders, updated ShareHolders, originalPassword string, updatedPassword string) error
	Delete(ins ShareHolders, password string) error
}
