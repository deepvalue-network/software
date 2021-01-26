package specifiers

import "github.com/steve-care-software/products/libs/hash"

type specifier struct {
	identifier  Identifier
	identifiers Identifiers
}

func createSpecifierWithIdentifier(
	identifier Identifier,
) Specifier {
	return createSpecifierInternally(identifier, nil)
}

func createSpecifierWithIdentifiers(
	identifiers Identifiers,
) Specifier {
	return createSpecifierInternally(nil, identifiers)
}

func createSpecifierInternally(
	identifier Identifier,
	identifiers Identifiers,
) Specifier {
	out := specifier{
		identifier:  identifier,
		identifiers: identifiers,
	}

	return &out
}

// Hash returns the hash
func (obj *specifier) Hash() hash.Hash {
	if obj.IsIdentifier() {
		return obj.Identifier().Hash()
	}

	return obj.Identifiers().Hash()
}

// IsIdentifier returns true if there is an identifier, false otherwise
func (obj *specifier) IsIdentifier() bool {
	return obj.identifier != nil
}

// Identifier returns the identifier, if any
func (obj *specifier) Identifier() Identifier {
	return obj.identifier
}

// IsIdentifiers returns true if there is identifiers, false otherwise
func (obj *specifier) IsIdentifiers() bool {
	return obj.identifiers != nil
}

// Identifiers returns the identifiers, if any
func (obj *specifier) Identifiers() Identifiers {
	return obj.identifiers
}
