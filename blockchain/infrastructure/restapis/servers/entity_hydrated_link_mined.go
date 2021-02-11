package servers

import (
	"time"
)

type entityHydratedLinkMined struct {
	Link      *entityHydratedLink `json:"link" hydro:"0"`
	Results   string              `json:"results" hydro:"1"`
	CreatedOn string              `json:"created_on" hydro:"2"`
}

func linkMinedOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if createdOn, ok := ins.(time.Time); ok {
		return createdOn.Format(internalTimeLayout), nil
	}

	return nil, nil
}

func linkMinedOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "CreatedOn" {
		createdOn, err := time.Parse(internalTimeLayout, ins.(string))
		if err != nil {
			return nil, err
		}

		return createdOn, nil
	}

	return nil, nil
}
