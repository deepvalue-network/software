package servers

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	mined_block "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	mined_link "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hashtree"
	uuid "github.com/satori/go.uuid"
)

func newBlock(
	ht hashtree.Compact,
) (blocks.Block, error) {
	hashes := []hash.Hash{}
	leaves := ht.Leaves().Leaves()
	for _, oneLeaf := range leaves {
		hashes = append(hashes, oneLeaf.Head())
	}

	return blocks.NewBuilder().
		Create().
		WithHashes(hashes).
		Now()
}

func newBlockMined(
	block blocks.Block,
	results string,
	createdOn time.Time,
) (mined_block.Block, error) {
	return mined_block.NewBuilder().
		Create().
		WithBlock(block).
		WithResults(results).
		CreatedOn(createdOn).
		Now()
}

func newLink(
	index uint,
	prevMinedLink hash.Hash,
	nextBlock blocks.Block,
) (links.Link, error) {
	return links.NewBuilder().
		Create().
		WithIndex(index).
		WithPreviousMinedLink(prevMinedLink).
		WithNextBlock(nextBlock).
		Now()
}

func newLinkMined(
	link links.Link,
	results string,
	createdOn time.Time,
) (mined_link.Link, error) {
	return mined_link.NewBuilder().
		Create().
		WithLink(link).
		WithResults(results).
		CreatedOn(createdOn).
		Now()
}

func newPeer(
	server string,
	createdOn time.Time,
	lastUpdatedOn *time.Time,
) (peers.Peer, error) {
	builder := peers.NewPeerBuilder().Create().CreatedOn(createdOn).WithServer(server)
	if lastUpdatedOn != nil {
		builder.LastUpdatedOn(*lastUpdatedOn)
	}

	return builder.Now()
}

func newPeers(
	id *uuid.UUID,
	syncInterval time.Duration,
	list []peers.Peer,
	lastSyncTime *time.Time,
) (peers.Peers, error) {
	builder := peers.NewBuilder().Create().WithID(id).WithSyncDuration(syncInterval).WithList(list)
	if lastSyncTime != nil {
		builder.LastSyncTime(*lastSyncTime)
	}

	return builder.Now()
}

func newGenesis(
	miningValue uint8,
	blockBaseDifficulty uint,
	blockIncreasePerHashDifficulty float64,
	linkDifficulty uint,
) (genesis.Genesis, error) {
	return genesis.NewBuilder().Create().
		WithMiningValue(miningValue).
		WithBlockBaseDifficulty(blockBaseDifficulty).
		WithBlockIncreasePerHashDifficulty(blockIncreasePerHashDifficulty).
		WithLinkDifficulty(linkDifficulty).
		Now()
}

func newChain(
	id *uuid.UUID,
	peers peers.Peers,
	root mined_block.Block,
	gen genesis.Genesis,
	createdOn time.Time,
	head mined_link.Link,
) (chains.Chain, error) {
	builder := chains.NewBuilder(internalPeerSyncInterval).
		Create().
		WithID(id).
		WithPeers(peers).
		WithRoot(root).
		WithGenesis(gen).
		CreatedOn(createdOn)

	if head != nil {
		builder.WithHead(head)
	}

	if head != nil {
		// retrieve previous version of chain:
		prevChain, err := internalChainRepository.Retrieve(id)
		if err != nil {
			return nil, err
		}

		builder.WithOriginal(prevChain)
	}

	return builder.Now()
}
