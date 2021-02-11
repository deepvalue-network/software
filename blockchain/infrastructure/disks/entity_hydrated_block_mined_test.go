package disks

import (
	"os"
	"reflect"
	"testing"
	"time"

	blocks_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
)

func TestHydrate_block_mined_Success(t *testing.T) {
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
