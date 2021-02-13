package disks

import (
	"os"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_chain_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a chain:
	chain := chains.CreateChainForTests()

	// save the root block:
	root := chain.Root()
	err := internalServiceBlockMined.Insert(root)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// save the head, if any:
	if chain.HasHead() {
		head := chain.Head()
		err := internalServiceLinkMined.Insert(head)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}
	}

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, chain, t)
}
