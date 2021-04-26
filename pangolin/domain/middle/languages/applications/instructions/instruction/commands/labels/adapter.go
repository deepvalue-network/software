package labels

import (
	label_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels/label/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	langLabelInsAdapter label_instruction.Adapter
	builder             Builder
	instructionBuilder  InstructionBuilder
}

func createAdapter(
	langLabelInsAdapter label_instruction.Adapter,
	builder Builder,
	instructionBuilder InstructionBuilder,
) Adapter {
	out := adapter{
		langLabelInsAdapter: langLabelInsAdapter,
		builder:             builder,
		instructionBuilder:  instructionBuilder,
	}

	return &out
}

// ToLabel converts a parsed label command to a label instance
func (app *adapter) ToLabel(parsed parsers.LabelCommand) (Label, error) {
	instructions := []Instruction{}
	parsedInstructions := parsed.Instructions()
	for _, oneInstruction := range parsedInstructions {
		parsedLangLabelIns := oneInstruction.Instruction()
		langLabelIns, err := app.langLabelInsAdapter.ToInstruction(parsedLangLabelIns)
		if err != nil {
			return nil, err
		}

		builder := app.instructionBuilder.Create().WithInstruction(langLabelIns)
		if oneInstruction.HasScopes() {
			scopes := []bool{}
			parsedScopes := oneInstruction.Scopes().All()
			for _, oneScope := range parsedScopes {
				scopes = append(scopes, oneScope.IsExternal())
			}

			builder.WithScopes(scopes)
		}

		ins, err := builder.Now()
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, ins)

	}

	variable := parsed.Variable()
	name := parsed.Name()
	return app.builder.Create().
		WithInstructions(instructions).
		WithName(name).
		WithVariable(variable).
		Now()
}
