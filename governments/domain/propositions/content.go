package propositions

import (
	"time"

	"github.com/deepvalue-network/software/governments/domain/governments"
	"github.com/deepvalue-network/software/libs/hash"
)

type content struct {
	hash      hash.Hash
	gov       governments.Government
	section   Section
	activeOn  time.Time
	deadline  time.Time
	createdOn time.Time
}

func createContent(
	hash hash.Hash,
	gov governments.Government,
	section Section,
	activeOn time.Time,
	deadline time.Time,
	createdOn time.Time,
) Content {
	out := content{
		hash:      hash,
		gov:       gov,
		section:   section,
		activeOn:  activeOn,
		deadline:  deadline,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Government returns the government
func (obj *content) Government() governments.Government {
	return obj.gov
}

// Section returns the section
func (obj *content) Section() Section {
	return obj.section
}

// ActiveOn returns the activation time
func (obj *content) ActiveOn() time.Time {
	return obj.activeOn
}

// Deadline returns the deadline time
func (obj *content) Deadline() time.Time {
	return obj.deadline
}

// CreatedOn returns the creation time
func (obj *content) CreatedOn() time.Time {
	return obj.createdOn
}
