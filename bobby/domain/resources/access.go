package resources

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
)

type access struct {
	resource  Mutable
	owners    []*uuid.UUID
	encrypted public.Key
}

func createAccess(
	resource Mutable,
	owners []*uuid.UUID,
) Access {
	return createAccessInternally(resource, owners, nil)
}

func createAccessWithEncryptedPubkey(
	resource Mutable,
	owners []*uuid.UUID,
	encrypted public.Key,
) Access {
	return createAccessInternally(resource, owners, encrypted)
}

func createAccessInternally(
	resource Mutable,
	owners []*uuid.UUID,
	encrypted public.Key,
) Access {
	out := access{
		resource:  resource,
		owners:    owners,
		encrypted: encrypted,
	}

	return &out
}

// Resource returns the resource
func (obj *access) Resource() Mutable {
	return obj.resource
}

// Owners returns the owners
func (obj *access) Owners() []*uuid.UUID {
	return obj.owners
}

// IsEncrypted returns true if encrypted, false otherwise
func (obj *access) IsEncrypted() bool {
	return obj.encrypted != nil
}

// Encrypted returns the encryption pubKey, if any
func (obj *access) Encrypted() public.Key {
	return obj.encrypted
}
