package parsers

type languageLabelInstruction struct {
	langIns LanguageInstruction
	lblIns  LabelInstruction
	tok     Token
}

func createLanguageLabelInstructionWithLanguageInstruction(
	langIns LanguageInstruction,
) LanguageLabelInstruction {
	return createLanguageLabelInstructionInternally(langIns, nil, nil)
}

func createLanguageLabelInstructionWithLabelInstruction(
	lblIns LabelInstruction,
) LanguageLabelInstruction {
	return createLanguageLabelInstructionInternally(nil, lblIns, nil)
}

func createLanguageLabelInstructionWithToken(
	tok Token,
) LanguageLabelInstruction {
	return createLanguageLabelInstructionInternally(nil, nil, tok)
}

func createLanguageLabelInstructionInternally(
	langIns LanguageInstruction,
	lblIns LabelInstruction,
	tok Token,
) LanguageLabelInstruction {
	out := languageLabelInstruction{
		langIns: langIns,
		lblIns:  lblIns,
		tok:     tok,
	}

	return &out
}

// IsLanguageInstruction returns true if there is a language instruction, false otherwise
func (obj *languageLabelInstruction) IsLanguageInstruction() bool {
	return obj.langIns != nil
}

// LanguageInstruction returns the language instruction, if any
func (obj *languageLabelInstruction) LanguageInstruction() LanguageInstruction {
	return obj.langIns
}

// IsLabelInstruction returns true if there is a label instruction, false otherwise
func (obj *languageLabelInstruction) IsLabelInstruction() bool {
	return obj.lblIns != nil
}

// LabelInstruction returns the label instruction, if any
func (obj *languageLabelInstruction) LabelInstruction() LabelInstruction {
	return obj.lblIns
}

// IsToken returns true if there is a token, false otherwise
func (obj *languageLabelInstruction) IsToken() bool {
	return obj.tok != nil
}

// Token returns the token, if any
func (obj *languageLabelInstruction) Token() Token {
	return obj.tok
}
