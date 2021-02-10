package disks

import (
	"os"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_peer_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, time.Duration(time.Second))

	// build a peer:
	peer := peers.CreatePeerForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, peer, t)
}
