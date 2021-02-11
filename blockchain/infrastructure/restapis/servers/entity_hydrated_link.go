package servers

import (
	"github.com/deepvalue-network/software/libs/hash"
)

type entityHydratedLink struct {
	Index         uint                 `json:"index" hydro:"0"`
	PrevMinedLink string               `json:"prev_mined_link" hydro:"1"`
	NextBlock     *entityHydratedBlock `json:"next_block" hydro:"2"`
}

func linkOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if hsh, ok := ins.(hash.Hash); ok {
		return hsh.String(), nil
	}

	return nil, nil
}

func linkOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "PrevMinedLink" {
		hsh, err := hash.NewAdapter().FromString(ins.(string))
		if err != nil {
			return nil, err
		}

		return *hsh, nil
	}

	return nil, nil
}
