package shareholders

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/governments/domain/shareholders/transactions"
)

// Builder represents a shareholder builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSeed(seed string) Builder
	WithGovernment(gov governments.Government) Builder
	WithPublic(public shareholders.ShareHolder) Builder
	WithTransactions(transactions transactions.Transactions) Builder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Name() string
	Seed() string
	Government() governments.Government
	Public() shareholders.ShareHolder
	Transactions() transactions.Transactions
}
