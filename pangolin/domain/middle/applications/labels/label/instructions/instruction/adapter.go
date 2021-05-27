package instruction

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels/label/instructions/instruction"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels/label/instructions/instruction/token"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	labelAdapter    label_instruction.Adapter
	languageAdapter language_instruction.Adapter
	tokenAdapter    token.Adapter
	builder         Builder
}

func createAdapter(
	labelAdapter label_instruction.Adapter,
	languageAdapter language_instruction.Adapter,
	tokenAdapter token.Adapter,
	builder Builder,
) Adapter {
	out := adapter{
		labelAdapter:    labelAdapter,
		languageAdapter: languageAdapter,
		tokenAdapter:    tokenAdapter,
		builder:         builder,
	}

	return &out
}

// ToInstruction converts a parsed language label instruction to instruction instance
func (app *adapter) ToInstruction(parsed parsers.LanguageLabelInstruction) (Instruction, error) {
	builder := app.builder.Create()
	if parsed.IsLanguageInstruction() {
		parsedLangIns := parsed.LanguageInstruction()
		langIns, err := app.languageAdapter.ToInstruction(parsedLangIns)
		if err != nil {
			return nil, err
		}

		builder.WithLanguage(langIns)
	}

	if parsed.IsLabelInstruction() {
		parsedLabelIns := parsed.LabelInstruction()
		labelIns, err := app.labelAdapter.ToInstruction(parsedLabelIns)
		if err != nil {
			return nil, err
		}

		builder.WithLabel(labelIns)
	}

	if parsed.IsToken() {
		parsedToken := parsed.Token()
		tok, err := app.tokenAdapter.ToToken(parsedToken)
		if err != nil {
			return nil, err
		}

		builder.WithToken(tok)
	}

	return builder.Now()
}
