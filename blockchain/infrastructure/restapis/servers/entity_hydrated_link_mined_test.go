package servers

import (
	"testing"
	"time"

	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_linkMined_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a link:
	link := link_mined.CreateLinkForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, link, t)
}
