package votes

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/propositions"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash          hash.Hash
	prop          propositions.Proposition
	createdOn     time.Time
	isApproved    bool
	isCancel      bool
	isDisapproved bool
}

func createContentWithApproved(
	hash hash.Hash,
	prop propositions.Proposition,
	createdOn time.Time,
) Content {
	return createContentInternally(hash, prop, createdOn, true, false, false)
}

func createContentWithCancel(
	hash hash.Hash,
	prop propositions.Proposition,
	createdOn time.Time,
) Content {
	return createContentInternally(hash, prop, createdOn, false, true, false)
}

func createContentWithDisapproved(
	hash hash.Hash,
	prop propositions.Proposition,
	createdOn time.Time,
) Content {
	return createContentInternally(hash, prop, createdOn, false, false, true)
}

func createContentInternally(
	hash hash.Hash,
	prop propositions.Proposition,
	createdOn time.Time,
	isApproved bool,
	isCancel bool,
	isDisapproved bool,
) Content {
	out := content{
		hash:          hash,
		prop:          prop,
		createdOn:     createdOn,
		isApproved:    isApproved,
		isCancel:      isCancel,
		isDisapproved: isDisapproved,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Proposition returns the proposition
func (obj *content) Proposition() propositions.Proposition {
	return obj.prop
}

// IsApproved returns true if the proposition is approved, false otherwise
func (obj *content) IsApproved() bool {
	return obj.isApproved
}

// IsCancel returns true if the proposition is canceled, false otherwise
func (obj *content) IsCancel() bool {
	return obj.isCancel
}

// IsDisapproved returns true if the proposition is disapproved, false otherwise
func (obj *content) IsDisapproved() bool {
	return obj.isDisapproved
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
