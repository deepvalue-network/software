package grammar

import (
	"fmt"
	"regexp"
	"strings"
)

type ruleAdapter struct {
	rawTokenBuilder    RawTokenBuilder
	ruleSectionBuilder RuleSectionBuilder
	ruleBuilder        RuleBuilder
}

func createRuleAdapter(
	rawTokenBuilder RawTokenBuilder,
	ruleSectionBuilder RuleSectionBuilder,
	ruleBuilder RuleBuilder,
) RuleAdapter {
	out := ruleAdapter{
		rawTokenBuilder:    rawTokenBuilder,
		ruleSectionBuilder: ruleSectionBuilder,
		ruleBuilder:        ruleBuilder,
	}

	return &out
}

// ToRules converts a script to []Rule
func (app *ruleAdapter) ToRules(script string, grammarName string) ([]Rule, error) {
	anythingExceptEnd := fmt.Sprintf(anythingExcept, end)
	rulePattern := fmt.Sprintf(
		"(%s)%s%s(%s)%s",
		rulePattern,
		potentialWhitespaces,
		begin,
		anythingExceptEnd,
		end,
	)

	pattern := regexp.MustCompile(rulePattern)
	matches := pattern.FindAllStringSubmatch(script, -1)
	indexes := pattern.FindAllStringSubmatchIndex(script, -1)

	rules := []Rule{}
	for index, oneMatch := range matches {
		// create the sections:
		sections, err := app.createRuleSections(oneMatch[2], indexes[index][2], grammarName)
		if err != nil {
			return nil, err
		}

		// build the rule:
		rule, err := app.ruleBuilder.Create().WithName(oneMatch[1]).WithSections(sections).Now()
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (app *ruleAdapter) createRuleSections(str string, baseIndex int, grammarName string) ([]RuleSection, error) {
	constants, err := app.findConstants(str, baseIndex, grammarName)
	if err != nil {
		return nil, err
	}

	patterns, err := app.findPatterns(str, constants, baseIndex, grammarName)
	if err != nil {
		return nil, err
	}

	sections := []RuleSection{}
	for _, oneConstant := range constants {
		section, err := app.ruleSectionBuilder.Create().WithConstant(oneConstant).Now()
		if err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	for _, onePattern := range patterns {
		section, err := app.ruleSectionBuilder.Create().WithPattern(onePattern).Now()
		if err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}

func (app *ruleAdapter) findPatterns(str string, constants []RawToken, baseIndex int, grammarName string) ([]RawToken, error) {
	replacePattern := func(str string, pattern string) string {
		amount := len(pattern)
		replacement := ""
		for i := 0; i < amount; i++ {
			replacement = fmt.Sprintf("%s%s", replacement, "$")
		}

		return strings.Replace(str, pattern, replacement, 1)
	}

	delimiter := "##########"
	cpyString := str
	for _, oneConstant := range constants {
		cpyString = strings.Replace(cpyString, oneConstant.Code(), delimiter, 1)
	}

	out := []RawToken{}
	patternStrings := strings.Split(cpyString, delimiter)
	for _, onePattern := range patternStrings {
		trimmedPattern := strings.TrimSpace(onePattern)
		if len(trimmedPattern) <= 0 {
			continue
		}

		index := strings.Index(str, onePattern)
		idx := baseIndex + index

		isRules, ruleNames, ruleIndexes := app.findRuleNames(onePattern)
		if isRules {
			for index, oneRuleName := range ruleNames {
				idx := ruleIndexes[index] + baseIndex
				trimmedRuleName := strings.TrimSpace(oneRuleName)
				rawToken, err := app.rawTokenBuilder.Create().WithCode(oneRuleName).WithValue(trimmedRuleName).WithIndex(idx).WithGrammar(grammarName).Now()
				if err != nil {
					return nil, err
				}

				out = append(out, rawToken)
			}
		} else {
			rawToken, err := app.rawTokenBuilder.Create().WithCode(onePattern).WithValue(trimmedPattern).WithIndex(idx).WithGrammar(grammarName).Now()
			if err != nil {
				return nil, err
			}

			out = append(out, rawToken)
		}

		// replace the pattern:
		str = replacePattern(str, onePattern)
	}

	return out, nil
}

func (app *ruleAdapter) findConstants(str string, baseIndex int, grammarName string) ([]RawToken, error) {
	anythingExceptConstantEnd := fmt.Sprintf(anythingExcept, constantEnd)
	patternStr := fmt.Sprintf("%s(%s)%s", constantBegin, anythingExceptConstantEnd, constantEnd)
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(str, -1)
	indexes := pattern.FindAllStringSubmatchIndex(str, -1)

	out := []RawToken{}
	for index, oneMatch := range matches {
		idx := baseIndex + indexes[index][0]
		rawToken, err := app.rawTokenBuilder.Create().WithCode(oneMatch[0]).WithValue(oneMatch[1]).WithIndex(idx).WithGrammar(grammarName).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, rawToken)
	}

	return out, nil
}

func (app *ruleAdapter) findRuleNames(str string) (bool, []string, []int) {
	pattern := regexp.MustCompile(rulePattern)
	matches := pattern.FindAllString(str, -1)
	indexes := pattern.FindAllStringIndex(str, -1)

	if len(matches) > 0 {
		total := strings.Join(matches, " ")
		total = strings.TrimSpace(total)
		str = strings.TrimSpace(str)
		if total == str {
			idx := []int{}
			for _, oneIndex := range indexes {
				idx = append(idx, oneIndex[0])
			}

			return true, matches, idx
		}

		return false, nil, nil
	}

	return false, nil, nil
}
