package accesses

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/products/identity/domain/accesses/access"
)

type accesses struct {
	accessFactory access.Factory
	mp            map[string]access.Access
}

func createAccesses(
	accessFactory access.Factory,
	mp map[string]access.Access,
) Accesses {
	out := accesses{
		accessFactory: accessFactory,
		mp:            mp,
	}

	return &out
}

// All returns all accesses
func (obj *accesses) All() map[string]access.Access {
	return obj.mp
}

// Create creates a new access
func (obj *accesses) Create(name string) error {
	if _, ok := obj.mp[name]; ok {
		str := fmt.Sprintf("there is already an access for name: %s", name)
		return errors.New(str)
	}

	ins, err := obj.accessFactory.Create()
	if err != nil {
		return err
	}

	obj.mp[name] = ins
	return nil
}

// Fetch fetches an access by name
func (obj *accesses) Fetch(name string) (access.Access, error) {
	if _, ok := obj.mp[name]; !ok {
		str := fmt.Sprintf("the access (name: %s) could not be fetched because it doesn't exists", name)
		return nil, errors.New(str)
	}

	return obj.mp[name], nil
}

// Delete deletes an access by name
func (obj *accesses) Delete(name string) error {
	if _, ok := obj.mp[name]; !ok {
		str := fmt.Sprintf("the access (name: %s) could not be deleted because it doesn't exists", name)
		return errors.New(str)
	}

	delete(obj.mp, name)
	return nil
}
