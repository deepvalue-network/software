package connections

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/requests"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	request   requests.Request
	requestee hash.Hash
	pubKey    public.Key
	server    servers.Server
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	request requests.Request,
	requestee hash.Hash,
	pubKey public.Key,
	server servers.Server,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		request:   request,
		requestee: requestee,
		pubKey:    pubKey,
		server:    server,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Request returns the request
func (obj *content) Request() requests.Request {
	return obj.request
}

// Requestee returns the requestee
func (obj *content) Requestee() hash.Hash {
	return obj.requestee
}

// PublicKey returns the encryption publicKey
func (obj *content) PublicKey() public.Key {
	return obj.pubKey
}

// Server returns the server
func (obj *content) Server() servers.Server {
	return obj.server
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
