package services

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/blockchain/application/repositories"
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	mined_block "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/blockchain/domain/genesis"
	"github.com/steve-care-software/products/blockchain/domain/links"
	mined_link "github.com/steve-care-software/products/blockchain/domain/links/mined"
	"github.com/steve-care-software/products/libs/hash"
)

// maxMiningValue represents the max mining value before adding another miner number to the slice
const maxMiningValue = 2147483647

// maxMiningTries represents the max mining characters to try before abandonning
const maxMiningTries = 2147483647

// maxDifficulty represents the max difficulty a block can have
const maxDifficulty = 127

// NewApplication creates a new application instance
func NewApplication(
	block Block,
	minedBlock MinedBlock,
	link Link,
	minedLink MinedLink,
	chain Chain,
) Application {
	return createApplication(
		block,
		minedBlock,
		link,
		minedLink,
		chain,
	)
}

// NewChain creates a new chain application instance
func NewChain(
	peerSyncInterval time.Duration,
	remoteAppBuilder repositories.RemoteBuilder,
	chainService chains.Service,
	blockApp Block,
	minedBlockApp MinedBlock,
	minedLinkApp MinedLink,
	minedLinkRepository repositories.MinedLink,
	chainRepositoryApp repositories.Chain,
) Chain {
	chainBuilder := chains.NewBuilder(peerSyncInterval)
	genesisBuilder := genesis.NewBuilder()
	return createChain(
		chainService,
		chainBuilder,
		genesisBuilder,
		blockApp,
		minedBlockApp,
		minedLinkApp,
		minedLinkRepository,
		remoteAppBuilder,
		chainRepositoryApp,
	)
}

// NewMinedLink creates a new mined link application instance
func NewMinedLink(
	minedLinkService mined_link.Service,
	linkRepositoryApp repositories.Link,
	minedLinkRepositoryApp repositories.MinedLink,
	minerApp Miner,
) MinedLink {
	minedLinkBuilder := mined_link.NewBuilder()
	return createMinedLink(minedLinkService, minedLinkBuilder, linkRepositoryApp, minedLinkRepositoryApp, minerApp)
}

// NewLink creates a new link application instance
func NewLink(
	linkService links.Service,
	blockRepository repositories.Block,
	linkRepository repositories.Link,
	minedLinkRepository repositories.MinedLink,
) Link {
	linkBuilder := links.NewBuilder()
	return createLink(
		linkService,
		linkBuilder,
		blockRepository,
		linkRepository,
		minedLinkRepository,
	)
}

// NewMinedBlock creates a new mined block application instance
func NewMinedBlock(
	mineBlockService mined_block.Service,
	mineBlockRepositoryApp repositories.MinedBlock,
	blockRepositoryApp repositories.Block,
	minerApp Miner,
) MinedBlock {
	mineBlockBuilder := mined_block.NewBuilder()
	return createMinedBlock(
		mineBlockBuilder,
		mineBlockService,
		mineBlockRepositoryApp,
		blockRepositoryApp,
		minerApp,
	)
}

// NewBlock creates a new block application instance
func NewBlock(
	blockRepository blocks.Repository,
	blockService blocks.Service,
) Block {
	blockBuilder := blocks.NewBuilder()
	return createBlock(
		blockBuilder,
		blockRepository,
		blockService,
	)
}

// NewMiner creates a new miner application instance
func NewMiner() Miner {
	hashAdapter := hash.NewAdapter()
	return createMiner(hashAdapter)
}

// Application represents the blockchain application
type Application interface {
	Block() Block
	MinedBlock() MinedBlock
	Link() Link
	MinedLink() MinedLink
	Chain() Chain
}

// Miner represents a miner application
type Miner interface {
	Test(difficulty uint) (string, *time.Duration, error)
	Mine(miningValue uint8, difficulty uint, hash hash.Hash) (string, *time.Duration, error)
}

// Block represents a block application
type Block interface {
	Create(hashes []hash.Hash) (blocks.Block, error)
	Delete(hash hash.Hash) error
}

// MinedBlock represents the mined block application
type MinedBlock interface {
	Mine(miningValue uint8, baseDifficulty uint, incrPerHash float64, blockHash hash.Hash) (mined_block.Block, error)
	MineList(miningValue uint8, baseDifficulty uint, incrPerHash float64) ([]mined_block.Block, error)
	Delete(hash hash.Hash) error
}

// Link represents the link application
type Link interface {
	Create(prevMinedLinkHash hash.Hash, nextBlockHash hash.Hash) (links.Link, error)
	Delete(hash hash.Hash) error
}

// MinedLink represents the mined link application
type MinedLink interface {
	Mine(miningValue uint8, difficulty uint, linkHash hash.Hash) (mined_link.Link, error)
	MineList(miningValue uint8, difficulty uint) ([]mined_link.Link, error)
	Delete(hash hash.Hash) error
}

// Chain represents a chain application
type Chain interface {
	Update(id *uuid.UUID) error
	Delete(id *uuid.UUID) error
	Sync(waitPeriod time.Duration)
	Create(
		id *uuid.UUID,
		miningValue uint8,
		blockBaseDifficulty uint,
		blockIncreaseDiffPerHash float64,
		linkDifficulty uint,
		initialHashes []hash.Hash,
	) (chains.Chain, error)
}
