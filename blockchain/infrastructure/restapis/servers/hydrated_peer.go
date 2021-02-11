package servers

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
)

type hydratedPeer struct {
	Content       string `json:"server" hydro:"0"`
	CreatedOn     string `json:"created_on" hydro:"1"`
	LastUpdatedOn string `json:"last_updated_on,omitempty" hydro:"2"`
}

func peerOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "CreatedOn" {
		if createdOn, ok := ins.(time.Time); ok {
			return createdOn.Format(internalTimeLayout), nil
		}
	}

	if fieldName == "LastUpdatedOn" {
		if lastUpdatedOn, ok := ins.(time.Time); ok {
			return lastUpdatedOn.Format(internalTimeLayout), nil
		}
	}

	if fieldName == "Content" {
		if content, ok := ins.(peers.Content); ok {
			return content.String(), nil
		}
	}

	return nil, nil
}

func peerOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "CreatedOn" {
		createdOn, err := time.Parse(internalTimeLayout, ins.(string))
		if err != nil {
			return nil, err
		}

		return createdOn, nil
	}

	if fieldName == "LastUpdatedOn" {
		if str, ok := ins.(string); ok {
			if str == "" {
				return nil, nil
			}

			lastUpdatedOn, err := time.Parse(internalTimeLayout, str)
			if err != nil {
				return nil, err
			}

			return &lastUpdatedOn, nil
		}
	}

	return nil, nil
}
