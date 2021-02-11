package servers

import (
	"testing"
	"time"

	blocks_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_block_mined_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a mined block:
	minedBlock := blocks_mined.CreateBlockForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, minedBlock, t)
}
