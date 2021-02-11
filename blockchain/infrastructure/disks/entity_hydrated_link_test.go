package disks

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
)

func TestHydrate_link_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// build a link:
	link := links.CreateLinkForTests()

	// save the link:
	err := internalServiceLink.Insert(link)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve the link:
	retLink, err := internalRepositoryLink.Retrieve(link.Hash())
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
