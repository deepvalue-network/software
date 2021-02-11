package disks

import (
	"os"
	"testing"
	"time"

	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	files_disks "github.com/deepvalue-network/software/libs/files/infrastructure/disks"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_linkMined_Success(t *testing.T) {
	basePath := "./test_files"
	defer func() {
		os.RemoveAll(basePath)
	}()

	// init:
	Init(basePath, 0777, time.Duration(time.Second))

	// creates the link service:
	serviceFileBlock := files_disks.NewService(internalHydroAdapter, basePath, 0777)
	serviceBlock := NewServiceBlock(serviceFileBlock)
	serviceLink := NewServiceLink(serviceBlock, serviceFileBlock)

	// build a link:
	link := link_mined.CreateLinkForTests()

	// save the link:
	err := serviceLink.Insert(link.Link())
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, link, t)
}
