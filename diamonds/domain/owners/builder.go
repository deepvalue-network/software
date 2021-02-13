package owners

import (
	"errors"
	"strconv"
	"time"

	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	encPKAdapter encryption.Adapter
	hashAdapter  hash.Adapter
	seed         string
	sigPK        signature.PrivateKey
	encPK        encryption.PrivateKey
	createdOn    *time.Time
}

func createBuilder(
	encPKAdapter encryption.Adapter,
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		encPKAdapter: encPKAdapter,
		hashAdapter:  hashAdapter,
		seed:         "",
		sigPK:        nil,
		encPK:        nil,
		createdOn:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.encPKAdapter, app.hashAdapter)
}

// WithSeed adds a seed to the builder
func (app *builder) WithSeed(seed string) Builder {
	app.seed = seed
	return app
}

// WithSignaturePK adds a signature PK to the builder
func (app *builder) WithSignaturePK(sigPK signature.PrivateKey) Builder {
	app.sigPK = sigPK
	return app
}

// WithEncryptionPK adds an encryption PK to the builder
func (app *builder) WithEncryptionPK(encPK encryption.PrivateKey) Builder {
	app.encPK = encPK
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Owner instance
func (app *builder) Now() (Owner, error) {
	if app.seed == "" {
		return nil, errors.New("the seed is mandatory in order to build an Owner instance")
	}

	if app.sigPK == nil {
		return nil, errors.New("the signature PK is mandatory in order to build an Owner instance")
	}

	if app.encPK == nil {
		return nil, errors.New("the encryption PK is mandatory in order to build an Owner instance")
	}

	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.seed),
		[]byte(app.sigPK.String()),
		app.encPKAdapter.ToBytes(app.encPK),
		[]byte(strconv.Itoa(int(app.createdOn.Second()))),
	})

	if err != nil {
		return nil, err
	}

	return createOwner(*hsh, app.seed, app.sigPK, app.encPK, *app.createdOn), nil
}
