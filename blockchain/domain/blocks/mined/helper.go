package mined

import "github.com/deepvalue-network/software/libs/hash"

func calculateDifficulty(baseDifficulty uint, incrPerHash float64, amountHashes int) uint {
	sum := float64(0)
	base := float64(baseDifficulty)
	for i := 0; i < int(amountHashes); i++ {
		sum += incrPerHash
	}

	return uint(sum + base)
}

func minerHash(results string, hash hash.Hash, hashAdapter hash.Adapter) (*hash.Hash, error) {
	return hashAdapter.FromMultiBytes([][]byte{
		[]byte(results),
		hash.Bytes(),
	})
}
