package hashtree

import (
	"github.com/steve-care-software/products/pangolin/libs/hash"
)

// NewAdapter creates a new hashtree adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// NewBuilder creates a new hashtree builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the hashtree adapter
type Adapter interface {
	FromJSON(js *JSONCompact) (Compact, error)
	ToJSON(ht Compact) *JSONCompact
}

// Builder represents an hashtree builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks [][]byte) Builder
	WithJS(js []byte) Builder
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
