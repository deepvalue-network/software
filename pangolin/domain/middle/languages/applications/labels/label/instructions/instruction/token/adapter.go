package token

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	builder          Builder
	codeMatchBuilder CodeMatchBuilder
	codeBuilder      CodeBuilder
}

func createAdapter(
	builder Builder,
	codeMatchBuilder CodeMatchBuilder,
	codeBuilder CodeBuilder,
) Adapter {
	out := adapter{
		builder:          builder,
		codeMatchBuilder: codeMatchBuilder,
		codeBuilder:      codeBuilder,
	}

	return &out
}

// ToToken converts a parsed token to a token instance
func (app *adapter) ToToken(parsed parsers.Token) (Token, error) {
	builder := app.builder.Create()
	if parsed.IsCodeMatch() {
		parsedCodeMatch := parsed.CodeMatch()
		content := parsedCodeMatch.Content()
		section := parsedCodeMatch.Section()
		patternVariables := parsedCodeMatch.PatternVariables()
		codeMatch, err := app.codeMatchBuilder.Create().WithPatterns(patternVariables).WithSectionName(section).WithReturn(content).Now()
		if err != nil {
			return nil, err
		}

		builder.WithCodeMatch(codeMatch)
	}

	if parsed.IsTokenSection() {
		codeBuilder := app.codeBuilder.Create()
		parsedTokenSection := parsed.TokenSection()
		if parsedTokenSection.IsVariableName() {
			variableName := parsedTokenSection.VariableName()
			code, err := codeBuilder.WithReturn(variableName).Now()
			if err != nil {
				return nil, err
			}

			builder.WithCode(code)
		}

		if parsedTokenSection.IsSpecific() {
			specific := parsedTokenSection.Specific()
			variableName := specific.VariableName()
			patternVariable := specific.PatternVariable()

			codeBuilder.WithPattern(patternVariable).WithReturn(variableName)
			if specific.HasAmount() {
				amount := specific.Amount()
				codeBuilder.WithAmount(amount)
			}

			code, err := codeBuilder.Now()
			if err != nil {
				return nil, err
			}

			builder.WithCode(code)
		}
	}

	return builder.Now()
}
