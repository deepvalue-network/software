package servers

import (
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_peers_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a peers:
	peers := peers.CreatePeersForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, peers, t)
}
