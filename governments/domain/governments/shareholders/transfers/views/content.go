package views

import "github.com/deepvalue-network/software/libs/hash"

type content struct {
	hash     hash.Hash
	section  Section
	newOwner []hash.Hash
}

func createContent(
	hash hash.Hash,
	section Section,
	newOwner []hash.Hash,
) Content {
	out := content{
		hash:     hash,
		section:  section,
		newOwner: newOwner,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Section returns the section
func (obj *content) Section() Section {
	return obj.section
}

// NewOwner returns the newOwner hashes
func (obj *content) NewOwner() []hash.Hash {
	return obj.newOwner
}
