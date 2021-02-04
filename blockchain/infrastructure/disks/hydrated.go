package disks

import (
	"time"
)

type hydratedPeer struct {
	Server        string    `json:"server" hydro:"0"`
	CreatedOn     time.Time `json:"created_on" hydro:"1"`
	LastUpdatedOn time.Time `json:"last_updated_on" hydro:"2"`
}

type hydratedPeers struct {
	ID           string          `json:"id" hydro:"0"`
	SyncInterval time.Duration   `json:"sync_interval" hydro:"1"`
	List         []*hydratedPeer `json:"list" hydro:"2"`
	LastSyncTime *time.Time      `json:"last_sync_time" hydro:"3"`
}

type hydratedGenesis struct {
	Hash        string                     `json:"hash" hydro:"0"`
	MiningValue uint8                      `json:"mining_value" hydro:"1"`
	Difficulty  *hydratedGenesisDifficulty `json:"difficulty" hydro:"2"`
}

type hydratedGenesisDifficulty struct {
	Block *hydratedGenesisDifficultyBlock `json:"block" hydro:"0"`
	Link  uint                            `json:"link" hydro:"1"`
}

type hydratedGenesisDifficultyBlock struct {
	Base            uint   `json:"base" hydro:"0"`
	IncreasePerHash string `json:"increase_per_hash" hydro:"1"`
}
