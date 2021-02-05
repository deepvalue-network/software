package selectors

import (
	"github.com/deepvalue-network/software/bobby/domain/selectors/specifiers"
	"github.com/deepvalue-network/software/libs/hash"
)

type database struct {
	hash      hash.Hash
	graphbase specifiers.Specifier
	content   DatabaseContent
}

func createDatabase(
	hash hash.Hash,
	graphbase specifiers.Specifier,
	content DatabaseContent,
) Database {
	out := database{
		hash:      hash,
		graphbase: graphbase,
		content:   content,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// Graphbase returns the graphbase
func (obj *database) Graphbase() specifiers.Specifier {
	return obj.graphbase
}

// Content returns the content
func (obj *database) Content() DatabaseContent {
	return obj.content
}
