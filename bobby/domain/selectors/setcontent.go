package selectors

import "github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"

type setContent struct {
	specifier specifiers.Specifier
}

func createSetContentWithSpecifier(
	specifier specifiers.Specifier,
) SetContent {
	return createSetContentInternally(specifier)
}

func createSetContentInternally(
	specifier specifiers.Specifier,
) SetContent {
	out := setContent{
		specifier: specifier,
	}

	return &out
}

// IsSpecifier returns true if there is a specifier, false otherwise
func (obj *setContent) IsSpecifier() bool {
	return obj.specifier != nil
}

// Specifier returns the specifier, if any
func (obj *setContent) Specifier() specifiers.Specifier {
	return obj.specifier
}
