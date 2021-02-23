package connections

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// Builder represents a connection builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(sig signature.Signature) Builder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Hash() hash.Hash
	Content() Content
	Signature() signature.Signature
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithRequestee(requestee hash.Hash) ContentBuilder
	WithPublicKey(pubKey public.Key) ContentBuilder
	WithServer(server servers.Server) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a content request
type Content interface {
	Hash() hash.Hash
	Requestee() hash.Hash
	PublicKey() public.Key
	Server() servers.Server
	CreatedOn() time.Time
}
