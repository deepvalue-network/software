package grammar

type tokens struct {
	tokens       map[string]Token
	replacements []ReplacementToken
}

func createTokens(tokens map[string]Token) (Tokens, error) {
	return createTokensInternally(tokens, nil)
}

func createTokensWithReplacements(tokens map[string]Token, replacements []ReplacementToken) (Tokens, error) {
	return createTokensInternally(tokens, replacements)
}

func createTokensInternally(toks map[string]Token, replacements []ReplacementToken) (Tokens, error) {
	out := tokens{
		tokens:       toks,
		replacements: replacements,
	}

	return &out, nil
}

// Tokens returns the tokens
func (obj *tokens) Tokens() map[string]Token {
	return obj.tokens
}

// Replace replace the tokens
func (obj *tokens) Replace(tok Token) Tokens {
	obj.tokens[tok.Name()] = tok
	return obj
}

// HasReplacements return true if there is a replacements, false otherwise
func (obj *tokens) HasReplacements() bool {
	return obj.replacements != nil
}

// Replacements returns the replacementTokens
func (obj *tokens) Replacements() []ReplacementToken {
	return obj.replacements
}
