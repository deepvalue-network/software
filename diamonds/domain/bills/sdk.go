package bills

import (
	"time"

	"github.com/steve-care-software/products/diamonds/domain/genesis/spends/views"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

// Bill represents a bill
type Bill interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.RingSignature
}

// Content represents a bill content
type Content interface {
	Hash() hash.Hash
	Origin() Origin
	Amount() []byte
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

// ViewBill represents a view bill
type ViewBill interface {
	Hash() hash.Hash
	Bill() Bill
	Seed() string
	Amount() uint64
}
