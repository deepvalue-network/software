package mined

import "github.com/deepvalue-network/software/libs/hash"

func minerHash(results string, hash hash.Hash, hashAdapter hash.Adapter) (*hash.Hash, error) {
	return hashAdapter.FromMultiBytes([][]byte{
		[]byte(results),
		hash.Bytes(),
	})
}
