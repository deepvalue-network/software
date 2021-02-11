package servers

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type entityHydratedChain struct {
	ID        string                    `json:"id" hydro:"0"`
	Peers     *hydratedPeers            `json:"peers" hydro:"1"`
	Genesis   *hydratedGenesis          `json:"genesis" hydro:"2"`
	Root      *entityHydratedBlockMined `json:"root_block_mined_hash" hydro:"3"`
	CreatedOn string                    `json:"created_on" hydro:"4"`
	Head      *entityHydratedLinkMined  `json:"head_mined_link_hash" hydro:"5"`
}

func chainOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "ID" {
		if id, ok := ins.(*uuid.UUID); ok {
			return id.String(), nil
		}
	}

	if fieldName == "CreatedOn" {
		if createdOn, ok := ins.(time.Time); ok {
			return createdOn.Format(internalTimeLayout), nil
		}
	}

	return nil, nil
}

func chainOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "ID" {
		id, err := uuid.FromString(ins.(string))
		if err != nil {
			return nil, err
		}

		return &id, nil
	}

	if fieldName == "CreatedOn" {
		createdOn, err := time.Parse(internalTimeLayout, ins.(string))
		if err != nil {
			return nil, err
		}

		return createdOn, nil
	}

	return nil, nil
}
