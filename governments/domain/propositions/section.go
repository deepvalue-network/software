package propositions

import (
	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/governments/domain/governments/shareholders"
	"github.com/deepvalue-network/software/libs/hash"
)

type section struct {
	gov     governments.Content
	holders shareholders.ShareHolders
	custom  *hash.Hash
}

func createSectionWithGovernment(
	gov governments.Content,
) Section {
	return createSectionInternally(gov, nil, nil)
}

func createSectionWithShareHolders(
	holders shareholders.ShareHolders,
) Section {
	return createSectionInternally(nil, holders, nil)
}

func createSectionWithCustom(
	custom *hash.Hash,
) Section {
	return createSectionInternally(nil, nil, custom)
}

func createSectionInternally(
	gov governments.Content,
	holders shareholders.ShareHolders,
	custom *hash.Hash,
) Section {
	out := section{
		gov:     gov,
		holders: holders,
		custom:  custom,
	}

	return &out
}

// Hash returns the hash
func (obj *section) Hash() hash.Hash {
	if obj.IsGovernment() {

	}

	if obj.IsShareHolders() {

	}

	return *obj.custom
}

// IsGovernment returns true if there is a government, false otherwise
func (obj *section) IsGovernment() bool {
	return obj.gov != nil
}

// Government returns the government, if any
func (obj *section) Government() governments.Content {
	return obj.gov
}

// IsShareHolders returns true if there is shareHolders, false otherwise
func (obj *section) IsShareHolders() bool {
	return obj.holders != nil
}

// ShareHolders returns the shareHolders, if any
func (obj *section) ShareHolders() shareholders.ShareHolders {
	return obj.holders
}

// IsCustom returns true if there is custom, false otherwise
func (obj *section) IsCustom() bool {
	return obj.custom != nil
}

// Custom returns custom, if any
func (obj *section) Custom() *hash.Hash {
	return obj.custom
}
