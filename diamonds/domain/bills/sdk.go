package bills

import (
	"time"

	"github.com/deepvalue-network/software/diamonds/domain/genesis/spends/views"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder(minPubKeysAmount uint) ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter, minPubKeysAmount)
}

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
	Create() ContentBuilder
	WithOriginGenesis(originGenesis views.Genesis) ContentBuilder
	WithOriginBill(originBill ViewBill) ContentBuilder
	WithAmount(amount uint64) ContentBuilder
	WithHashedAmount(hashedAmount hash.Hash) ContentBuilder
	WithSeed(seed string) ContentBuilder
	WithHashedSeed(hashedSeed hash.Hash) ContentBuilder
	WithPubKeysOwner(pubKeys []signature.PublicKey) ContentBuilder
	WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a bill content
type Content interface {
	Hash() hash.Hash
	Origin() Origin
	Amount() hash.Hash
	Seed() hash.Hash
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
	WithAmount(amount uint64) ViewBillBuilder
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
