package parsers

type languageTestInstruction struct {
	lang        LanguageInstructionCommon
	test        TestInstruction
	isInterpret bool
}

func createLanguageTestInstructionWithLanguage(
	lang LanguageInstructionCommon,
) LanguageTestInstruction {
	return createLanguageTestInstructionInternally(lang, nil, false)
}

func createLanguageTestInstructionWithTest(
	test TestInstruction,
) LanguageTestInstruction {
	return createLanguageTestInstructionInternally(nil, test, false)
}

func createLanguageTestInstructionWithIntepret() LanguageTestInstruction {
	return createLanguageTestInstructionInternally(nil, nil, true)
}

func createLanguageTestInstructionInternally(
	lang LanguageInstructionCommon,
	test TestInstruction,
	isInterpret bool,
) LanguageTestInstruction {
	out := languageTestInstruction{
		lang:        lang,
		test:        test,
		isInterpret: isInterpret,
	}

	return &out
}

// IsLanguageInstruction returns true if there is a language instruction, false otherwise
func (obj *languageTestInstruction) IsLanguageInstruction() bool {
	return obj.lang != nil
}

// LanguageInstruction returns the language instruction, if any
func (obj *languageTestInstruction) LanguageInstruction() LanguageInstructionCommon {
	return obj.lang
}

// IsTestInstruction returns true if there is a test instruction, false otherwise
func (obj *languageTestInstruction) IsTestInstruction() bool {
	return obj.test != nil
}

// TestInstruction returns the test instruction, if any
func (obj *languageTestInstruction) TestInstruction() TestInstruction {
	return obj.test
}

// IsInterpret returns true if the instruction is interpret
func (obj *languageTestInstruction) IsInterpret() bool {
	return obj.isInterpret
}
