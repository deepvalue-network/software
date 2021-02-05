package services

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/libs/hash"
)

type chain struct {
	chainService        chains.Service
	chainBuilder        chains.Builder
	genesisBuilder      genesis.Builder
	blockApp            Block
	minedBlockApp       MinedBlock
	minedLinkApp        MinedLink
	minedLinkRepository repositories.MinedLink
	remoteAppBuilder    repositories.RemoteBuilder
	chainRepository     repositories.Chain
}

func createChain(
	chainService chains.Service,
	chainBuilder chains.Builder,
	genesisBuilder genesis.Builder,
	blockApp Block,
	minedBlockApp MinedBlock,
	minedLinkApp MinedLink,
	minedLinkRepository repositories.MinedLink,
	remoteAppBuilder repositories.RemoteBuilder,
	chainRepository repositories.Chain,
) Chain {
	out := chain{
		chainService:        chainService,
		chainBuilder:        chainBuilder,
		genesisBuilder:      genesisBuilder,
		blockApp:            blockApp,
		minedBlockApp:       minedBlockApp,
		minedLinkApp:        minedLinkApp,
		minedLinkRepository: minedLinkRepository,
		remoteAppBuilder:    remoteAppBuilder,
		chainRepository:     chainRepository,
	}
	return &out
}

// Update updates a chain by id
func (app *chain) Update(id *uuid.UUID) error {
	// retrieve the chain:
	chain, err := app.chainRepository.Retrieve(id)
	if err != nil {
		return err
	}

	// mine the block:
	genesis := chain.Genesis()
	genesisMiningValue := genesis.MiningValue()
	genesisDiff := genesis.Difficulty()
	genesisBlockDiff := genesisDiff.Block()

	// mine the blocks:
	baseDifficulty := genesisBlockDiff.Base()
	incrPerHash := genesisBlockDiff.IncreasePerHash()
	_, err = app.minedBlockApp.MineList(genesisMiningValue, baseDifficulty, incrPerHash)
	if err != nil {
		return err
	}

	// mine the links:
	genesisLinkDiff := genesisDiff.Link()
	_, err = app.minedLinkApp.MineList(genesisMiningValue, genesisLinkDiff)
	if err != nil {
		return err
	}

	// retrieve the head mined link:
	head, err := app.minedLinkRepository.Head()
	if err != nil {
		return err
	}

	// build the updated chain:
	updated, err := app.chainBuilder.Create().WithOriginal(chain).WithHead(head).Now()
	if err != nil {
		return err
	}

	// update the chain:
	return app.chainService.Update(chain, updated)
}

// Delete deletes a chain by id
func (app *chain) Delete(id *uuid.UUID) error {
	chain, err := app.chainRepository.Retrieve(id)
	if err != nil {
		return err
	}

	return app.chainService.Delete(chain)
}

// Create creates a new chain
func (app *chain) Create(
	id *uuid.UUID,
	miningValue uint8,
	blockBaseDifficulty uint,
	blockIncreaseDiffPerHash float64,
	linkDifficulty uint,
	initialHashes []hash.Hash,
) (chains.Chain, error) {
	// build the genesis instance:
	gen, err := app.genesisBuilder.Create().
		WithMiningValue(miningValue).
		WithBlockBaseDifficulty(blockBaseDifficulty).
		WithBlockIncreasePerHashDifficulty(blockIncreaseDiffPerHash).
		WithLinkDifficulty(linkDifficulty).
		Now()

	if err != nil {
		return nil, err
	}

	// create the root block:
	root, err := app.blockApp.Create(initialHashes)
	if err != nil {
		return nil, err
	}

	// mine the root block:
	blockHash := root.Tree().Head()
	minedRoot, err := app.minedBlockApp.Mine(miningValue, blockBaseDifficulty, blockIncreaseDiffPerHash, blockHash)
	if err != nil {
		return nil, err
	}

	// build the chain:
	chain, err := app.chainBuilder.Create().WithGenesis(gen).WithRoot(minedRoot).Now()
	if err != nil {
		return nil, err
	}

	// save the save:
	err = app.chainService.Insert(chain)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

// Sync sync the chains
func (app *chain) Sync(waitPeriod time.Duration) {
	for {
		// sync the chains:
		err := app.syncOnce()
		if err != nil {
			// log the error:
		}

		// wait:
		time.Sleep(waitPeriod)
	}
}

// syncOnce sync chains once
func (app *chain) syncOnce() error {
	// retrieve the chain ids:
	chainIDs, err := app.chainRepository.List()
	if err != nil {
		return err
	}

	// for each chain, sync:
	for _, oneChainID := range chainIDs {
		err := app.syncByID(oneChainID)
		if err != nil {
			// log the error:
		}
	}

	return nil
}

// syncByID sync chain by ID
func (app *chain) syncByID(chainID *uuid.UUID) error {
	// retrieve the chain:
	chain, err := app.chainRepository.Retrieve(chainID)
	if err != nil {
		return err
	}

	// fetch the peers:
	peers := chain.Peers()

	// if the peers has been recently synced, skip them:
	now := time.Now().UTC()
	syncInterval := peers.SyncInterval()
	after := now.Add(syncInterval * -1)
	if peers.HasLastSync() {
		lastSync := peers.LastSync()
		if lastSync.After(after) {
			return nil
		}
	}

	// sync the peers:
	list := peers.All()
	for _, onePeer := range list {
		// update the chain by peer:
		err := app.syncChainByPeer(chain, onePeer)
		if err != nil {
			return err
		}
	}

	return nil
}

// syncChainByPeer sync a chain by peer
func (app *chain) syncChainByPeer(localChain chains.Chain, ins peers.Peer) error {
	remoteApp, err := app.remoteAppBuilder.Create().WithPeer(ins).Now()
	if err != nil {
		// remove the peer from the peers list:
		err := localChain.Peers().Delete(ins)
		if err != nil {
			return err
		}

		return err
	}

	// fetch the chainID
	localChainID := localChain.ID()

	// retrieve the chain of the remote peer:
	remoteChain, err := remoteApp.Chain().Retrieve(localChainID)
	if err != nil {
		return err
	}

	// update the chain if needed:
	return app.updateFromRemote(localChain, remoteChain)
}

func (app *chain) updateFromRemote(local chains.Chain, remote chains.Chain) error {
	return nil
}
