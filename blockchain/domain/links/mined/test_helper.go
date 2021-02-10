package mined

import "github.com/deepvalue-network/software/blockchain/domain/links"

// CreateLinkForTests creates link instance for tests
func CreateLinkForTests() Link {
	link := links.CreateLinkForTests()
	results := "345234"
	ins, err := NewBuilder().Create().WithLink(link).WithResults(results).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
