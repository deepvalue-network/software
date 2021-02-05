package access

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
)

type access struct {
	id  *uuid.UUID            `hydro:"ID, ID"`
	sig signature.PrivateKey  `hydro:"Signature, SigPK"`
	enc encryption.PrivateKey `hydro:"Encryption, EncPK"`
}

func createAccess(
	id *uuid.UUID,
	sig signature.PrivateKey,
	enc encryption.PrivateKey,
) Access {
	out := access{
		id:  id,
		sig: sig,
		enc: enc,
	}

	return &out
}

// ID returns the id
func (obj *access) ID() *uuid.UUID {
	return obj.id
}

// Signature returns the signature PK
func (obj *access) Signature() signature.PrivateKey {
	return obj.sig
}

// Encryption returns the encryption PK
func (obj *access) Encryption() encryption.PrivateKey {
	return obj.enc
}
