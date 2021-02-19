package shareholders

import (
	"hash"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/governments/domain/identities/shareholders/transactions"
)

// Builder represents a shareholders builder
type Builder interface {
	Create() Builder
	WithShareHolders(shareHolders []ShareHolder) Builder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	All() []ShareHolder
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
