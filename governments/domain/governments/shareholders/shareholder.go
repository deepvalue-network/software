package shareholders

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

type shareHolder struct {
	hash      hash.Hash
	chain     chains.Chain
	keys      []hash.Hash
	power     uint
	createdOn time.Time
}

func createShareHolder(
	hash hash.Hash,
	chain chains.Chain,
	keys []hash.Hash,
	power uint,
	createdOn time.Time,
) ShareHolder {
	out := shareHolder{
		hash:      hash,
		chain:     chain,
		keys:      keys,
		power:     power,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *shareHolder) Hash() hash.Hash {
	return obj.hash
}

// Chain returns the chain
func (obj *shareHolder) Chain() chains.Chain {
	return obj.chain
}

// Keys returns the keys
func (obj *shareHolder) Keys() []hash.Hash {
	return obj.keys
}

// Power returns the power
func (obj *shareHolder) Power() uint {
	return obj.power
}

// CreatedOn returns the creation time
func (obj *shareHolder) CreatedOn() time.Time {
	return obj.createdOn
}

// Same retruns true if the given hashes are all contained inside the shareholder, and the length of the pubKeys are the same as the shareholder's keys
func (obj *shareHolder) Same(pubKeyHashes []hash.Hash) bool {
	if len(obj.keys) != len(pubKeyHashes) {
		return false
	}

	for _, oneHash := range pubKeyHashes {
		if !obj.Contains(oneHash) {
			return false
		}
	}

	return true
}

// Contains returns true if the given hash is contained within the hashed keys, false otherwise
func (obj *shareHolder) Contains(hashedPubKey hash.Hash) bool {
	for _, oneKey := range obj.keys {
		if oneKey.Compare(hashedPubKey) {
			return true
		}
	}

	return false
}
