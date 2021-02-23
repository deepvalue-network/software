package connections

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/connections/requests"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	pubKeyAdapter := public.NewAdapter()
	return createContentBuider(hashAdapter, pubKeyAdapter)
}

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
	WithRequest(request requests.Request) ContentBuilder
	WithRequestee(requestee hash.Hash) ContentBuilder
	WithPublicKey(pubKey public.Key) ContentBuilder
	WithServer(server servers.Server) ContentBuilder
	CreatedOn(createdOn time.Time) ContentBuilder
	Now() (Content, error)
}

// Content represents a content request
type Content interface {
	Hash() hash.Hash
	Request() requests.Request
	Requestee() hash.Hash
	PublicKey() public.Key
	Server() servers.Server
	CreatedOn() time.Time
}
