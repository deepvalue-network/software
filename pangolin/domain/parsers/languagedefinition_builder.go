package parsers

import "errors"

type languageDefinitionBuilder struct {
	values []LanguageValue
}

func createLanguageDefinitionBuilder() LanguageDefinitionBuilder {
	out := languageDefinitionBuilder{
		values: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageDefinitionBuilder) Create() LanguageDefinitionBuilder {
	return createLanguageDefinitionBuilder()
}

// WithValues add values to the builder
func (app *languageDefinitionBuilder) WithValues(values []LanguageValue) LanguageDefinitionBuilder {
	app.values = values
	return app
}

// Now builds a new LanguageDefinition instance
func (app *languageDefinitionBuilder) Now() (LanguageDefinition, error) {
	if app.values == nil {
		app.values = []LanguageValue{}
	}

	root := ""
	var tokens RelativePath
	var channels RelativePath
	var rules RelativePath
	var logic RelativePath
	input := ""
	extends := []RelativePath{}
	patternMatches := []PatternMatch{}
	for _, oneValue := range app.values {
		if oneValue.IsRoot() {
			root = oneValue.Root()
			continue
		}

		if oneValue.IsPatternMatches() {
			patternMatches = oneValue.PatternMatches()
			continue
		}

		if oneValue.IsTokens() {
			tokens = oneValue.Tokens()
			continue
		}

		if oneValue.IsChannels() {
			channels = oneValue.Channels()
			continue
		}

		if oneValue.IsRules() {
			rules = oneValue.Rules()
			continue
		}

		if oneValue.IsLogic() {
			logic = oneValue.Logic()
			continue
		}

		if oneValue.IsInputVariable() {
			input = oneValue.InputVariable()
			continue
		}

		if oneValue.IsExtends() {
			extends = oneValue.Extends()
			continue
		}
	}

	if root == "" {
		return nil, errors.New("the root pattern is mandatory in order to build a LanguageDefinition instance")
	}

	if tokens == nil {
		return nil, errors.New("the token's RelativePath is mandatory in order to build a LanguageDefinition instance")
	}

	if rules == nil {
		return nil, errors.New("the rule's RelativePath is mandatory in order to build a LanguageDefinition instance")
	}

	if logic == nil {
		return nil, errors.New("the logic's RelativePath is mandatory in order to build a LanguageDefinition instance")
	}

	if input == "" {
		return nil, errors.New("the input variable is mandatory in order to build a LanguageDefinition instance")
	}

	if len(patternMatches) <= 0 {
		return nil, errors.New("the patternMatches are mandatory in order to build a LanguageDefinition instance")
	}

	if channels != nil && len(extends) > 0 {
		return createLanguageDefinitionWithChannelsAndExtends(root, patternMatches, tokens, rules, logic, input, channels, extends), nil
	}

	if channels != nil {
		return createLanguageDefinitionWithChannels(root, patternMatches, tokens, rules, logic, input, channels), nil
	}

	if len(extends) > 0 {
		return createLanguageDefinitionWithExtends(root, patternMatches, tokens, rules, logic, input, extends), nil
	}

	return createLanguageDefinition(root, patternMatches, tokens, rules, logic, input), nil
}
