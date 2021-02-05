package internals

import "github.com/deepvalue-network/software/libs/hash"

/*
 * simpleInterface -> dehydrateSimpleStruct
 */

// DehydrateSimpleStruct repreents a dehydrated simple struct
type DehydrateSimpleStruct struct {
	first  string `hydro:"First, Fr"`
	second string `hydro:"Second, Se"`
}

func createSimpleStruct(first string, second string) SimpleInterface {
	out := DehydrateSimpleStruct{
		first:  first,
		second: second,
	}

	return &out
}

// First returns the first element
func (obj *DehydrateSimpleStruct) First() string {
	return obj.first
}

// Second returns the second element
func (obj *DehydrateSimpleStruct) Second() string {
	return obj.second
}

/*
 * complexInterface -> dehydrateComplexStruct
 */

// DehydrateComplexStruct represents a dehydrated complex struct
type DehydrateComplexStruct struct {
	simple  SimpleInterface `hydro:"Simple, Sim"`
	another uint            `hydro:"Another, An"`
	hash    hash.Hash       `hydro:"Hash, Hsh"`
}

func createComplexStruct(
	simple SimpleInterface,
	another uint,
	hash hash.Hash,
) ComplexInterface {
	out := DehydrateComplexStruct{
		simple:  simple,
		another: another,
		hash:    hash,
	}

	return &out
}

// Simple returns the simple instance
func (obj *DehydrateComplexStruct) Simple() SimpleInterface {
	return obj.simple
}

// Another returns the another number
func (obj *DehydrateComplexStruct) Another() uint {
	return obj.another
}

// Hash returns the hash
func (obj *DehydrateComplexStruct) Hash() hash.Hash {
	return obj.hash
}
