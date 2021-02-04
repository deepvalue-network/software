package disks

import (
	"time"
)

type entityHydratedLinkMined struct {
	Link      string    `json:"link_hash" hydro:"0"`
	Results   string    `json:"results" hydro:"1"`
	CreatedOn time.Time `json:"created_on" hydro:"2"`
}

type entityHydratedChain struct {
	ID                string           `json:"hash" hydro:"0"`
	Peers             *hydratedPeers   `json:"peers" hydro:"1"`
	Root              string           `json:"root_block_mined_hash" hydro:"2"`
	Genesis           *hydratedGenesis `json:"genesis" hydro:"3"`
	CreatedOn         time.Time        `json:"created_on" hydro:"4"`
	PeerSyncInterval  time.Duration    `json:"peer_sync_interval" hydro:"5"`
	Head              string           `json:"head_mined_link_hash" hydro:"6"`
	PreviousChainHash string           `json:"previous_chain_hash" hydro:"7"`
}
