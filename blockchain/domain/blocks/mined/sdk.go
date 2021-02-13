package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewValidator creates a new validator instance
func NewValidator() Validator {
	hashAdapter := hash.NewAdapter()
	return createValidator(hashAdapter)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// CalculateDifficulty calculates the difficulty
func CalculateDifficulty(baseDifficulty uint, incrPerHash float64, amountHashes int) uint {
	return calculateDifficulty(baseDifficulty, incrPerHash, amountHashes)
}

// NewPointer creates a new mined block pointer
func NewPointer() Block {
	return new(block)
}

// Validator represents a mined block validator
type Validator interface {
	Execute(gen genesis.Genesis, block Block) error
}

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithBlock(block blocks.Block) Builder
	WithResults(results string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a mined block
type Block interface {
	Hash() hash.Hash
	Block() blocks.Block
	Results() string
	CreatedOn() time.Time
}

// Repository represents a block repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(minedBlockHash hash.Hash) (Block, error)
	RetrieveByBlockHash(blockHash hash.Hash) (Block, error)
}

// Service represents a block service
type Service interface {
	Insert(block Block) error
	Delete(block Block) error
	DeleteByBlock(block blocks.Block) error
}
