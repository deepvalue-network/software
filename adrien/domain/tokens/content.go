package tokens

import "github.com/deepvalue-network/software/adrien/domain/rules"

type content struct {
	token string
	rule  rules.Rule
}

func createContentWithToken(
	token string,
) Content {
	return createContentInternally(token, nil)
}

func createContentWithRule(
	rule rules.Rule,
) Content {
	return createContentInternally("", rule)
}

func createContentInternally(
	token string,
	rule rules.Rule,
) Content {
	out := content{
		token: token,
		rule:  rule,
	}

	return &out
}

// Name returns the name
func (obj *content) Name() string {
	if obj.IsToken() {
		return obj.Token()
	}

	return obj.Rule().Name()
}

// IsToken returns true if there is a token, false otherwise
func (obj *content) IsToken() bool {
	return obj.token != ""
}

// Token returns the token reference
func (obj *content) Token() string {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *content) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule
func (obj *content) Rule() rules.Rule {
	return obj.rule
}
