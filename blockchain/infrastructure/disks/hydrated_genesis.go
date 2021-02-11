package disks

import "github.com/deepvalue-network/software/libs/hash"

// HydratedGenesis represents an hydrated genesis
type HydratedGenesis struct {
	Hash                           string  `json:"hash"`
	MiningValue                    uint8   `json:"mining_value" hydro:"0"`
	BlockBaseDifficulty            uint    `json:"block_base_difficulty" hydro:"1"`
	BlockIncreasePerHashDifficulty float64 `json:"block_increase_per_hash_difficulty" hydro:"2"`
	LinkDifficulty                 uint    `json:"link_difficulty" hydro:"3"`
}

func genesisOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Hash" {
		if hsh, ok := ins.(hash.Hash); ok {
			return hsh.String(), nil
		}
	}

	return nil, nil
}
