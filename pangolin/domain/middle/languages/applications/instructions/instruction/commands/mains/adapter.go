package mains

import (
	application_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	instructionAdapter application_instruction.Adapter
	builder            Builder
	instructionBuilder InstructionBuilder
}

func createAdapter(
	instructionAdapter application_instruction.Adapter,
	builder Builder,
	instructionBuilder InstructionBuilder,
) Adapter {
	out := adapter{
		instructionAdapter: instructionAdapter,
		builder:            builder,
		instructionBuilder: instructionBuilder,
	}

	return &out
}

// ToMain converts a parsed main command to a main instance
func (app *adapter) ToMain(parsed parsers.MainCommand) (Main, error) {
	instructions := []Instruction{}
	mainCommandInstructions := parsed.Instructions()
	for _, oneMainCommandIns := range mainCommandInstructions {
		parsedLangIns := oneMainCommandIns.Instruction()
		langIns, err := app.instructionAdapter.ToInstruction(parsedLangIns)
		if err != nil {
			return nil, err
		}

		builder := app.instructionBuilder.Create().WithInstruction(langIns)
		if oneMainCommandIns.HasScopes() {
			scopes := []bool{}
			parsedScopes := oneMainCommandIns.Scopes().All()
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
	return app.builder.Create().
		WithVariable(variable).
		WithInstructions(instructions).
		Now()
}
