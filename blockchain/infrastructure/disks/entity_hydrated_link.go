package disks

import (
	"github.com/deepvalue-network/software/libs/hash"
)

// EntityHydratedLink represents an entity hydrated link
type EntityHydratedLink struct {
	Hash          string `json:"hash"`
	Index         uint   `json:"index" hydro:"0"`
	PrevMinedLink string `json:"prev_mined_link" hydro:"1"`
	NextBlock     string `json:"next_block" hydro:"2"`
}

func linkOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Hash" {
		if hsh, ok := ins.(hash.Hash); ok {
			return hsh.String(), nil
		}
	}

	if fieldName == "NextBlock" {
		if hydratedBlock, ok := ins.(*EntityHydratedBlock); ok {
			return hydratedBlock.Tree.Head, nil
		}
	}

	if fieldName == "PrevMinedLink" {
		if hsh, ok := ins.(hash.Hash); ok {
			return hsh.String(), nil
		}
	}

	return nil, nil
}

func linkOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "NextBlock" {
		if strHash, ok := ins.(string); ok {
			hsh, err := hash.NewAdapter().FromString(strHash)
			if err != nil {
				return nil, err
			}

			return internalRepositoryBlock.Retrieve(*hsh)
		}
	}

	if fieldName == "PrevMinedLink" {
		hsh, err := hash.NewAdapter().FromString(ins.(string))
		if err != nil {
			return nil, err
		}

		return *hsh, nil
	}

	return nil, nil
}
