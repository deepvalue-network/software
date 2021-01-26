package application

import (
	"github.com/steve-care-software/products/blockchain/application/repositories"
	"github.com/steve-care-software/products/blockchain/application/services"
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
