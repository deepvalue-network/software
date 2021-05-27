package definitions

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

type patternMatchAdapter struct {
	builder PatternMatchBuilder
}

func createPatternMatchAdapter(
	builder PatternMatchBuilder,
) PatternMatchAdapter {
	out := patternMatchAdapter{
		builder: builder,
	}

	return &out
}

// ToPatternMatch converts parsed PatternMatch to a PatternMatch instance
func (app *patternMatchAdapter) ToPatternMatch(parsed parsers.PatternMatch) (PatternMatch, error) {
	pattern := parsed.Pattern()
	builder := app.builder.Create().WithPattern(pattern)

	labels := parsed.Labels()
	if labels.HasEnterLabel() {
		enterLabel := labels.EnterLabel()
		builder.WithEnterLabel(enterLabel)
	}

	if labels.HasExitLabel() {
		exitLabel := labels.ExitLabel()
		builder.WithExitLabel(exitLabel)
	}

	return builder.Now()
}

// ToPatternMatches converts parsed PatternMatches to a PatternMatch instances
func (app *patternMatchAdapter) ToPatternMatches(parsed []parsers.PatternMatch) ([]PatternMatch, error) {
	out := []PatternMatch{}
	for _, oneParsed := range parsed {
		patternMatch, err := app.ToPatternMatch(oneParsed)
		if err != nil {
			return nil, err
		}

		out = append(out, patternMatch)
	}

	return out, nil
}
