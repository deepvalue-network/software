package application

import (
	"github.com/steve-care-software/products/bobby/application/data"
	"github.com/steve-care-software/products/bobby/application/transactions"
)

// Application represents the bobby application
type Application interface {
	Transaction() transactions.Application
	Structure() data.Application
}
