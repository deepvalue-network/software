package disks

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	mined_block "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/blockchain/domain/chains/peers"
	"github.com/steve-care-software/products/blockchain/domain/genesis"
	"github.com/steve-care-software/products/blockchain/domain/links"
	mined_link "github.com/steve-care-software/products/blockchain/domain/links/mined"
	"github.com/steve-care-software/products/libs/hash"
	"github.com/steve-care-software/products/libs/hashtree"
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

func newChain(
	id *uuid.UUID,
	peers peers.Peers,
	root mined_block.Block,
	gen genesis.Genesis,
	createdOn time.Time,
	peerSyncInterval time.Duration,
	head mined_link.Link,
	previousChain chains.Chain,
) (chains.Chain, error) {
	builder := chains.NewBuilder(peerSyncInterval).
		Create().
		WithID(id).
		WithPeers(peers).
		WithRoot(root).
		WithGenesis(gen).
		CreatedOn(createdOn)

	if head != nil {
		builder.WithHead(head)
	}

	if previousChain != nil {
		builder.WithOriginal(previousChain)
	}

	return builder.Now()
}
