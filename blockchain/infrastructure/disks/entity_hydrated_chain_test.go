package disks

import (
	"os"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_chain_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// create the services:
	fileService := files_disks.NewService(internalHydroAdapter, basePath, 0777)
	blockService := NewServiceBlock(fileService)
	minedBlockService := NewServiceBlockMined(blockService, fileService)
	linkService := NewServiceLink(blockService, fileService)
	minedLinkService := NewServiceLinkMined(linkService, fileService)

	// build a chain:
	chain := chains.CreateChainForTests()

	// save the root block:
	root := chain.Root()
	err := minedBlockService.Insert(root)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the head, if any:
	if chain.HasHead() {
		head := chain.Head()
		err := minedLinkService.Insert(head)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}
	}

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, chain, t)
}
