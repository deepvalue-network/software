package hashtree

import (
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new hashtree builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an hashtree builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks [][]byte) Builder
	Now() (HashTree, error)
}

// HashTree represents an hashtree
type HashTree interface {
	Height() int
	Length() int
	Head() hash.Hash
	Parent() ParentLeaf
	Compact() Compact
	Order(data [][]byte) ([][]byte, error)
}

// Block represents a block of hashes
type Block interface {
	Leaves() Leaves
	HashTree() (HashTree, error)
}

// ParentLeaf represents an hashtree parent leaf
type ParentLeaf interface {
	Left() Leaf
	Right() Leaf
	BlockLeaves() Leaves
	HashTree() (HashTree, error)
}

// Leaf represents an hashtree leaf
type Leaf interface {
	Head() hash.Hash
	HasParent() bool
	Parent() ParentLeaf
	Leaves() Leaves
	Height() int
}

// Leaves represents a list of Leaf instances
type Leaves interface {
	Leaves() []Leaf
	Merge(lves Leaves) Leaves
	HashTree() (HashTree, error)
}

// Compact represents a compact hashtree
type Compact interface {
	Head() hash.Hash
	Leaves() Leaves
	Length() int
}
