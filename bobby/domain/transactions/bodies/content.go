package bodies

import (
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/access"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/contents"
	"github.com/steve-care-software/products/libs/hash"
)

type content struct {
	container containers.Transaction
	content   contents.Transaction
	access    access.Transaction
}

func createContentWithContainer(
	container containers.Transaction,
) Content {
	return createContentInternally(container, nil, nil)
}

func createContentWithContent(
	content contents.Transaction,
) Content {
	return createContentInternally(nil, content, nil)
}

func createContentWithAccess(
	access access.Transaction,
) Content {
	return createContentInternally(nil, nil, access)
}

func createContentInternally(
	container containers.Transaction,
	cnt contents.Transaction,
	access access.Transaction,
) Content {
	out := content{
		container: container,
		content:   cnt,
		access:    access,
	}

	return &out
}

func (obj *content) Hash() hash.Hash {
	if obj.IsContainer() {
		return obj.Container().Hash()
	}

	if obj.IsContent() {
		return obj.Content().Hash()
	}

	return obj.Access().Hash()
}

// IsContainer returns true if there is a container, false otherwise
func (obj *content) IsContainer() bool {
	return obj.container != nil
}

// Container returns the container, if any
func (obj *content) Container() containers.Transaction {
	return obj.container
}

// IsContent returns true if there is a content, false otherwise
func (obj *content) IsContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *content) Content() contents.Transaction {
	return obj.content
}

// IsAccess returns true if there is an access, false otherwise
func (obj *content) IsAccess() bool {
	return obj.access != nil
}

// Access returns the access, if any
func (obj *content) Access() access.Transaction {
	return obj.access
}
