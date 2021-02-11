package servers

import (
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_link_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a link:
	link := links.CreateLinkForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, link, t)
}
