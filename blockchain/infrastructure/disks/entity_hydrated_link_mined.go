package disks

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hash"
)

type entityHydratedLinkMined struct {
	Link      *entityHydratedLink `json:"link" hydro:"0"`
	Results   string              `json:"results" hydro:"1"`
	CreatedOn string              `json:"created_on" hydro:"2"`
}

func linkMinedOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if link, ok := ins.(links.Link); ok {
		return link.Hash().String(), nil
	}

	if createdOn, ok := ins.(time.Time); ok {
		return createdOn.Format(timeLayout), nil
	}

	return nil, nil
}

func linkMinedOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Link" {
		if hsh, ok := ins.(hash.Hash); ok {
			return internalRepositoryLink.Retrieve(hsh)
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
