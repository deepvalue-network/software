package access

import (
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type transaction struct {
	hash      hash.Hash
	resources selectors.Selector
	content   Content
}

func createTransaction(
	hash hash.Hash,
	resources selectors.Selector,
	content Content,
) Transaction {
	out := transaction{
		hash:      hash,
		resources: resources,
		content:   content,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Resources returns the resources
func (obj *transaction) Resources() selectors.Selector {
	return obj.resources
}

// Content returns the content
func (obj *transaction) Content() Content {
	return obj.content
}
