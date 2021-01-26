package internals

import "github.com/steve-care-software/products/libs/hash"

// NewComplex creates a new complex instance
func NewComplex(simple SimpleInterface, another uint, hash hash.Hash) (ComplexInterface, error) {
	return createComplexStruct(simple, another, hash), nil
}

// NewSimple creates a new simple instance
func NewSimple(first string, second string) (SimpleInterface, error) {
	return createSimpleStruct(first, second), nil
}

// ComplexStructOnHydrateEventFn represents the complex struct onHydrate event func
func ComplexStructOnHydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if hsh, ok := ins.(hash.Hash); ok {
		return hsh.String(), nil
	}

	return nil, nil
}

// ComplexStructOnDehydrateEventFn represents the complex struct onDehydrate event func
func ComplexStructOnDehydrateEventFn(ins interface{}, fieldName string, structName string) (interface{}, error) {
	if fieldName == "Hsh" && structName == "HydrateComplexStruct" {
		if hashAsString, ok := ins.(string); ok {
			hash, err := hash.NewAdapter().FromString(hashAsString)
			if err != nil {
				return nil, err
			}

			return *hash, nil
		}
	}

	return nil, nil
}

// SimpleInterface represents a simple interface
type SimpleInterface interface {
	First() string
	Second() string
}

// ComplexInterface represents a complex interface
type ComplexInterface interface {
	Simple() SimpleInterface
	Another() uint
	Hash() hash.Hash
}

// HydrateSimpleStruct represents an hydrate simple struct
type HydrateSimpleStruct struct {
	Fr string `json:"first" hydro:"0"`
	Se string `json:"second" hydro:"1"`
}

// HydrateComplexStruct represents an hydrate complex struct
type HydrateComplexStruct struct {
	Sim *HydrateSimpleStruct `json:"simple" hydro:"0"`
	An  uint                  `json:"another" hydro:"1"`
	Hsh string                `json:"hash" hydro:"2"`
}
