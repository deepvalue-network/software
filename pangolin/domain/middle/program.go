package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables"
)

type program struct {
	testable testables.Testable
	language applications.Application
}

func createProgramWithTestable(
	testable testables.Testable,
) Program {
	return createProgramInternally(testable, nil)
}

func createProgramWithLanguage(
	language applications.Application,
) Program {
	return createProgramInternally(nil, language)
}

func createProgramInternally(
	testable testables.Testable,
	language applications.Application,
) Program {
	out := program{
		testable: testable,
		language: language,
	}

	return &out
}

// IsTestable returns true if the program is testable, false otherwise
func (obj *program) IsTestable() bool {
	return obj.testable != nil
}

// Testable returns the testable instance, if any
func (obj *program) Testable() testables.Testable {
	return obj.testable
}

// IsLanguage returns true if the program is language, false otherwise
func (obj *program) IsLanguage() bool {
	return obj.language != nil
}

// Language returns the language instance, if any
func (obj *program) Language() applications.Application {
	return obj.language
}
