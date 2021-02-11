package servers

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/blockchain/domain/links"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hydro"
	"github.com/gorilla/mux"
)

const internalErrorOutput = "server error"

const invalidHashErrorOutput = "the given hash, in the URL, is invalid"

const invalidIDErrorOutput = "the given id, in the URL, is invalid"

const missingParamErrorOutput = "the '%s' parameter was expected, none given"

const hashKeyname = "hash"

const idKeyname = "id"

const retrievePattern = "%s/%s"

// internal elements provided by outside:
var internalPeerSyncInterval time.Duration
var internalChainRepository chains.Repository
var internalTimeLayout string

// internal elements provided created by init:
var internalHydroAdapter hydro.Adapter

// Init initializes the package
func Init(
	peerSyncInterval time.Duration,
	chainRepository chains.Repository,
	timeLayout string,
) {

	internalPeerSyncInterval = peerSyncInterval
	internalChainRepository = chainRepository
	internalTimeLayout = timeLayout
}

// NewServer creates a new server instance
func NewServer(
	rep repositories.Application,
	router *mux.Router,
	waitPeriod time.Duration,
	port uint,
) Server {
	hashAdapter := hash.NewAdapter()
	return createServer(rep, hashAdapter, router, waitPeriod, port)
}

// Server represents a rest api server
type Server interface {
	Start()
	Stop()
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
