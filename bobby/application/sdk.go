package application

import (
	"github.com/deepvalue-network/software/bobby/application/data"
	"github.com/deepvalue-network/software/bobby/application/transactions"
)

// Application represents the bobby application
type Application interface {
	Transaction() transactions.Application
	Structure() data.Application
}
