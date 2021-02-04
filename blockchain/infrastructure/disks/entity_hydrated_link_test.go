package disks

import (
	"os"
	"testing"

	"github.com/steve-care-software/products/blockchain/domain/links"
	files_disks "github.com/steve-care-software/products/libs/files/infrastructure/disks"
	"github.com/steve-care-software/products/libs/hydro"
)

func TestHydrate_link_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath)

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
