package services

import (
	"github.com/deepvalue-network/software/bobby/application/transactions/services/states"
)

// Application represents the service application
type Application interface {
	Begin() (states.State, error)
	Commit(state states.State)
	Rollback() error
	Push() error
}
