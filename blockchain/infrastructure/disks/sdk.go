package disks

import (
	"os"
	"path/filepath"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

const timeLayout = "2006-01-02T15:04:05.000Z"

// peer sync interval:
var internalPeerSyncInterval time.Duration

// adapters:
var internalHydroAdapter hydro.Adapter

// repositories:
var internalRepositoryBlock blocks.Repository
var internalRepositoryBlockMined block_mined.Repository
var internalRepositoryLink links.Repository
var internalRepositoryLinkMined link_mined.Repository
var internalRepositoryChain chains.Repository

// services:
var internalServiceBlock blocks.Service
var internalServiceBlockMined block_mined.Service

// Init initializes the package
func Init(
	basePath string,
	fileMode os.FileMode,
	peerSyncInterval time.Duration,
) {
	// interval assign
	internalPeerSyncInterval = peerSyncInterval

	// create the block repository:
	blockPtr := new(entityHydratedBlock)
	blockBasePath := filepath.Join(basePath, "blocks")
	repositoryFileBlock := files_disks.NewRepository(internalHydroAdapter, blockBasePath, blockPtr)
	repositoryBlock := NewRepositoryBlock(repositoryFileBlock)

	// create the mined block repository:
	minedBlockPtr := new(entityHydratedBlock)
	minedBlockBasePath := filepath.Join(basePath, "blocks_mined")
	repositoryFileMinedBlock := files_disks.NewRepository(internalHydroAdapter, minedBlockBasePath, minedBlockPtr)
	repositoryBlockMined := NewRepositoryBlockMined(repositoryFileMinedBlock)

	// create the link repository:
	linkPtr := new(entityHydratedLink)
	linkBasePath := filepath.Join(basePath, "links")
	repositoryFileLink := files_disks.NewRepository(internalHydroAdapter, linkBasePath, linkPtr)
	linkRepository := NewRepositoryLink(repositoryFileLink)

	// create the link mined repository:
	minedLinkPtr := new(entityHydratedLinkMined)
	minedLinkBasePath := filepath.Join(basePath, "links_mined")
	repositoryFileLinkMined := files_disks.NewRepository(internalHydroAdapter, minedLinkBasePath, minedLinkPtr)
	minedLinkRepository := NewRepositoryLinkMined(repositoryFileLinkMined)

	// create the chain repository:
	chainLinkPtr := new(entityHydratedChain)
	chainBasePath := filepath.Join(basePath, "chains")
	repositoryFileChain := files_disks.NewRepository(internalHydroAdapter, chainBasePath, chainLinkPtr)
	chainRepository := NewRepositoryChain(repositoryFileChain)

	// repository assign:
	internalRepositoryBlock = repositoryBlock
	internalRepositoryBlockMined = repositoryBlockMined
	internalRepositoryLink = linkRepository
	internalRepositoryLinkMined = minedLinkRepository
	internalRepositoryChain = chainRepository

	// create the block service:
	blockFileService := files_disks.NewService(internalHydroAdapter, blockBasePath, fileMode)
	blockService := NewServiceBlock(blockFileService)

	// create the mined block service:
	minedBlockFileService := files_disks.NewService(internalHydroAdapter, minedBlockBasePath, fileMode)
	minedBlockService := NewServiceBlockMined(blockService, minedBlockFileService)

	// service assign:
	internalServiceBlock = blockService
	internalServiceBlockMined = minedBlockService
}

// NewRepositoryChain creates a new chain repository
func NewRepositoryChain(
	fileRepository files.Repository,
) chains.Repository {
	return createRepositoryChain(fileRepository)
}

// NewServiceLinkMined creates a new disk link mined service instance
func NewServiceLinkMined(
	linkService links.Service,
	fileService files.Service,
) link_mined.Service {
	return createServiceLinkMined(linkService, fileService)
}

// NewRepositoryLinkMined represents a new disk link mined repository instance
func NewRepositoryLinkMined(
	fileRepository files.Repository,
) link_mined.Repository {
	return createRepositoryLinkMined(fileRepository)
}

// NewServiceLink creates a new disk link service instance
func NewServiceLink(
	blockService blocks.Service,
	fileService files.Service,
) links.Service {
	return createServiceLink(blockService, fileService)
}

// NewRepositoryLink creates a new disk link repository instance
func NewRepositoryLink(
	fileRepository files.Repository,
) links.Repository {
	return createRepositoryLink(fileRepository)
}

// NewRepositoryBlock creates a new disk block repository instance
func NewRepositoryBlock(
	fileRepository files.Repository,
) blocks.Repository {
	return createRepositoryBlock(fileRepository)
}

// NewServiceBlock creates a new disk block service
func NewServiceBlock(
	fileService files.Service,
) blocks.Service {
	return createServiceBlock(fileService)
}

// NewRepositoryBlockMined creates a new disk mined block repository
func NewRepositoryBlockMined(
	fileRepository files.Repository,
) block_mined.Repository {
	return createRepositoryBlockMined(fileRepository)
}

// NewServiceBlockMined creates a new disk mined block service
func NewServiceBlockMined(
	blockService blocks.Service,
	fileService files.Service,
) block_mined.Service {
	return createServiceBlockMined(blockService, fileService)
}

func init() {
	blockBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*blocks.Block)(nil)).
		WithDehydratedConstructor(newBlock).
		WithDehydratedPointer(blocks.NewPointer()).
		WithHydratedPointer(new(entityHydratedBlock)).
		OnHydrate(blockOnHydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	blockMinedBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*block_mined.Block)(nil)).
		WithDehydratedConstructor(newBlockMined).
		WithDehydratedPointer(block_mined.NewPointer()).
		WithHydratedPointer(new(entityHydratedBlockMined)).
		OnHydrate(blockMinedOnHydrateEventFn).
		OnDehydrate(blockMinedOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	linkBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*links.Link)(nil)).
		WithDehydratedConstructor(newLink).
		WithDehydratedPointer(links.NewPointer()).
		WithHydratedPointer(new(entityHydratedLink)).
		OnHydrate(linkOnHydrateEventFn).
		OnDehydrate(linkOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	minedLinkBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*link_mined.Link)(nil)).
		WithDehydratedConstructor(newLinkMined).
		WithDehydratedPointer(link_mined.NewPointer()).
		WithHydratedPointer(new(entityHydratedLinkMined)).
		OnHydrate(linkMinedOnHydrateEventFn).
		OnDehydrate(linkMinedOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	peerBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*peers.Peer)(nil)).
		WithDehydratedConstructor(newPeer).
		WithDehydratedPointer(peers.NewPeerPointer()).
		WithHydratedPointer(createPeerForBridge()).
		OnHydrate(peerOnHydrateEventFn).
		OnDehydrate(peerOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	peersBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*peers.Peers)(nil)).
		WithDehydratedConstructor(newPeers).
		WithDehydratedPointer(peers.NewPointer()).
		WithHydratedPointer(createPeersForBridge()).
		OnHydrate(peersOnHydrateEventFn).
		OnDehydrate(peersOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	genesisBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*genesis.Genesis)(nil)).
		WithDehydratedConstructor(newGenesis).
		WithDehydratedPointer(genesis.NewPointer()).
		WithHydratedPointer(createGenesisForBridge()).
		Now()

	if err != nil {
		panic(err)
	}

	chainBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*chains.Chain)(nil)).
		WithDehydratedConstructor(newChain).
		WithDehydratedPointer(chains.NewPointer()).
		WithHydratedPointer(new(entityHydratedChain)).
		OnHydrate(chainOnHydrateEventFn).
		OnDehydrate(chainOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	// build the manager:
	manager := hydro.NewManagerFactory().Create()

	// register the bridges:
	manager.Register(blockBridge)
	manager.Register(blockMinedBridge)
	manager.Register(linkBridge)
	manager.Register(minedLinkBridge)
	manager.Register(peerBridge)
	manager.Register(peersBridge)
	manager.Register(genesisBridge)
	manager.Register(chainBridge)

	// create the adapter:
	internalHydroAdapter, err = hydro.NewAdapterBuilder().Create().WithManager(manager).Now()
	if err != nil {
		panic(err)
	}
}
