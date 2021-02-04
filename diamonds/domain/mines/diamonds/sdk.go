package diamonds

import (
	"github.com/steve-care-software/products/diamonds/domain/genesis"
	"github.com/steve-care-software/products/libs/hashtree"
)

// Builder represents a diamond mine
type Builder interface {
	Create() Builder
	WithTree(tree hashtree.HashTree) Builder
	WithDiamonds(diamonds Diamonds) Builder
	Now() (Diamonds, error)
}

// Diamonds represents initial diamonds
type Diamonds interface {
	Tree() hashtree.HashTree
	All() []genesis.Genesis
}
