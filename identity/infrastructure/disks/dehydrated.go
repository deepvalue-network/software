package disks

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/identity/domain/accesses"
	"github.com/deepvalue-network/software/identity/domain/accesses/access"
	"github.com/deepvalue-network/software/identity/domain/users"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
)

func newUser(name string, seed string, accesses accesses.Accesses) (users.User, error) {
	return users.NewBuilder(EncPKBitrate).Create().WithName(name).WithSeed(seed).WithAccesses(accesses).Now()
}

func newAccesses(mp map[string]access.Access) (accesses.Accesses, error) {
	return accesses.NewBuilder(EncPKBitrate).Create().WithMap(mp).Now()
}

func newAccess(id uuid.UUID, sigPK signature.PrivateKey, encPK encryption.PrivateKey) (access.Access, error) {
	return access.NewBuilder().Create().WithID(&id).WithSignature(sigPK).WithEncryption(encPK).Now()
}
