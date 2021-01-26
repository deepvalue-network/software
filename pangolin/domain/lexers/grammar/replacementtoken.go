package grammar

type replacementToken struct {
	toGrammar string
	fromToken string
}

func createReplacementToken(toGrammar string, fromToken string) (ReplacementToken, error) {
	out := replacementToken{
		toGrammar: toGrammar,
		fromToken: fromToken,
	}

	return &out, nil
}

// ToGrammar returns the to grammar
func (obj *replacementToken) ToGrammar() string {
	return obj.toGrammar
}

// FromToken returns the from token
func (obj *replacementToken) FromToken() string {
	return obj.fromToken
}
