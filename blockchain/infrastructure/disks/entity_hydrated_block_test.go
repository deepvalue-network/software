package disks

import (
	"testing"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hydro"
)

func TestHydrate_block_Success(t *testing.T) {
	// build a block:
	block := blocks.CreateBlockForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, block, t)
}
