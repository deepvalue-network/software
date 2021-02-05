package disks

import (
	"os"
	"testing"

	blocks_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_block_mined_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath)

	// creates the block service:
	serviceFile := files_disks.NewService(internalHydroAdapter, basePath, 0777)
	serviceBlock := NewServiceBlock(serviceFile)

	// build a mined block:
	minedBlock := blocks_mined.CreateBlockForTests()

	// save the block:
	err := serviceBlock.Insert(minedBlock.Block())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, minedBlock, t)
}
