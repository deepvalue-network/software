package disks

import (
	"time"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hash"
)

type entityHydratedBlockMined struct {
	Block     *entityHydratedBlock `json:"block" hydro:"0"`
	Results   string               `json:"results" hydro:"1"`
	CreatedOn string               `json:"created_on" hydro:"2"`
}

func blockMinedOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if block, ok := ins.(blocks.Block); ok {
		return block.Tree().Head().String(), nil
	}

	if createdOn, ok := ins.(time.Time); ok {
		return createdOn.Format(timeLayout), nil
	}

	return nil, nil
}

func blockMinedOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Block" {
		if hsh, ok := ins.(hash.Hash); ok {
			return internalRepositoryBlock.Retrieve(hsh)
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
