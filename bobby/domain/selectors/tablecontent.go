package selectors

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/libs/hash"
)

type tableContent struct {
	specifier specifiers.Specifier
}

func createTableContentWithSpecifier(
	specifier specifiers.Specifier,
) TableContent {
	return createTableContentInternally(specifier)
}

func createTableContentInternally(
	specifier specifiers.Specifier,
) TableContent {
	out := tableContent{
		specifier: specifier,
	}

	return &out
}

// Hash returns the hash
func (obj *tableContent) Hash() hash.Hash {
	return obj.Specifier().Hash()
}

// IsSpecifier returns true if there is a specifier, false otherwise
func (obj *tableContent) IsSpecifier() bool {
	return obj.specifier != nil
}

// Specifier returns the specifier, if any
func (obj *tableContent) Specifier() specifiers.Specifier {
	return obj.specifier
}
