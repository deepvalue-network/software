package disks

import (
	"os"

	"github.com/deepvalue-network/software/libs/files/domain/files"
	"github.com/deepvalue-network/software/libs/hydro"
)

const fileDoesNotExistsPattern = "the file (path: %s) does not exists"

const fileAlreadyExistsPattern = "the file (path: %s) already exists"

// NewRepository creates a new disk repository instance
func NewRepository(
	hydroAdapter hydro.Adapter,
	basePath string,
	ptr interface{},
) files.Repository {
	return createRepository(hydroAdapter, basePath, ptr)
}

// NewService creates a new service instance
func NewService(
	hydroAdapter hydro.Adapter,
	basePath string,
	fileMode os.FileMode,
) files.Service {
	return creatService(hydroAdapter, basePath, fileMode)
}
