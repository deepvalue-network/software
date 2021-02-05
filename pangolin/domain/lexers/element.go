package lexers

import "github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"

type element struct {
	rule grammar.Rule
	code string
}

func createElement(rule grammar.Rule, code string) Element {
	out := element{
		rule: rule,
		code: code,
	}

	return &out
}

// Rule returns the rule
func (obj *element) Rule() grammar.Rule {
	return obj.rule
}

// Code returns the code
func (obj *element) Code() string {
	return obj.code
}
