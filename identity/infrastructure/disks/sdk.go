package disks

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/identity/domain/accesses"
	"github.com/deepvalue-network/software/identity/domain/accesses/access"
	"github.com/deepvalue-network/software/identity/domain/users"
	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hydro"
)

// EncPKBitrate represents the encryption PrivateKey bitrate
const EncPKBitrate = 4096

var hydroAdapter hydro.Adapter

func init() {

	accessBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*access.Access)(nil)).
		WithDehydratedConstructor(newAccess).
		WithDehydratedPointer(access.NewPointer()).
		WithHydratedPointer(new(hydratedAccess)).
		OnHydrate(accessOnHydrateEventFn).
		OnDehydrate(accessOnDehydrateEventFn).
		Now()

	if err != nil {
		panic(err)
	}

	accessesBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*accesses.Accesses)(nil)).
		WithDehydratedConstructor(newAccesses).
		WithDehydratedPointer(accesses.NewPointer()).
		WithHydratedPointer(new(hydratedAccesses)).
		Now()

	if err != nil {
		panic(err)
	}

	userBridge, err := hydro.NewBridgeBuilder().Create().
		WithDehydratedInterface((*users.User)(nil)).
		WithDehydratedConstructor(newUser).
		WithDehydratedPointer(users.NewPointer()).
		WithHydratedPointer(new(hydratedUser)).
		Now()

	if err != nil {
		panic(err)
	}

	// build the manager:
	manager := hydro.NewManagerFactory().Create()

	// register the bridges:
	manager.Register(accessBridge)
	manager.Register(accessesBridge)
	manager.Register(userBridge)

	// create the adapter:
	hydroAdapter, err = hydro.NewAdapterBuilder().Create().WithManager(manager).Now()
	if err != nil {
		panic(err)
	}

}

func accessOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if id, ok := ins.(*uuid.UUID); ok {
		return id.String(), nil
	}

	if sigPK, ok := ins.(signature.PrivateKey); ok {
		return sigPK.String(), nil
	}

	if encPK, ok := ins.(encryption.PrivateKey); ok {
		return encryption.NewAdapter().ToEncoded(encPK), nil
	}

	return nil, nil
}

func accessOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "ID" {
		return uuid.FromString(ins.(string))
	}

	if fieldName == "SigPK" {
		return signature.NewPrivateKeyAdapter().ToPrivateKey(ins.(string))
	}

	if fieldName == "EncPK" {
		return encryption.NewAdapter().FromEncoded(ins.(string))
	}

	return nil, nil
}
