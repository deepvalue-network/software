package requests

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	sig         signature.RingSignature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		sig:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithContent adds a content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(sig signature.RingSignature) Builder {
	app.sig = sig
	return app
}

// Now builds a new Request instance
func (app *builder) Now() (Request, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Request instance")
	}

	if app.sig == nil {
		return nil, errors.New("the ring signature is mandatory in order to build a Request instance")
	}

	msg := app.content.Hash().String()
	if !app.sig.Verify(msg) {
		str := fmt.Sprintf("the ring signature (%s) could not validate against the content (hash: %s) while building a Request instance", app.sig.String(), msg)
		return nil, errors.New(str)
	}

	if !app.content.Stake().Origin().Same(app.content.To()) {
		str := fmt.Sprintf("the stake origin (shareHolder's hash: %s) does not match the to pubKey hashes", app.content.Stake().Origin().Hash().String())
		return nil, errors.New(str)
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.sig.String()),
	})

	if err != nil {
		return nil, err
	}

	return createRequest(*hash, app.content, app.sig), nil
}
