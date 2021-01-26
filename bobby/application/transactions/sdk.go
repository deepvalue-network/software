package transactions

import (
	"github.com/steve-care-software/products/bobby/application/transactions/repositories"
	"github.com/steve-care-software/products/bobby/application/transactions/services"
)

// Application represents the transaction application
type Application interface {
	Repository() repositories.Application
	Service() services.Application
}
