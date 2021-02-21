package identities

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type shareHolderBuilder struct {
	hashAdapter hash.Adapter
	gov         governments.Government
	public      shareholders.ShareHolder
	sigPK       signature.PrivateKey
}

func createShareHolderBuilder(
	hashAdapter hash.Adapter,
) ShareHolderBuilder {
	out := shareHolderBuilder{
		hashAdapter: hashAdapter,
		gov:         nil,
		public:      nil,
		sigPK:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *shareHolderBuilder) Create() ShareHolderBuilder {
	return createShareHolderBuilder(app.hashAdapter)
}

// WithGovernment adds a government to the builder
func (app *shareHolderBuilder) WithGovernment(gov governments.Government) ShareHolderBuilder {
	app.gov = gov
	return app
}

// WithPublic adds a public shareHolder to the builder
func (app *shareHolderBuilder) WithPublic(public shareholders.ShareHolder) ShareHolderBuilder {
	app.public = public
	return app
}

// WithSigPK adds a signature's privateKey to the builder
func (app *shareHolderBuilder) WithSigPK(sigPK signature.PrivateKey) ShareHolderBuilder {
	app.sigPK = sigPK
	return app
}

// Now builds a new ShareHolder instance
func (app *shareHolderBuilder) Now() (ShareHolder, error) {
	if app.gov == nil {
		return nil, errors.New("the government is mandatory in order to build a ShareHolder instance")
	}

	if app.public == nil {
		return nil, errors.New("the public shareHolder is mandatory in order to build a ShareHolder instance")
	}

	if app.sigPK == nil {
		return nil, errors.New("the signature privateKey is mandatory in order to build a ShareHolder instance")
	}

	pubKey := app.sigPK.PublicKey()
	pubKeyHash, err := app.hashAdapter.FromBytes([]byte(pubKey.String()))
	if err != nil {
		return nil, err
	}

	if !app.public.Contains(*pubKeyHash) {
		str := fmt.Sprintf("the public shareHolder was expected to contain the signature PrivateKey's hashed PublicKey (%s)", pubKeyHash.String())
		return nil, errors.New(str)
	}

	if !app.gov.ShareHolders().Has(app.public) {
		str := fmt.Sprintf("the given government (ID: %s) does not contain the given public ShareHolder (hash: %s)", app.gov.ID().String(), app.public.Hash().String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.gov.Hash().Bytes(),
		app.public.Hash().Bytes(),
		[]byte(app.sigPK.String()),
	})

	if err != nil {
		return nil, err
	}

	return createShareHolder(*hash, app.gov, app.public, app.sigPK), nil
}
