package grammar

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type validateTokenFn func(index int, str string) bool

type tokensAdapter struct {
	rawTokenBuilder         RawTokenBuilder
	tokenRuleBuilder        TokenRuleBuilder
	blocksBuilder           TokenBlocksBuilder
	tokenBuilder            TokenBuilder
	tokensBuilder           TokensBuilder
	replacementTokenAdapter ReplacementTokenAdapter
}

func createTokensAdapter(
	rawTokenBuilder RawTokenBuilder,
	tokenRuleBuilder TokenRuleBuilder,
	blocksBuilder TokenBlocksBuilder,
	tokenBuilder TokenBuilder,
	tokensBuilder TokensBuilder,
	replacementTokenAdapter ReplacementTokenAdapter,
) TokensAdapter {
	out := tokensAdapter{
		rawTokenBuilder:         rawTokenBuilder,
		tokenRuleBuilder:        tokenRuleBuilder,
		blocksBuilder:           blocksBuilder,
		tokenBuilder:            tokenBuilder,
		tokensBuilder:           tokensBuilder,
		replacementTokenAdapter: replacementTokenAdapter,
	}

	return &out
}

// ToTokens converts a script to Tokens instance
func (app *tokensAdapter) ToTokens(
	script string,
	tokenPattern string,
	replacementPattern string,
	grammarName string,
	extends map[string]Grammar,
	rules []Rule,
) (Tokens, error) {
	pattern := regexp.MustCompile(tokenPattern)
	matches := pattern.FindAllStringSubmatch(script, -1)

	tokens := map[string]Token{}
	for _, oneMatch := range matches {
		orStatements, err := app.findOrStatements(oneMatch[2])
		if err != nil {
			return nil, err
		}

		tokenBlocks := []TokenBlocks{}
		for _, oneStatement := range orStatements {
			blk, err := app.createTokenBlocks(oneStatement, rules, grammarName, extends)
			if err != nil {
				return nil, err
			}

			tokenBlocks = append(tokenBlocks, blk)
		}

		tok, err := app.tokenBuilder.Create().WithName(oneMatch[1]).WithBlocks(tokenBlocks).Now()
		if err != nil {
			return nil, err
		}

		tokens[tok.Name()] = tok
	}

	repl, err := app.replacementTokenAdapter.ToReplacementTokens(script, replacementPattern)
	if err != nil {
		return nil, err
	}

	builder := app.tokensBuilder.Create().WithTokens(tokens)
	if len(repl) > 0 {
		builder.WithReplacements(repl)
	}

	return builder.Now()
}

func (app *tokensAdapter) createTokenBlocks(str string, rules []Rule, grammarName string, extends map[string]Grammar) (TokenBlocks, error) {
	tokenRules, err := app.filterRules(str, rules, grammarName)
	if err != nil {
		return nil, err
	}

	tokPattern := fmt.Sprintf("((%s)\\.)?(%s)(\\@(%s))?", grammarNamePattern, tokenPattern, tokenPattern)
	tokens, err := app.buildRawTokensWithPotentialGrammar(tokPattern, str, grammarName, extends)
	if err != nil {
		return nil, err
	}

	multipleOptionalPattern := fmt.Sprintf(singleFormat, multipleOptional)
	multipleOptionals, err := app.buildRawTokens(multipleOptionalPattern, str, grammarName, nil)
	if err != nil {
		return nil, err
	}

	multipleMandatoryPattern := fmt.Sprintf(singleFormat, multipleMandatory)
	multipleMandatories, err := app.buildRawTokens(multipleMandatoryPattern, str, grammarName, nil)
	if err != nil {
		return nil, err
	}

	optionalPattern := fmt.Sprintf(singleFormat, optional)
	optionals, err := app.buildRawTokens(optionalPattern, str, grammarName, nil)
	if err != nil {
		return nil, err
	}

	tokenBuilder := app.blocksBuilder.Create()
	if len(tokenRules) > 0 {
		tokenBuilder.WithRules(tokenRules)
	}

	if len(tokens) > 0 {
		tokenBuilder.WithTokens(tokens)
	}

	if len(multipleOptionals) > 0 {
		tokenBuilder.WithMultipleOptionals(multipleOptionals)
	}

	if len(multipleMandatories) > 0 {
		tokenBuilder.WithMultipleMandatories(multipleMandatories)
	}

	if len(optionals) > 0 {
		tokenBuilder.WithOptionals(optionals)
	}

	return tokenBuilder.Now()
}

