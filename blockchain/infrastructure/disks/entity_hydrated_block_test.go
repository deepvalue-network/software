package disks

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
)

func TestHydrate_block_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a mined block:
	block := blocks.CreateBlockForTests()

	// save the block:
	err := internalServiceBlock.Insert(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the block:
	retBlock, err := internalRepositoryBlock.Retrieve(block.Tree().Head())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hydrated, err := internalHydroAdapter.Hydrate(block)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retHydrated, err := internalHydroAdapter.Hydrate(retBlock)
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
