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
	"github.com/deepvalue-network/software/libs/events"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hydro"
)

const (
	// EventBlockInsert represents an insert block event
	EventBlockInsert = iota

	// EventBlockDelete represents a delete block event
	EventBlockDelete

	// EventBlockMinedInsert represents an insert mined block event
	EventBlockMinedInsert

	// EventBlockMinedDelete represents a delete mined block event
	EventBlockMinedDelete

	// EventLinkInsert represents an insert link event
	EventLinkInsert

	// EventLinkDelete represents a delete link event
	EventLinkDelete

	// EventLinkMinedInsert represents an insert mined link event
	EventLinkMinedInsert

	// EventLinkMinedDelete represents a delete mined link event
	EventLinkMinedDelete

	// EventChainInsert represents an insert chain event
	EventChainInsert

	// EventChainUpdate represents an update chain event
	EventChainUpdate

	// EventChainDelete represents a delete chain event
	EventChainDelete
)

const timeLayout = "2006-01-02T15:04:05.000Z"

// event manager:
var internalEventManager events.Manager

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
var internalServiceLink links.Service
var internalServiceLinkMined link_mined.Service

// Init initializes the package
func Init(
	basePath string,
	fileMode os.FileMode,
	peerSyncInterval time.Duration,
) {

	// init events:
	eventManager, err := initEventManager()
	if err != nil {
		panic(err)
	}

	// assign event manager:
	internalEventManager = eventManager

	// interval assign
	internalPeerSyncInterval = peerSyncInterval

	// create the block repository:
	blockPtr := new(EntityHydratedBlock)
	blockBasePath := filepath.Join(basePath, "blocks")
	repositoryFileBlock := files_disks.NewRepository(internalHydroAdapter, blockBasePath, blockPtr)
	repositoryBlock := NewRepositoryBlock(repositoryFileBlock)

	// create the mined block repository:
	minedBlockPtr := new(EntityHydratedBlockMined)
	minedBlockBasePath := filepath.Join(basePath, "blocks_mined")
	pointerMinedBlockBasePath := filepath.Join(basePath, "blocks_mined_pointers")
	repositoryFileMinedBlock := files_disks.NewRepository(internalHydroAdapter, minedBlockBasePath, minedBlockPtr)
	repositoryPointerFileMinedBlock := files_disks.NewRepository(internalHydroAdapter, pointerMinedBlockBasePath, nil)
	repositoryBlockMined := NewRepositoryBlockMined(repositoryFileMinedBlock, repositoryPointerFileMinedBlock)

	// create the link repository:
	linkPtr := new(EntityHydratedLink)
	linkBasePath := filepath.Join(basePath, "links")
	blockPointerLinkBasePath := filepath.Join(basePath, "links_blocks_pointers")
	minedLinkPointerLinkBasePath := filepath.Join(basePath, "links_minedlinks_pointers")
	repositoryFileLink := files_disks.NewRepository(internalHydroAdapter, linkBasePath, linkPtr)
	repositoryBlockPointerFileLink := files_disks.NewRepository(internalHydroAdapter, blockPointerLinkBasePath, nil)
	repositoryMinedLinkPointerFileLink := files_disks.NewRepository(internalHydroAdapter, minedLinkPointerLinkBasePath, nil)
	linkRepository := NewRepositoryLink(repositoryFileLink, repositoryBlockPointerFileLink, repositoryMinedLinkPointerFileLink)

	// create the link mined repository:
	minedLinkPtr := new(EntityHydratedLinkMined)
	minedLinkBasePath := filepath.Join(basePath, "links_mined")
	linkPointerMinedLinkBasePath := filepath.Join(basePath, "links_mined_links_pointers")
	headPointerMinedLinkBasePath := filepath.Join(basePath, "links_mined_head_pointer")
	headFileName := "head.hash"
	repositoryFileLinkMined := files_disks.NewRepository(internalHydroAdapter, minedLinkBasePath, minedLinkPtr)
	linkPointerFileRepository := files_disks.NewRepository(internalHydroAdapter, linkPointerMinedLinkBasePath, nil)
	headPointerFileRepository := files_disks.NewRepository(internalHydroAdapter, headPointerMinedLinkBasePath, nil)
	minedLinkRepository := NewRepositoryLinkMined(repositoryFileLinkMined, linkPointerFileRepository, headPointerFileRepository, headFileName)

	// create the chain repository:
	chainLinkPtr := new(EntityHydratedChain)
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
	blockService := NewServiceBlock(internalEventManager, blockFileService)

	// create the mined block service:
	minedBlockFileService := files_disks.NewService(internalHydroAdapter, minedBlockBasePath, fileMode)
	minedBlockPointerFileService := files_disks.NewService(internalHydroAdapter, pointerMinedBlockBasePath, fileMode)
	minedBlockService := NewServiceBlockMined(internalEventManager, repositoryBlockMined, blockService, minedBlockFileService, minedBlockPointerFileService)

	// create the link service:
	linkFileService := files_disks.NewService(internalHydroAdapter, linkBasePath, fileMode)
	linkBlockPointerFileService := files_disks.NewService(internalHydroAdapter, blockPointerLinkBasePath, fileMode)
	linkMinedLinkPointerFileService := files_disks.NewService(internalHydroAdapter, minedLinkPointerLinkBasePath, fileMode)
	linkService := NewServiceLink(internalEventManager, linkRepository, blockService, linkFileService, linkBlockPointerFileService, linkMinedLinkPointerFileService)

	// create the mined link service:
	minedLinkFileService := files_disks.NewService(internalHydroAdapter, minedLinkBasePath, fileMode)
	minedLinkLinkPointerFileService := files_disks.NewService(internalHydroAdapter, linkPointerMinedLinkBasePath, fileMode)
	headPointerFileService := files_disks.NewService(internalHydroAdapter, headPointerMinedLinkBasePath, fileMode)
	minedLinkService := NewServiceLinkMined(internalEventManager, minedLinkRepository, linkService, minedLinkFileService, minedLinkLinkPointerFileService, headPointerFileService, headFileName)

	// service assign:
	internalServiceBlock = blockService
	internalServiceBlockMined = minedBlockService
	internalServiceLink = linkService
	internalServiceLinkMined = minedLinkService
}

