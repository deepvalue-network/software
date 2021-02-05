package mined

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/links"
	"github.com/deepvalue-network/software/libs/hash"
)

type link struct {
	hash      hash.Hash  `hydro:"Hash, Hash"`
	link      links.Link `hydro:"Link, Link"`
	results   string     `hydro:"Results, Results"`
	createdOn time.Time  `hydro:"CreatedOn, CreatedOn"`
}

func createLink(
	hash hash.Hash,
	lnk links.Link,
	results string,
	createdOn time.Time,
) Link {
	out := link{
		hash:      hash,
		link:      lnk,
		results:   results,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Link returns the link
func (obj *link) Link() links.Link {
	return obj.link
}

// Results returns the results
func (obj *link) Results() string {
	return obj.results
}

// CreatedOn returns the creation time
func (obj *link) CreatedOn() time.Time {
	return obj.createdOn
}
