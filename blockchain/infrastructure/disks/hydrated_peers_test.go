package disks

import (
	"os"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_peers_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, time.Duration(time.Second))

	// build a peers:
	peers := peers.CreatePeersForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, peers, t)
}
