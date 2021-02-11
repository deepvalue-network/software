package disks

import (
	"os"
	"reflect"
	"testing"
	"time"

	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
)

func TestHydrate_linkMined_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a link:
	link := link_mined.CreateLinkForTests()

	// save the link:
	err := internalServiceLinkMined.Insert(link)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the link:
	retLink, err := internalRepositoryLinkMined.Retrieve(link.Hash())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hydrated, err := internalHydroAdapter.Hydrate(link)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retHydrated, err := internalHydroAdapter.Hydrate(retLink)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// compare:
	if !reflect.DeepEqual(hydrated, retHydrated) {
		t.Errorf("the compared instances are different")
		return
	}
}
