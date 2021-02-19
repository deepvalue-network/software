package shareholders

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

// Builder represents a shareholder builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithGovernment(gov governments.Government) Builder
	WithPublic(public shareholders.ShareHolder) Builder
	WithSignaturePK(sigPK signature.PrivateKey) Builder
	WithEncryptionPK(encPK encryption.PrivateKey) Builder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Hash() hash.Hash
	ID() *uuid.UUID
	Government() governments.Government
	Public() shareholders.ShareHolder
	Signature() signature.PrivateKey
	Encryption() encryption.PrivateKey
}
