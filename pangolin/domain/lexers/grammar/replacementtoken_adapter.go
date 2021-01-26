package grammar

import (
	"errors"
	"fmt"
	"regexp"
)

type replacementTokenAdapter struct {
	builder ReplacementTokenBuilder
}

func createReplacementTokenAdapter(builder ReplacementTokenBuilder) ReplacementTokenAdapter {
	out := replacementTokenAdapter{
		builder: builder,
	}

	return &out
}

// ToReplacementTokens converts a grammar script and a replacement pattern to a ReplacementToken map
func (app *replacementTokenAdapter) ToReplacementTokens(script string, replacementPattern string) ([]ReplacementToken, error) {
	pattern := regexp.MustCompile(replacementPattern)
	matches := pattern.FindAllStringSubmatch(script, -1)

	out := []ReplacementToken{}
	for _, oneMatch := range matches {

		if len(oneMatch) != 3 {
			str := fmt.Sprintf("the ReplacementToken pattern (%s) was expected to contain %d element per match, %d returned", oneMatch[0], 3, len(oneMatch))
			return nil, errors.New(str)
		}

		repl, err := app.builder.Create().WithToGrammar(oneMatch[1]).WithFromToken(oneMatch[2]).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, repl)
	}

	return out, nil
}
