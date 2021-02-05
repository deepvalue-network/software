package access

import access "github.com/deepvalue-network/software/bobby/domain/resources"

type content struct {
	isRemove bool
	add      access.Access
}

func createContentWithRemove() Content {
	return createContentInternally(true, nil)
}

func createContentWithAdd(add access.Access) Content {
	return createContentInternally(false, add)
}

func createContentInternally(
	isRemove bool,
	add access.Access,
) Content {
	out := content{
		isRemove: isRemove,
		add:      add,
	}

	return &out
}

// IsRemove returns true if there is remove, false otherwise
func (obj *content) IsRemove() bool {
	return obj.isRemove
}

// IsAdd returns true if there is an access addition, false otherwise
func (obj *content) IsAdd() bool {
	return obj.add != nil
}

// Add returns the access to be added, if any
func (obj *content) Add() access.Access {
	return obj.add
}
