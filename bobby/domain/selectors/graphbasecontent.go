package selectors

import "github.com/steve-care-software/products/bobby/domain/selectors/specifiers"

type graphbaseContent struct {
	specifier specifiers.Specifier
	metaData  Table
}

func createGraphbaseContentWithSpecifier(
	specifier specifiers.Specifier,
) GraphbaseContent {
	return createGraphbaseContentInternally(specifier, nil)
}

func createGraphbaseContentWithMetaData(
	metaData Table,
) GraphbaseContent {
	return createGraphbaseContentInternally(nil, metaData)
}

func createGraphbaseContentInternally(
	specifier specifiers.Specifier,
	metaData Table,
) GraphbaseContent {
	out := graphbaseContent{
		specifier: specifier,
		metaData:  metaData,
	}

	return &out
}

// IsSpecifier returns true if there is a specifier, false otherwise
func (obj *graphbaseContent) IsSpecifier() bool {
	return obj.specifier != nil
}

// Specifier returns the specifier, if any
func (obj *graphbaseContent) Specifier() specifiers.Specifier {
	return obj.specifier
}

// IsMetaData returns true if there is a metaData, false otherwise
func (obj *graphbaseContent) IsMetaData() bool {
	return obj.metaData != nil
}

// MetaData returns the metaData, if any
func (obj *graphbaseContent) MetaData() Table {
	return obj.metaData
}
