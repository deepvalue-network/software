package shareholders

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a shareholder builder
type Builder interface {
	Create() Builder
	WithGovernment(gov governments.Government) Builder
	WithKeys(keys []hash.Hash) Builder
	WithPower(power uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Hash() hash.Hash
	Government() governments.Government
	Keys() []hash.Hash
	Power() uint
	CreatedOn() time.Time
}
