package application

import (
	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/blockchain/application/services"
)

// NewApplication creates a new application instance
func NewApplication(
	rep repositories.Application,
	serv services.Application,
) Application {
	return createApplication(rep, serv)
}

// Application represents a blockchain application
type Application interface {
	Repositories() repositories.Application
	Services() services.Application
}
