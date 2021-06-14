package tokens

import (
	"errors"
	"fmt"
)

type tokens struct {
	list []Token
	mp   map[string]Token
}

func createTokens(
	list []Token,
	mp map[string]Token,
) Tokens {
	out := tokens{
		list: list,
		mp:   mp,
	}

	return &out
}

// All returns the tokens list
func (obj *tokens) All() []Token {
	return obj.list
}

// Find returns a token by name, if any
func (obj *tokens) Find(name string) (Token, error) {
	if tok, ok := obj.mp[name]; ok {
		return tok, nil
	}

	str := fmt.Sprintf("the token (name: %s) is not declared", name)
	return nil, errors.New(str)
}
