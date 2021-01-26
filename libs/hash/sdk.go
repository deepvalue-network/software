package hash

const amountLettersPerFolder = 16

// NewAdapter returns a new hash adapter
func NewAdapter() Adapter {
	return createAdapter()
}

// Hash represents a hash
type Hash [64]byte

// Adapter represents an hash adapter
type Adapter interface {
	FromBytes(input []byte) (*Hash, error)
	FromMultiBytes(input [][]byte) (*Hash, error)
	FromString(input string) (*Hash, error)
}
