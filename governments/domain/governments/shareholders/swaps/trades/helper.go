package trades

import "github.com/deepvalue-network/software/libs/hash"

func compareHashes(origin []hash.Hash, compareTo []hash.Hash) bool {
	if len(origin) != len(compareTo) {
		return false
	}

	for _, oneCompareTo := range compareTo {
		for _, oneOrigin := range origin {
			if !oneCompareTo.Compare(oneOrigin) {
				return false
			}
		}
	}

	return true
}
