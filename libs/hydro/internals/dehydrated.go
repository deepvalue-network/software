package internals

import "github.com/steve-care-software/products/libs/hash"

/*
 * simpleInterface -> dehydrateSimpleStruct
 */

type dehydrateSimpleStruct struct {
	first  string `hydro:"First, Fr"`
	second string `hydro:"Second, Se"`
}

func createSimpleStruct(first string, second string) SimpleInterface {
	out := dehydrateSimpleStruct{
		first:  first,
		second: second,
	}

	return &out
}

// First returns the first element
func (obj *dehydrateSimpleStruct) First() string {
	return obj.first
}

// Second returns the second element
func (obj *dehydrateSimpleStruct) Second() string {
	return obj.second
}

/*
 * complexInterface -> dehydrateComplexStruct
 */

type dehydrateComplexStruct struct {
	simple  SimpleInterface `hydro:"Simple, Sim"`
	another uint            `hydro:"Another, An"`
	hash    hash.Hash       `hydro:"Hash, Hsh"`
}

func createComplexStruct(
	simple SimpleInterface,
	another uint,
	hash hash.Hash,
) ComplexInterface {
	out := dehydrateComplexStruct{
		simple:  simple,
		another: another,
		hash:    hash,
	}

	return &out
}

// Simple returns the simple instance
func (obj *dehydrateComplexStruct) Simple() SimpleInterface {
	return obj.simple
}

// Another returns the another number
func (obj *dehydrateComplexStruct) Another() uint {
	return obj.another
}

// Hash returns the hash
func (obj *dehydrateComplexStruct) Hash() hash.Hash {
	return obj.hash
}
