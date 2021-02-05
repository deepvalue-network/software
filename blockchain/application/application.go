package application

import (
	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/blockchain/application/services"
)

type application struct {
	rep  repositories.Application
	serv services.Application
}

func createApplication(
	rep repositories.Application,
	serv services.Application,
) Application {
	out := application{
		rep:  rep,
		serv: serv,
	}

	return &out
}

// Repositories returns the repositories application
func (app *application) Repositories() repositories.Application {
	return app.rep
}

// Services returns the services application
func (app *application) Services() services.Application {
	return app.serv
}
