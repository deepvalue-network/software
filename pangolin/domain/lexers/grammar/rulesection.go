package grammar

import (
	"regexp"
	"strings"
)

type ruleSection struct {
	constant RawToken
	pattern  RawToken
}

func createRuleSectionWithConstant(constant RawToken) RuleSection {
	return createRuleSectionInternally(constant, nil)
}

func createRuleSectionWithPattern(pattern RawToken) RuleSection {
	return createRuleSectionInternally(nil, pattern)
}

func createRuleSectionInternally(constant RawToken, pattern RawToken) RuleSection {
	out := ruleSection{
		constant: constant,
		pattern:  pattern,
	}

	return &out
}

// HasConstant returns true if there is a constant, false otherwise
func (obj *ruleSection) HasConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *ruleSection) Constant() RawToken {
	return obj.constant
}

// HasPattern returns true if there is a pattern, false otherwise
func (obj *ruleSection) HasPattern() bool {
	return obj.pattern != nil
}

// Pattern returns the pattern, if any
func (obj *ruleSection) Pattern() RawToken {
	return obj.pattern
}

// FindFirst finds the first occurence of the string, using the rule pattern
func (obj *ruleSection) FindFirst(str string) (string, error) {
	if obj.HasPattern() {
		regexp, err := regexp.Compile(obj.pattern.Name())
		if err != nil {
			return "", nil
		}

		return regexp.FindString(str), nil
	}

	constant := obj.Constant().Name()
	pos := strings.Index(str, constant)
	if pos != -1 {
		return constant, nil
	}

	return "", nil
}
