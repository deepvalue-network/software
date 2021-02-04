package bills

import (
	"time"

	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents a bill builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(signature signature.RingSignature) Builder
	Now() (Bill, error)
}

// Bill represents a bill
type Bill interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// ContentBuilder represents a bill content builder
type ContentBuilder interface {
	Create() Builder
	WithOriginGenesis(originGenesis views.Genesis) Builder
	WithOriginBill(originBill ViewBill) Builder
	WithHashedAmount(hashedAmount hash.Hash) Builder
	WithEncryptedSeed(encSeed []byte) Builder
	WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Bill, error)
}

// Content represents a bill content
type Content interface {
	Hash() hash.Hash
	Origin() Origin
	Amount() hash.Hash
	Seed() []byte
	Owner() []hash.Hash
	CreatedOn() time.Time
}

// Origin represents a bill origin
type Origin interface {
	Hash() hash.Hash
	IsGenesis() bool
	Genesis() views.Genesis
	IsBill() bool
	Bill() ViewBill
}

// ViewBillBuilder represents a view bill builder
type ViewBillBuilder interface {
	Create() ViewBillBuilder
	WithBill(bill Bill) ViewBillBuilder
	WithSeed(seed string) ViewBillBuilder
	Now() (ViewBill, error)
}

// ViewBill represents a view bill
type ViewBill interface {
	Hash() hash.Hash
	Bill() Bill
	Seed() string
	Amount() uint64
}

// Repository represents a bill repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Bill, error)
}

// Service represents a bill service
type Service interface {
	Save(bill Bill) error
	Delete(bill Bill) error
}

// ViewRepository represents a view bill repository
type ViewRepository interface {
	List() ([]hash.Hash, error)
	ListByBill(bill Bill) ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (ViewBill, error)
}

// ViewService represents a view bill service
type ViewService interface {
	Insert(bill ViewBill) error
	Delete(bill ViewBill) error
}
