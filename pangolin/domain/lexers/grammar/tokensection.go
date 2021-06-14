package grammar

type tokenSection struct {
	rule  TokenRule
	token RawToken
}

func createTokenSectionWithToken(token RawToken) TokenSection {
	return createTokenSectionInternally(token, nil)
}

func createTokenSectionWithRule(rule TokenRule) TokenSection {
	return createTokenSectionInternally(nil, rule)
}

func createTokenSectionInternally(token RawToken, rule TokenRule) TokenSection {
	out := tokenSection{
		rule:  rule,
		token: token,
	}

	return &out
}

// HasRule returns true if there is a rule, false otherwise
func (obj *tokenSection) HasRule() bool {
	return obj.rule != nil
}

// Rule returns the rule, if any
func (obj *tokenSection) Rule() TokenRule {
	return obj.rule
}

// HasToken returns true if there is a token, false otherwise
func (obj *tokenSection) HasToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *tokenSection) Token() RawToken {
	return obj.token
}

// NextRuleToken returns the RuleSection or Token name
func (obj *tokenSection) NextRuleToken() (Rule, string) {
	if obj.HasRule() {
		return obj.rule.Rule(), ""
	}

	return nil, obj.token.Name()
}
