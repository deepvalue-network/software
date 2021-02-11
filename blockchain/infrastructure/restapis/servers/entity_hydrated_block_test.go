package servers

import (
	"testing"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_block_Success(t *testing.T) {
	// build a block:
	block := blocks.CreateBlockForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, block, t)
}
