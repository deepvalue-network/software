package selectors

import (
	"github.com/steve-care-software/products/bobby/domain/selectors/specifiers"
	"github.com/steve-care-software/products/libs/hash"
)

type graphbase struct {
	hash    hash.Hash
	content GraphbaseContent
	parent  specifiers.Specifier
}

func createGraphbase(
	hash hash.Hash,
	content GraphbaseContent,
) Graphbase {
	return createGraphbaseInternally(hash, content, nil)
}

func createGraphbaseWithParent(
	hash hash.Hash,
	content GraphbaseContent,
	parent specifiers.Specifier,
) Graphbase {
	return createGraphbaseInternally(hash, content, parent)
}

func createGraphbaseInternally(
	hash hash.Hash,
	content GraphbaseContent,
	parent specifiers.Specifier,
) Graphbase {
	out := graphbase{
		hash:    hash,
		content: content,
		parent:  parent,
	}

	return &out
}

// Hash returns the hash
func (obj *graphbase) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *graphbase) Content() GraphbaseContent {
	return obj.content
}

// HasParent returns true if there is a parent, false otherwise
func (obj *graphbase) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *graphbase) Parent() specifiers.Specifier {
	return obj.parent
}
