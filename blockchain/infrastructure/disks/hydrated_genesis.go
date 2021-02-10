package disks

type hydratedGenesis struct {
	MiningValue                    uint8   `json:"mining_value" hydro:"0"`
	BlockBaseDifficulty            uint    `json:"block_base_difficulty" hydro:"1"`
	BlockIncreasePerHashDifficulty float64 `json:"block_increase_per_hash_difficulty" hydro:"2"`
	LinkDifficulty                 uint    `json:"link_difficulty" hydro:"3"`
}