// NewRepositoryChain creates a new chain repository
func NewRepositoryChain(
	fileRepository files.Repository,
) chains.Repository {
	return createRepositoryChain(fileRepository)
}

// NewServiceLinkMined creates a new disk link mined service instance
func NewServiceLinkMined(
	eventManager events.Manager,
	minedLinkRepository link_mined.Repository,
	linkService links.Service,
	fileService files.Service,
	linkPointerFileService files.Service,
	headPointerFileService files.Service,
	headFileName string,
) link_mined.Service {
	return createServiceLinkMined(eventManager, minedLinkRepository, linkService, fileService, linkPointerFileService, headPointerFileService, headFileName)
}

// NewRepositoryLinkMined represents a new disk link mined repository instance
func NewRepositoryLinkMined(
	fileRepository files.Repository,
	linkPointerFileRepository files.Repository,
	headPointerFileRepository files.Repository,
	headFileName string,
) link_mined.Repository {
	hashAdapter := hash.NewAdapter()
	return createRepositoryLinkMined(hashAdapter, fileRepository, linkPointerFileRepository, headPointerFileRepository, headFileName)
}

// NewServiceLink creates a new disk link service instance
func NewServiceLink(
	eventManager events.Manager,
	linkRepository links.Repository,
	blockService blocks.Service,
	fileService files.Service,
	blockPointerFileService files.Service,
	minedLinkPointerFileService files.Service,
) links.Service {
	return createServiceLink(
		eventManager,
		linkRepository,
		blockService,
		fileService,
		blockPointerFileService,
		minedLinkPointerFileService,
	)
}

// NewRepositoryLink creates a new disk link repository instance
func NewRepositoryLink(
	fileRepository files.Repository,
	blockPointerFileRepository files.Repository,
	minedLinkPointerFileRepository files.Repository,
) links.Repository {
	hashAdapter := hash.NewAdapter()
	return createRepositoryLink(hashAdapter, fileRepository, blockPointerFileRepository, minedLinkPointerFileRepository)
}

// NewRepositoryBlock creates a new disk block repository instance
func NewRepositoryBlock(
	fileRepository files.Repository,
) blocks.Repository {
	return createRepositoryBlock(fileRepository)
}

// NewServiceBlock creates a new disk block service
func NewServiceBlock(
	eventManager events.Manager,
	fileService files.Service,
) blocks.Service {
	return createServiceBlock(eventManager, fileService)
}

// NewRepositoryBlockMined creates a new disk mined block repository
func NewRepositoryBlockMined(
	fileRepository files.Repository,
	pointerFileRepository files.Repository,
) block_mined.Repository {
	hashAdapter := hash.NewAdapter()
	return createRepositoryBlockMined(hashAdapter, fileRepository, pointerFileRepository)
}

// NewServiceBlockMined creates a new disk mined block service
func NewServiceBlockMined(
	eventManager events.Manager,
	minedBlockRepository block_mined.Repository,
	blockService blocks.Service,
	fileService files.Service,
	pointerFileService files.Service,
) block_mined.Service {
	return createServiceBlockMined(eventManager, minedBlockRepository, blockService, fileService, pointerFileService)
}

func init() {
	// create bridges:
	blockBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*blocks.Block)(nil)).
		WithDehydratedConstructor(newBlock).
		WithDehydratedPointer(blocks.NewPointer()).
		WithHydratedPointer(new(EntityHydratedBlock)).
		OnHydrate(blockOnHydrateEventFn).
		OnDehydrate(blockOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	blockMinedBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*block_mined.Block)(nil)).
		WithDehydratedConstructor(newBlockMined).
		WithDehydratedPointer(block_mined.NewPointer()).
		WithHydratedPointer(new(EntityHydratedBlockMined)).
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
		WithHydratedPointer(new(EntityHydratedLink)).
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
		WithHydratedPointer(new(EntityHydratedLinkMined)).
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
		OnHydrate(genesisOnHydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	chainBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*chains.Chain)(nil)).
		WithDehydratedConstructor(newChain).
		WithDehydratedPointer(chains.NewPointer()).
		WithHydratedPointer(new(EntityHydratedChain)).
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
