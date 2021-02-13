package disks

import (
	"os"
	"reflect"
	"testing"
	"time"

	blocks_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
)

func TestHydrate_blockMined_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a mined block:
	minedBlock := blocks_mined.CreateBlockForTests()

	// save the block:
	err := internalServiceBlockMined.Insert(minedBlock)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the block:
	retMinedBlock, err := internalRepositoryBlockMined.Retrieve(minedBlock.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hydrated, err := internalHydroAdapter.Hydrate(minedBlock)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retHydrated, err := internalHydroAdapter.Hydrate(retMinedBlock)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	if !reflect.DeepEqual(hydrated, retHydrated) {
		t.Errorf("the compared instances are different")
		return
	}
}

func TestHydrate_blockMined_insert_deleteBlock_expectBlockMinedDeleted_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a mined block:
	minedBlock := blocks_mined.CreateBlockForTests()

	// save the mined block:
	err := internalServiceBlockMined.Insert(minedBlock)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the mined block:
	retFirstMinedBlock, err := internalRepositoryBlockMined.Retrieve(minedBlock.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare the blocks:
	if !retFirstMinedBlock.Hash().Compare(minedBlock.Hash()) {
		t.Errorf("the returned block is invalid, \nexpected: %s, \nreturned: %s", minedBlock.Hash().String(), retFirstMinedBlock.Hash().String())
		return
	}

	// delete the underlying block:
	err = internalServiceBlock.Delete(minedBlock.Block())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the mined block again, expect an error:
	_, err = internalRepositoryBlockMined.Retrieve(minedBlock.Hash())
	if err == nil {
		t.Errorf("the retrieval of the mined block was expected to return an error, nil returned")
		return
	}
}
