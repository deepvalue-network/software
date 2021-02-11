package disks

import (
	"time"

	"github.com/deepvalue-network/software/libs/hash"
)

// EntityHydratedBlockMined represents an entity hydrated mined block
type EntityHydratedBlockMined struct {
	Hash      string `json:"hash"`
	Block     string `json:"block" hydro:"0"`
	Results   string `json:"results" hydro:"1"`
	CreatedOn string `json:"created_on" hydro:"2"`
}

func blockMinedOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Hash" {
		if hsh, ok := ins.(hash.Hash); ok {
			return hsh.String(), nil
		}
	}

	if fieldName == "Block" {
		if hydratedBlock, ok := ins.(*EntityHydratedBlock); ok {
			return hydratedBlock.Tree.Head, nil
		}
	}

	if fieldName == "CreatedOn" {
		if createdOn, ok := ins.(time.Time); ok {
			return createdOn.Format(timeLayout), nil
		}
	}

	return nil, nil
}

func blockMinedOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Block" {
		if strHash, ok := ins.(string); ok {
			hsh, err := hash.NewAdapter().FromString(strHash)
			if err != nil {
				return nil, err
			}

			return internalRepositoryBlock.Retrieve(*hsh)
		}
	}

	if fieldName == "CreatedOn" {
		createdOn, err := time.Parse(timeLayout, ins.(string))
		if err != nil {
			return nil, err
		}

		return createdOn, nil
	}

	return nil, nil
}
