package requests

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	requestor hash.Hash
	pubKey    public.Key
	server    servers.Server
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	requestor hash.Hash,
	pubKey public.Key,
	server servers.Server,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		requestor: requestor,
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

// Requestor returns the requestor
func (obj *content) Requestor() hash.Hash {
	return obj.requestor
}

// PublicKey returns the public key
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
