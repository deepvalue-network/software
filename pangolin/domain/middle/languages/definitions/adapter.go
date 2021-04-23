package definitions

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type adapter struct {
	patternMatchBuilder PatternMatchBuilder
	builder             Builder
}

func createAdapter(
	patternMatchBuilder PatternMatchBuilder,
	builder Builder,
) Adapter {
	out := adapter{
		patternMatchBuilder: patternMatchBuilder,
		builder:             builder,
	}

	return &out
}

// ToDefinition converts a parsed language definition to a language Definition instance
func (app *adapter) ToDefinition(parsed parsers.LanguageDefinition) (Definition, error) {
	patternMatches := []PatternMatch{}
	matches := parsed.PatternMatches()
	for _, onePatternMatch := range matches {
		pattern := onePatternMatch.Pattern()
		labels := onePatternMatch.Labels()
		matchBuilder := app.patternMatchBuilder.Create().WithPattern(pattern)
		if labels.HasEnterLabel() {
			enter := labels.EnterLabel()
			matchBuilder.WithEnterLabel(enter)
		}

		if labels.HasExitLabel() {
			exit := labels.ExitLabel()
			matchBuilder.WithExitLabel(exit)
		}

		match, err := matchBuilder.Now()
		if err != nil {
			return nil, err
		}

		patternMatches = append(patternMatches, match)
	}

	root := parsed.Root()
	tokens := parsed.Tokens().String()
	rules := parsed.Rules().String()
	logic := parsed.Logic().String()
	input := parsed.Input()
	builder := app.builder.Create().
		WithRoot(root).
		WithTokensPath(tokens).
		WithRulesPath(rules).
		WithLogicsPath(logic).
		WithInputVariable(input).
		WithPatternMatches(patternMatches)

	if parsed.HasChannels() {
		channels := parsed.Channels().String()
		builder.WithChannelsPath(channels)
	}

	if parsed.HasExtends() {
		extends := []string{}
		parsedExtends := parsed.Extends()
		for _, oneParsedExtend := range parsedExtends {
			extends = append(extends, oneParsedExtend.String())
		}

		builder.WithExtends(extends)
	}

	return builder.Now()
}
