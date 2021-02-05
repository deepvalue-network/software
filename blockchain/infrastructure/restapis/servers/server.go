package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/deepvalue-network/software/blockchain/application/repositories"
	"github.com/deepvalue-network/software/libs/hash"
)

type server struct {
	server      *http.Server
	rep         repositories.Application
	hashAdapter hash.Adapter
	router      *mux.Router
	waitPeriod  time.Duration
	port        uint
}

func createServer(
	rep repositories.Application,
	hashAdapter hash.Adapter,
	router *mux.Router,
	waitPeriod time.Duration,
	port uint,
) Server {
	out := server{
		server:      nil,
		rep:         rep,
		hashAdapter: hashAdapter,
		router:      router,
		waitPeriod:  waitPeriod,
		port:        port,
	}

	hashPattern := fmt.Sprintf("{%s:[0-9a-f]+}", hashKeyname)
	idPattern := fmt.Sprintf("{%s:[0-9a-f-]+}", idKeyname)

	blockURI := fmt.Sprintf("/blocks")
	blockRetrieveURI := fmt.Sprintf(retrievePattern, blockURI, hashPattern)

	minedBlockURI := fmt.Sprintf("/mblocks")
	minedBlockRetrieveURI := fmt.Sprintf(retrievePattern, minedBlockURI, hashPattern)

	linkURI := fmt.Sprintf("/links")
	linkRetrieveURI := fmt.Sprintf(retrievePattern, linkURI, hashPattern)

	minedLinkURI := fmt.Sprintf("/mlinks")
	minedLinkHeadURI := fmt.Sprintf("%s/head", minedLinkURI)
	minedLinkRetrieveURI := fmt.Sprintf(retrievePattern, minedLinkURI, hashPattern)

	chainURI := fmt.Sprintf("/chains")
	chainRetrieveURI := fmt.Sprintf(retrievePattern, chainURI, idPattern)

	out.router.HandleFunc(blockURI, out.blockList).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(blockRetrieveURI, out.blockRetrieve).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(minedBlockURI, out.minedBlockList).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(minedBlockRetrieveURI, out.minedBlockRetrieve).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(linkURI, out.linkList).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(linkRetrieveURI, out.linkRetrieve).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(minedLinkURI, out.minedLinkList).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(minedLinkHeadURI, out.minedLinkHead).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(minedLinkRetrieveURI, out.minedLinkRetrieve).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(chainURI, out.chainList).Methods(http.MethodGet, http.MethodOptions)
	out.router.HandleFunc(chainRetrieveURI, out.chainRetrieve).Methods(http.MethodGet, http.MethodOptions)

	return &out
}

// Start starts the server
func (app *server) Start() {
	if app.server != nil {
		return
	}

	app.server = &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", app.port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//block until we receive signal
	<-c

	// stops:
	app.Stop()
}

// Stop stops the server
func (app *server) Stop() {
	if app.server == nil {
		return
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), app.waitPeriod)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	app.server.Shutdown(ctx)

	// set the server as nil:
	app.server = nil

	// shut down:
	log.Println("shutting down")
	os.Exit(0)
}

func (app *server) blockList(w http.ResponseWriter, r *http.Request) {
	hashes, err := app.rep.Block().List()
	renderInsToJSON(w, hashes, err)
}

func (app *server) blockRetrieve(w http.ResponseWriter, r *http.Request) {
	hash := fetchHashFromParams(app.hashAdapter, w, r, hashKeyname)
	if hash == nil {
		return
	}

	block, err := app.rep.Block().Retrieve(*hash)
	renderInsToJSON(w, block, err)
}

func (app *server) minedBlockList(w http.ResponseWriter, r *http.Request) {
	hashes, err := app.rep.MinedBlock().List()
	renderInsToJSON(w, hashes, err)
}

func (app *server) minedBlockRetrieve(w http.ResponseWriter, r *http.Request) {
	hash := fetchHashFromParams(app.hashAdapter, w, r, hashKeyname)
	if hash == nil {
		return
	}

	block, err := app.rep.MinedBlock().Retrieve(*hash)
	renderInsToJSON(w, block, err)
}

func (app *server) linkList(w http.ResponseWriter, r *http.Request) {
	hashes, err := app.rep.Link().List()
	renderInsToJSON(w, hashes, err)
}

func (app *server) linkRetrieve(w http.ResponseWriter, r *http.Request) {
	hash := fetchHashFromParams(app.hashAdapter, w, r, hashKeyname)
	if hash == nil {
		return
	}

	block, err := app.rep.Link().Retrieve(*hash)
	renderInsToJSON(w, block, err)
}

func (app *server) minedLinkList(w http.ResponseWriter, r *http.Request) {
	hashes, err := app.rep.MinedLink().List()
	renderInsToJSON(w, hashes, err)
}

func (app *server) minedLinkHead(w http.ResponseWriter, r *http.Request) {
	head, err := app.rep.MinedLink().Head()
	renderInsToJSON(w, head, err)
}

func (app *server) minedLinkRetrieve(w http.ResponseWriter, r *http.Request) {
	hash := fetchHashFromParams(app.hashAdapter, w, r, hashKeyname)
	if hash == nil {
		return
	}

	block, err := app.rep.MinedLink().Retrieve(*hash)
	renderInsToJSON(w, block, err)
}

func (app *server) chainList(w http.ResponseWriter, r *http.Request) {
	hashes, err := app.rep.Chain().List()
	renderInsToJSON(w, hashes, err)
}

func (app *server) chainRetrieve(w http.ResponseWriter, r *http.Request) {
	id := fetchIDFromParams(w, r, idKeyname)
	if id == nil {
		return
	}

	chain, err := app.rep.Chain().Retrieve(id)
	renderInsToJSON(w, chain, err)
}
