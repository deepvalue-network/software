package identities

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type shareHolder struct {
	hash   hash.Hash
	gov    governments.Government
	public shareholders.ShareHolder
	sigPK  signature.PrivateKey
}

func createShareHolder(
	hash hash.Hash,
	gov governments.Government,
	public shareholders.ShareHolder,
	sigPK signature.PrivateKey,
) ShareHolder {
	out := shareHolder{
		hash:   hash,
		gov:    gov,
		public: public,
		sigPK:  sigPK,
	}

	return &out
}

// Hash returns the hash
func (obj *shareHolder) Hash() hash.Hash {
	return obj.hash
}

// Government returns the government
func (obj *shareHolder) Government() governments.Government {
	return obj.gov
}

// Public returns the public shareholder
func (obj *shareHolder) Public() shareholders.ShareHolder {
	return obj.public
}

// SigPK returns the signature privateKey
func (obj *shareHolder) SigPK() signature.PrivateKey {
	return obj.sigPK
}
