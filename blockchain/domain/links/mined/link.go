package mined

import (
	"time"

	"github.com/steve-care-software/products/blockchain/domain/links"
	"github.com/steve-care-software/products/libs/hash"
)

type link struct {
	hash      hash.Hash
	link      links.Link
	results   string
	createdOn time.Time
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
