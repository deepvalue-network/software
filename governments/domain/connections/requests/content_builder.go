package requests

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/hash"
)

type contentBuilder struct {
	hashAdapter   hash.Adapter
	pubKeyAdapter public.Adapter
	requestor     *hash.Hash
	pubKey        public.Key
	server        servers.Server
	createdOn     *time.Time
}

func createContentBuilder(
	hashAdapter hash.Adapter,
	pubKeyAdapter public.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter:   hashAdapter,
		pubKeyAdapter: pubKeyAdapter,
		requestor:     nil,
		pubKey:        nil,
		server:        nil,
		createdOn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(app.hashAdapter, app.pubKeyAdapter)
}

// WithRequestor adds a requestor to the builder
func (app *contentBuilder) WithRequestor(requestor hash.Hash) ContentBuilder {
	app.requestor = &requestor
	return app
}

// WithPublicKey adds a public key to the builder
func (app *contentBuilder) WithPublicKey(pubKey public.Key) ContentBuilder {
	app.pubKey = pubKey
	return app
}

// WithServer adds a server to the builder
func (app *contentBuilder) WithServer(server servers.Server) ContentBuilder {
	app.server = server
	return app
}

// CreatedOn adds a creation time to the builder
func (app *contentBuilder) CreatedOn(createdOn time.Time) ContentBuilder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.requestor == nil {
		return nil, errors.New("the requestor hash is mandatory in order to build a request Content instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the encryption publicKey is mandatory in order to build a request Content instance")
	}

	if app.server == nil {
		return nil, errors.New("the server is mandatory in order to build a request Content instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.requestor.Bytes(),
		app.pubKeyAdapter.ToBytes(app.pubKey),
		[]byte(app.server.String()),
		[]byte(strconv.Itoa(app.createdOn.Second())),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*hash, *app.requestor, app.pubKey, app.server, *app.createdOn), nil
}
