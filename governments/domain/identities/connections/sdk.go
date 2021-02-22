package connections

import (
	"github.com/deepvalue-network/software/governments/domain/connections"
	"github.com/deepvalue-network/software/governments/domain/connections/servers"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	uuid "github.com/satori/go.uuid"
)

// Connection represents a connection
type Connection interface {
	ID() *uuid.UUID
	Connection() connections.Connection
	Server() servers.Server
	SigPK() signature.PrivateKey
	EncPK() encryption.PrivateKey
}
