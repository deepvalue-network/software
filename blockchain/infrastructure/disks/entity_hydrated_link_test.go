package disks

import (
	"os"
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_link_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// creates the block service:
	serviceFileBlock := files_disks.NewService(internalHydroAdapter, basePath, 0777)
	serviceBlock := NewServiceBlock(serviceFileBlock)

	// build a link:
	link := links.CreateLinkForTests()

	// save the block:
	err := serviceBlock.Insert(link.NextBlock())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, link, t)
}
