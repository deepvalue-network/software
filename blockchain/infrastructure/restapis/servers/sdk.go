package servers

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/steve-care-software/products/blockchain/application/repositories"
	"github.com/steve-care-software/products/libs/hash"
)

const internalErrorOutput = "server error"

const invalidHashErrorOutput = "the given hash, in the URL, is invalid"

const invalidIDErrorOutput = "the given id, in the URL, is invalid"

const missingParamErrorOutput = "the '%s' parameter was expected, none given"

const hashKeyname = "hash"

const idKeyname = "id"

const retrievePattern = "%s/%s"

// NewServer creates a new server instance
func NewServer(
	rep repositories.Application,
	router *mux.Router,
	waitPeriod time.Duration,
	port uint,
) Server {
	hashAdapter := hash.NewAdapter()
	return createServer(rep, hashAdapter, router, waitPeriod, port)
}

// Server represents a rest api server
type Server interface {
	Start()
	Stop()
}
