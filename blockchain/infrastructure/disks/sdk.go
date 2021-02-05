package disks

import (
	"path/filepath"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/files/domain/files"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

const timeLayout = "2006-01-02T15:04:05.000Z"

var internalHydroAdapter hydro.Adapter
var internalRepositoryBlock blocks.Repository
var internalMinedRepositoryBlock block_mined.Repository

// Init initializes the package
func Init(
	basePath string,
) {
	// create the block repository:
	blockPtr := new(entityHydratedBlock)
	blockBasePath := filepath.Join(basePath, "blocks")
	repositoryFileBlock := files_disks.NewRepository(internalHydroAdapter, blockBasePath, blockPtr)
	repositoryBlock := NewRepositoryBlock(repositoryFileBlock)

	// create the mined block repository:
	minedBlockPtr := new(entityHydratedBlock)
	minedBlockBasePath := filepath.Join(basePath, "blocks/mined")
	repositoryFileMinedBlock := files_disks.NewRepository(internalHydroAdapter, minedBlockBasePath, minedBlockPtr)
	repositoryBlockMined := NewRepositoryBlockMined(repositoryFileMinedBlock)

	// assign:
	internalRepositoryBlock = repositoryBlock
	internalMinedRepositoryBlock = repositoryBlockMined
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

	// build the manager:
	manager := hydro.NewManagerFactory().Create()

	// register the bridges:
	manager.Register(blockBridge)
	manager.Register(blockMinedBridge)
	manager.Register(linkBridge)

	// create the adapter:
	internalHydroAdapter, err = hydro.NewAdapterBuilder().Create().WithManager(manager).Now()
	if err != nil {
		panic(err)
	}
}
