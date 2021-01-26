package elements

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/bobby/domain/structures/tables/schemas"
	"github.com/steve-care-software/products/libs/hash"
)

type elements struct {
	hash             hash.Hash
	list             []Element
	mpByPropertyHash map[string]Element
}

func createElements(
	hash hash.Hash,
	list []Element,
	mpByPropertyHash map[string]Element,
) Elements {
	out := elements{
		hash:             hash,
		list:             list,
		mpByPropertyHash: mpByPropertyHash,
	}

	return &out
}

// Hash returns the hash
func (obj *elements) Hash() hash.Hash {
	return obj.hash
}

// All returns the elements
func (obj *elements) All() []Element {
	return obj.list
}

// IsEmpty returns true if the list is empty, false otherwise
func (obj *elements) IsEmpty() bool {
	return len(obj.list) <= 0
}

// Fits returns nil if the given properties fits the elements.  Otherwise, it returns the first error
func (obj *elements) Fits(properties schemas.Properties) error {
	list := properties.All()
	for _, oneProperty := range list {
		keyname := oneProperty.Resource().Hash().String()
		if _, ok := obj.mpByPropertyHash[keyname]; !ok {
			str := fmt.Sprintf("the property (name: %s) does not have an associated value", oneProperty.Name())
			return errors.New(str)
		}
	}

	return nil
}
