package tests

import (
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests/test/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	testInstructionAdapter test_instruction.Adapter
	builder                Builder
	instructionBuilder     InstructionBuilder
}

func createAdapter(
	testInstructionAdapter test_instruction.Adapter,
	builder Builder,
	instructionBuilder InstructionBuilder,
) Adapter {
	out := adapter{
		testInstructionAdapter: testInstructionAdapter,
		builder:                builder,
		instructionBuilder:     instructionBuilder,
	}

	return &out
}

// ToTest converts a parsed test command to a test instance
func (app *adapter) ToTest(parsed parsers.TestCommand) (Test, error) {
	instructions := []Instruction{}
	parsedInstructions := parsed.Instructions()
	for _, oneInstruction := range parsedInstructions {
		parsedTestIns := oneInstruction.Instruction()
		testIns, err := app.testInstructionAdapter.ToInstruction(parsedTestIns)
		if err != nil {
			return nil, err
		}

		builder := app.instructionBuilder.Create().WithInstruction(testIns)
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

	name := parsed.Name()
	variable := parsed.Variable()
	return app.builder.Create().
		WithName(name).
		WithVariable(variable).
		WithInstructions(instructions).
		Now()
}
