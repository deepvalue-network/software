package testables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
)

type testable struct {
	executable executables.Executable
	language   definitions.Definition
}

func createTestableWithExecutable(
	executable executables.Executable,
) Testable {
	return createTestableInternally(executable, nil)
}

func createTestableWithLanguage(
	language definitions.Definition,
) Testable {
	return createTestableInternally(nil, language)
}

func createTestableInternally(
	executable executables.Executable,
	language definitions.Definition,
) Testable {
	out := testable{
		executable: executable,
		language:   language,
	}

	return &out
}

// IsExecutable returns true if there is an executable, false otherwise
func (obj *testable) IsExecutable() bool {
	return obj.executable != nil
}

// Executable returns the executable, if any
func (obj *testable) Executable() executables.Executable {
	return obj.executable
}

// IsLanguage returns true if there is a language, false otherwise
func (obj *testable) IsLanguage() bool {
	return obj.language != nil
}

// Language returns the language, if any
func (obj *testable) Language() definitions.Definition {
	return obj.language
}
