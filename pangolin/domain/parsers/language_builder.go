package parsers

import "errors"

type languageBuilder struct {
	values []LanguageValue
}

func createLanguageBuilder() LanguageBuilder {
	out := languageBuilder{
		values: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageBuilder) Create() LanguageBuilder {
	return createLanguageBuilder()
}

// WithValues add values to the builder
func (app *languageBuilder) WithValues(values []LanguageValue) LanguageBuilder {
	app.values = values
	return app
}

// Now builds a new Language instance
func (app *languageBuilder) Now() (Language, error) {
	if app.values == nil {
		app.values = []LanguageValue{}
	}

	root := ""
	var tokens RelativePath
	var channels RelativePath
	var rules RelativePath
	var logic RelativePath
	input := ""
	output := ""
	extends := []RelativePath{}
	patternMatches := []PatternMatch{}
	targets := []Target{}
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

		if oneValue.IsOutputVariable() {
			output = oneValue.OutputVariable()
			continue
		}

		if oneValue.IsExtends() {
			extends = oneValue.Extends()
			continue
		}

		if oneValue.IsTargets() {
			targets = oneValue.Targets()
			continue
		}
	}

	if root == "" {
		return nil, errors.New("the root pattern is mandatory in order to build a Language instance")
	}

	if tokens == nil {
		return nil, errors.New("the token's RelativePath is mandatory in order to build a Language instance")
	}

	if rules == nil {
		return nil, errors.New("the rule's RelativePath is mandatory in order to build a Language instance")
	}

	if logic == nil {
		return nil, errors.New("the logic's RelativePath is mandatory in order to build a Language instance")
	}

	if input == "" {
		return nil, errors.New("the input variable is mandatory in order to build a Language instance")
	}

	if output == "" {
		return nil, errors.New("the output variable is mandatory in order to build a Language instance")
	}

	if len(targets) <= 0 {
		return nil, errors.New("there must be at least 1 Target in order to build a Language instance")
	}

	if len(patternMatches) <= 0 {
		return nil, errors.New("the patternMatches are mandatory in order to build a Language instance")
	}

	if channels != nil && len(extends) > 0 {
		return createLanguageWithChannelsAndExtends(root, patternMatches, tokens, rules, logic, input, output, targets, channels, extends), nil
	}

	if channels != nil {
		return createLanguageWithChannels(root, patternMatches, tokens, rules, logic, input, output, targets, channels), nil
	}

	if len(extends) > 0 {
		return createLanguageWithExtends(root, patternMatches, tokens, rules, logic, input, output, targets, extends), nil
	}

	return createLanguage(root, patternMatches, tokens, rules, logic, input, output, targets), nil
}
