package grammar

type tokenRule struct {
	rule     Rule
	rawToken RawToken
}

func createTokenRule(rule Rule, rawToken RawToken) TokenRule {
	out := tokenRule{
		rule:     rule,
		rawToken: rawToken,
	}

	return &out
}

// Rule returns the rule
func (obj *tokenRule) Rule() Rule {
	return obj.rule
}

// RawToken returns the rawToken
func (obj *tokenRule) RawToken() RawToken {
	return obj.rawToken
}
