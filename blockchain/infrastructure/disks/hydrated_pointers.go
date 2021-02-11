package disks

import (
	"fmt"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	uuid "github.com/satori/go.uuid"
)

func createGenesisForBridge() *HydratedGenesis {
	out := HydratedGenesis{
		MiningValue:                    uint8(0x0),
		BlockBaseDifficulty:            uint(2),
		BlockIncreasePerHashDifficulty: float64(0.43),
		LinkDifficulty:                 uint(11),
	}

	return &out
}

func createPeerForBridge() *HydratedPeer {
	now := time.Now().UTC()
	out := HydratedPeer{
		Content:       fmt.Sprintf("%s://127.0.0.1:80", peers.NormalProtocol),
		CreatedOn:     now.String(),
		LastUpdatedOn: now.String(),
	}

	return &out
}

func createPeersForBridge() *HydratedPeers {
	now := time.Now().UTC()
	id := uuid.NewV4()
	out := HydratedPeers{
		ID:           id.String(),
		SyncInterval: int64(time.Duration(time.Second * 30)),
		List: []*HydratedPeer{
			createPeerForBridge(),
		},
		LastSyncTime: now.String(),
	}

	return &out
}