func (app *tokensAdapter) filterRules(str string, rules []Rule, grammarName string) ([]TokenRule, error) {
	createTokenRules := func(rawTokens []RawToken, rules map[string]Rule) ([]TokenRule, error) {
		tokenRules := []TokenRule{}
		for _, oneRawToken := range rawTokens {
			if rule, ok := rules[oneRawToken.Value()]; ok {
				tok, err := app.tokenRuleBuilder.Create().WithRawToken(oneRawToken).WithRule(rule).Now()
				if err != nil {
					return nil, err
				}

				tokenRules = append(tokenRules, tok)
				continue
			}

			str := fmt.Sprintf("the rule (ID: %s) is not defined", oneRawToken.Value())
			return nil, errors.New(str)
		}

		return tokenRules, nil
	}

	mp := map[string]Rule{}
	for _, oneRule := range rules {
		mp[oneRule.Name()] = oneRule
	}

	validate := func(idx int, str string) bool {
		// make sure there is no letters before the match:
		if idx > 0 {
			prev := idx - 1
			return strings.Index(lowercaseLetters, string(str[prev])) == -1
		}

		return true
	}

	ruleWithLocalNamePattern := fmt.Sprintf("(%s)(\\@(%s))?", rulePattern, rulePattern)
	rawTokens, err := app.buildRawTokens(ruleWithLocalNamePattern, str, grammarName, validate)
	if err != nil {
		return nil, err
	}

	return createTokenRules(rawTokens, mp)
}

func (app *tokensAdapter) buildRawTokensWithPotentialGrammar(patternStr string, str string, grammarName string, extends map[string]Grammar) ([]RawToken, error) {
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(str, -1)
	indexes := pattern.FindAllStringSubmatchIndex(str, -1)
	tokens := []RawToken{}
	for index, oneMatch := range matches {
		grName := oneMatch[2]
		tokenBuilder := app.rawTokenBuilder.Create().WithGrammar(grammarName)
		if grName != "" {
			if gr, ok := extends[grName]; ok {
				tokenBuilder.WithGrammar(gr.Name())
			} else {
				str := fmt.Sprintf("the Grammar (%s) specified in the matched Token (%s) is not defined", grName, oneMatch[0])
				return nil, errors.New(str)
			}
		}

		trimmedValue := strings.TrimSpace(oneMatch[3])
		rawTokenBuilder := tokenBuilder.WithCode(str).WithValue(trimmedValue).WithIndex(indexes[index][0])
		if len(oneMatch) >= 5 {
			trimmedName := strings.TrimSpace(oneMatch[5])
			if trimmedName != "" {
				rawTokenBuilder.WithName(trimmedName)
			}
		}

		rawToken, err := rawTokenBuilder.Now()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, rawToken)
	}

	return tokens, nil
}

func (app *tokensAdapter) buildRawTokens(patternStr string, str string, grammarName string, fn validateTokenFn) ([]RawToken, error) {
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(str, -1)
	indexes := pattern.FindAllStringSubmatchIndex(str, -1)

	tokens := []RawToken{}
	for index, oneMatch := range matches {
		if fn != nil && !fn(indexes[index][0], str) {
			continue
		}

		trimmedValue := strings.TrimSpace(oneMatch[0])
		rawTokenBuilder := app.rawTokenBuilder.Create().WithCode(str).WithValue(trimmedValue).WithIndex(indexes[index][0]).WithGrammar(grammarName)
		if len(oneMatch) > 1 {
			trimmedValue := strings.TrimSpace(oneMatch[1])
			if trimmedValue != "" {
				rawTokenBuilder.WithValue(trimmedValue)
			}

		}

		if len(oneMatch) >= 3 {
			trimmedName := strings.TrimSpace(oneMatch[3])
			if trimmedName != "" {
				rawTokenBuilder.WithName(trimmedName)
			}
		}

		rawToken, err := rawTokenBuilder.Now()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, rawToken)
	}

	return tokens, nil
}

func (app *tokensAdapter) findOrStatements(str string) ([]string, error) {
	out := []string{}
	sections := strings.Split(str, or)
	for _, oneSection := range sections {
		out = append(out, strings.TrimSpace(oneSection))
	}

	return out, nil
}
