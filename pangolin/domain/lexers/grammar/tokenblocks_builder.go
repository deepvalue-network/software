package grammar

import (
	"errors"
	"math"
	"sort"
)

type tokenBlocksBuilder struct {
	blockBuilder        TokenBlockBuilder
	sectionBuilder      TokenSectionBuilder
	rules               []TokenRule
	tokens              []RawToken
	optionals           []RawToken
	multipleOptionals   []RawToken
	multipleMandatories []RawToken
}

func createTokenBlocksBuilder(blockBuilder TokenBlockBuilder, sectionBuilder TokenSectionBuilder) TokenBlocksBuilder {
	out := tokenBlocksBuilder{
		blockBuilder:        blockBuilder,
		sectionBuilder:      sectionBuilder,
		rules:               nil,
		tokens:              nil,
		optionals:           nil,
		multipleOptionals:   nil,
		multipleMandatories: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBlocksBuilder) Create() TokenBlocksBuilder {
	return createTokenBlocksBuilder(app.blockBuilder, app.sectionBuilder)
}

// WithRules add rules to the builder
func (app *tokenBlocksBuilder) WithRules(rules []TokenRule) TokenBlocksBuilder {
	app.rules = rules
	return app
}

// WithTokens add tokens to the builder
func (app *tokenBlocksBuilder) WithTokens(tokens []RawToken) TokenBlocksBuilder {
	app.tokens = tokens
	return app
}

// WithOptionals add optionals to the builder
func (app *tokenBlocksBuilder) WithOptionals(optionals []RawToken) TokenBlocksBuilder {
	app.optionals = optionals
	return app
}

// WithMultipleOptionals add multipleOptionals to the builder
func (app *tokenBlocksBuilder) WithMultipleOptionals(multipleOptionals []RawToken) TokenBlocksBuilder {
	app.multipleOptionals = multipleOptionals
	return app
}

// WithMultipleMandatories add multipleMandatories to the builder
func (app *tokenBlocksBuilder) WithMultipleMandatories(multipleMandatories []RawToken) TokenBlocksBuilder {
	app.multipleMandatories = multipleMandatories
	return app
}

// Now builds a new TokenBlocks instance
func (app *tokenBlocksBuilder) Now() (TokenBlocks, error) {
	findNextRawToken := func(tok RawToken, tokens []RawToken) RawToken {
		index := tok.Index()
		next := index + 5000000000
		var out RawToken
		for _, oneToken := range tokens {
			if oneToken == nil {
				continue
			}

			tokIndex := oneToken.Index()
			if tokIndex < next && tokIndex > index {
				next = tokIndex
				out = oneToken
			}
		}

		return out
	}

	findNextTokenOrRule := func(tok RawToken, rules []RawToken, tokens []RawToken) RawToken {
		nextTok := findNextRawToken(tok, tokens)
		nextRule := findNextRawToken(tok, rules)

		nextTokIsNil := nextTok == nil
		nextRuleIsNil := nextRule == nil
		if nextTokIsNil && nextRuleIsNil {
			return nil
		}

		if nextTokIsNil {
			return nextRule
		}

		if nextRuleIsNil {
			return nextTok
		}

		ruleIndex := nextRule.Index()
		tokIndex := nextTok.Index()
		if ruleIndex < tokIndex {
			return nextRule
		}

		return nextTok
	}

	isWithin := func(begin int, end int, rawToken RawToken) bool {
		index := rawToken.Index()
		if index > begin {
			return index <= end
		}

		return false
	}

	initTokenBlockBuiler := func(tok RawToken, tokens []RawToken, rulesTokens []RawToken) TokenBlockBuilder {
		tokenBlockBuilder := app.blockBuilder.Create()

		next := findNextTokenOrRule(tok, rulesTokens, tokens)
		nextOptional := findNextRawToken(tok, app.optionals)
		nextMultipleOptional := findNextRawToken(tok, app.multipleOptionals)
		nextMultipleMandatory := findNextRawToken(tok, app.multipleMandatories)

		begin := tok.Index()
		end := math.MaxInt32
		if next != nil {
			end = next.Index()
		}

		if nextOptional != nil {
			if isWithin(begin, end, nextOptional) {
				tokenBlockBuilder.WithOptional(nextOptional)
			}
		}

		if nextMultipleOptional != nil {
			if isWithin(begin, end, nextMultipleOptional) {
				tokenBlockBuilder.WithMultipleOptional(nextMultipleOptional)
			}
		}

		if nextMultipleMandatory != nil {
			if isWithin(begin, end, nextMultipleMandatory) {
				tokenBlockBuilder.WithMultipleMandatory(nextMultipleMandatory)
			}
		}

		return tokenBlockBuilder
	}

	rulesRawTokens := []RawToken{}
	for _, oneRule := range app.rules {
		rulesRawTokens = append(rulesRawTokens, oneRule.RawToken())
	}

	tokenBlocks := []TokenBlock{}
	for _, oneToken := range app.tokens {
		section, err := app.sectionBuilder.Create().WithToken(oneToken).Now()
		if err != nil {
			return nil, err
		}

		blockToken, err := initTokenBlockBuiler(oneToken, app.tokens, rulesRawTokens).WithSection(section).Now()
		if err != nil {
			return nil, err
		}

		tokenBlocks = append(tokenBlocks, blockToken)
	}

	for index, oneRuleToken := range rulesRawTokens {
		section, err := app.sectionBuilder.Create().WithRule(app.rules[index]).Now()
		if err != nil {
			return nil, err
		}

		blockToken, err := initTokenBlockBuiler(oneRuleToken, app.tokens, rulesRawTokens).WithSection(section).Now()
		if err != nil {
			return nil, err
		}

		tokenBlocks = append(tokenBlocks, blockToken)
	}

	if len(tokenBlocks) <= 0 {
		return nil, errors.New("[]TokenBlock were expected in order to build a TokenBlocks instance")
	}

	keys := []int{}
	tokenBlocksMP := map[int]TokenBlock{}
	for _, oneTokenBlock := range tokenBlocks {
		section := oneTokenBlock.Section()
		index := -1
		if section.HasRule() {
			index = section.Rule().RawToken().Index()
		}

		if section.HasToken() {
			index = section.Token().Index()
		}

		keys = append(keys, index)
		tokenBlocksMP[index] = oneTokenBlock
	}

	sort.Ints(keys)
	orderedTokenBlocks := []TokenBlock{}
	for _, index := range keys {
		orderedTokenBlocks = append(orderedTokenBlocks, tokenBlocksMP[index])
	}

	return createTokenBlocks(orderedTokenBlocks), nil
}
