package owners

import (
	"github.com/deepvalue-network/software/diamonds/domain/bills"
	"github.com/deepvalue-network/software/diamonds/domain/owners"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuider creates a new builder instance
func NewBuider() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a bill builder
type Builder interface {
	Create() Builder
	WithOwner(owner owners.Owner) Builder
	WithViewBill(viewBill bills.ViewBill) Builder
	Now() (Bill, error)
}

// Bill represents an owned bill
type Bill interface {
	Hash() hash.Hash
	Owner() owners.Owner
	Bill() bills.ViewBill
}

// Repository represents a bill repository
type Repository interface {
	List(owner owners.Owner) []hash.Hash
	Retrieve(owner owners.Owner, hash hash.Hash) (Bill, error)
}

// Service represents a bill service
type Service interface {
	Insert(bill Bill) error
	Delete(bill Bill) error
}
