package shareholders

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
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
	Hash() hash.Hash
	Government() governments.Government
	Public() shareholders.ShareHolder
	SigPK() signature.PrivateKey
}
