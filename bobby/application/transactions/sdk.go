package transactions

import (
	"github.com/deepvalue-network/software/bobby/application/transactions/repositories"
	"github.com/deepvalue-network/software/bobby/application/transactions/services"
)

// Application represents the transaction application
type Application interface {
	Repository() repositories.Application
	Service() services.Application
}
