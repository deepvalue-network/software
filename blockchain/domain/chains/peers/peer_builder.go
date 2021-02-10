package peers

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

type peerBuilder struct {
	original      Peer
	server        string
	createdOn     *time.Time
	lastUpdatedOn *time.Time
}

func createPeerBuilder() PeerBuilder {
	out := peerBuilder{
		original:      nil,
		server:        "",
		createdOn:     nil,
		lastUpdatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *peerBuilder) Create() PeerBuilder {
	return createPeerBuilder()
}

// WithOriginal adds an original to the builder
func (app *peerBuilder) WithOriginal(original Peer) PeerBuilder {
	app.original = original
	return app
}

// WithServer adds a sever to the builder
func (app *peerBuilder) WithServer(server string) PeerBuilder {
	app.server = server
	return app
}

// CreatedOn adds a creation time to the builder
func (app *peerBuilder) CreatedOn(createdOn time.Time) PeerBuilder {
	app.createdOn = &createdOn
	return app
}

// LastUpdatedOn adds a lastUpdatedOn time to the builder
func (app *peerBuilder) LastUpdatedOn(lastUpdatedOn time.Time) PeerBuilder {
	app.lastUpdatedOn = &lastUpdatedOn
	return app
}

// Now builds a new Peer instance
func (app *peerBuilder) Now() (Peer, error) {
	var content Content
	if app.server != "" {
		server, scheme, err := app.extract(app.server)
		if err != nil {
			return nil, err
		}

		if scheme == NormalProtocol {
			content = createContentWithNormal(server)
		}

		if scheme == TorProtocol {
			content = createContentWithTor(server)
		}
	}

	if content == nil {
		return nil, errors.New("the normal or tor server is mandatory in order to build a Peer instance")
	}

	if app.original != nil {
		original := app.original.CreatedOn()
		app.createdOn = &original

		lastUpdatedOn := time.Now().UTC()
		app.lastUpdatedOn = &lastUpdatedOn
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.lastUpdatedOn == nil {
		lastUpdatedOn := time.Now().UTC()
		app.lastUpdatedOn = &lastUpdatedOn
	}

	return createPeer(content, *app.createdOn, *app.lastUpdatedOn), nil
}

func (app *peerBuilder) extract(str string) (Server, string, error) {
	ins, err := url.Parse(str)
	if err != nil {
		return nil, "", err
	}

	port, err := strconv.Atoi(ins.Port())
	if err != nil {
		return nil, "", err
	}

	host := ins.Hostname()
	return createServer(host, uint(port)), ins.Scheme, nil
}
