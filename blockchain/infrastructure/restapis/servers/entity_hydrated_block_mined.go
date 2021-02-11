package servers

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
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
		return createdOn.Format(internalTimeLayout), nil
	}

	return nil, nil
}

func blockMinedOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "CreatedOn" {
		createdOn, err := time.Parse(internalTimeLayout, ins.(string))
		if err != nil {
			return nil, err
		}

		return createdOn, nil
	}

	return nil, nil
}
