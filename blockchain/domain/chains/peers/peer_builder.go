package peers

import (
	"errors"
	"net/url"
	"strconv"
	"time"
)

type peerBuilder struct {
	original Peer
	normal   string
	tor      string
}

func createPeerBuilder() PeerBuilder {
	out := peerBuilder{
		original: nil,
		normal:   "",
		tor:      "",
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

// WithNormalServer adds a normal server to the builder
func (app *peerBuilder) WithNormalServer(normal string) PeerBuilder {
	app.normal = normal
	return app
}

// WithTorServer adds a tor server to the builder
func (app *peerBuilder) WithTorServer(tor string) PeerBuilder {
	app.tor = tor
	return app
}

// Now builds a new Peer instance
func (app *peerBuilder) Now() (Peer, error) {
	var content Content
	if app.normal != "" {
		server, err := app.extract(app.normal)
		if err != nil {
			return nil, err
		}

		content = createContentWithNormal(server)
	}

	if app.tor != "" {
		server, err := app.extract(app.tor)
		if err != nil {
			return nil, err
		}

		content = createContentWithTor(server)
	}

	if content == nil {
		return nil, errors.New("the normal or tor server is mandatory in order to build a Peer instance")
	}

	createdOn := time.Now().UTC()
	lastUpdatedOn := time.Now().UTC()
	if app.original != nil {
		createdOn = app.original.CreatedOn()
	}

	return createPeer(content, createdOn, lastUpdatedOn), nil
}

func (app *peerBuilder) extract(str string) (Server, error) {
	ins, err := url.Parse(str)
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(ins.Port())
	if err != nil {
		return nil, err
	}

	host := ins.Hostname()
	return createServer(host, uint(port)), nil
}
