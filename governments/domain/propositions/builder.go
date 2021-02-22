package propositions

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	sigs        []signature.RingSignature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		sigs:        nil,
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

// WithSignatures add signatures to the builder
func (app *builder) WithSignatures(sigs []signature.RingSignature) Builder {
	app.sigs = sigs
	return app
}

// Now builds a new Proposition instance
func (app *builder) Now() (Proposition, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Proposition instance")
	}

	if app.sigs == nil {
		app.sigs = []signature.RingSignature{}
	}

	if len(app.sigs) <= 0 {
		return nil, errors.New("there must be at least 1 Signature in order to build a Proposition instance")
	}

	contentHash := app.content.Hash()
	data := [][]byte{
		contentHash.Bytes(),
	}

	msg := contentHash.String()
	for _, oneSig := range app.sigs {
		if !oneSig.Verify(msg) {
			str := fmt.Sprintf("there is at least 1 ring signature (%s) that cannot be verified against the content (hash: %s)", msg, oneSig.String())
			return nil, errors.New(str)
		}

		data = append(data, []byte(oneSig.String()))
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createProposition(*hash, app.content, app.sigs), nil
}
