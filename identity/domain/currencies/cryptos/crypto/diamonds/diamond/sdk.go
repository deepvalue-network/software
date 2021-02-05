package diamond

import (
	bill_owners "github.com/deepvalue-network/software/diamonds/domain/bills/owners"
	genesis_owners "github.com/deepvalue-network/software/diamonds/domain/genesis/spends/owners"
	"github.com/deepvalue-network/software/libs/hash"
)

// Diamond represents diamond
type Diamond interface {
	Content() Content
}

// Content represents a diamond content
type Content interface {
	Hash() hash.Hash
	IsGenesis() bool
	Genesis() genesis_owners.Genesis
	IsBill() bool
	Bill() bill_owners.Bill
}
