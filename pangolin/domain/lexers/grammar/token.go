package grammar

import "fmt"

type token struct {
	name      string
	grammar   Grammar
	blocks    []TokenBlocks
	delimiter string
}

func createToken(delimiter string, name string, blocks []TokenBlocks) Token {
	return createTokenInternally(delimiter, name, blocks, nil)
}

func createTokenWithGrammar(delimiter string, name string, blocks []TokenBlocks, grammar Grammar) Token {
	return createTokenInternally(delimiter, name, blocks, grammar)
}

func createTokenInternally(delimiter string, name string, blocks []TokenBlocks, grammar Grammar) Token {
	out := token{
		delimiter: delimiter,
		name:      name,
		grammar:   grammar,
		blocks:    blocks,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Reference returns the reference
func (obj *token) Reference() string {
	if obj.HasGrammar() {
		return fmt.Sprintf("%s%s%s", obj.Grammar().Name(), obj.delimiter, obj.Name())
	}

	return obj.Name()
}

// Blocks returns the blocks
func (obj *token) Blocks() []TokenBlocks {
	return obj.blocks
}

// HasGrammar returns true if there is a grammar, false otherwise
func (obj *token) HasGrammar() bool {
	return obj.grammar != nil
}

// Grammar returns the grammar, if any
func (obj *token) Grammar() Grammar {
	return obj.grammar
}

// SubTokenNames returns the subToken names
func (obj *token) SubTokenNames() []string {
	names := []string{}
	for _, oneBlocks := range obj.blocks {
		subNames := oneBlocks.SubTokenNames()
		for _, oneSubName := range subNames {
			isUnique := true
			for _, oneName := range names {
				if oneSubName == oneName {
					isUnique = false
					break
				}
			}

			if isUnique {
				names = append(names, oneSubName)
			}
		}
	}

	return names
}

// SetName sets the name of the token
func (obj *token) SetName(name string) {
	obj.name = name
}
