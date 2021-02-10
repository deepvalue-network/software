package chains

import (
	"time"

	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
)

// CreateChainForTests creates a new chain instance for tests
func CreateChainForTests() Chain {
	peerSyncInterval := time.Duration(time.Second * 30)
	peers := peers.CreatePeersForTests()
	gen := genesis.CreateGenesisForTests()
	root := block_mined.CreateBlockForTests()
	head := link_mined.CreateLinkForTests()
	ins, err := NewBuilder(peerSyncInterval).Create().WithPeers(peers).WithGenesis(gen).WithRoot(root).WithHead(head).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
