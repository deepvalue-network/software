package parsers

import "errors"

type languageLabelInstructionBuilder struct {
	langIns LanguageInstruction
	lblIns  LabelInstruction
	tok     Token
}

func createLanguageLabelInstructionBuilder() LanguageLabelInstructionBuilder {
	out := languageLabelInstructionBuilder{
		langIns: nil,
		lblIns:  nil,
		tok:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageLabelInstructionBuilder) Create() LanguageLabelInstructionBuilder {
	return createLanguageLabelInstructionBuilder()
}

// WithLanguageInstruction adds a language instruction to the builder
func (app *languageLabelInstructionBuilder) WithLanguageInstruction(languageInstruction LanguageInstruction) LanguageLabelInstructionBuilder {
	app.langIns = languageInstruction
	return app
}

// WithLabelInstruction adds a label instruction to the builder
func (app *languageLabelInstructionBuilder) WithLabelInstruction(labelInstruction LabelInstruction) LanguageLabelInstructionBuilder {
	app.lblIns = labelInstruction
	return app
}

// WithToken adds a token to the builder
func (app *languageLabelInstructionBuilder) WithToken(token Token) LanguageLabelInstructionBuilder {
	app.tok = token
	return app
}

// Now builds a new LanguageLabelInstruction instance
func (app *languageLabelInstructionBuilder) Now() (LanguageLabelInstruction, error) {
	if app.langIns != nil {
		return createLanguageLabelInstructionWithLanguageInstruction(app.langIns), nil
	}

	if app.lblIns != nil {
		return createLanguageLabelInstructionWithLabelInstruction(app.lblIns), nil
	}

	if app.tok != nil {
		return createLanguageLabelInstructionWithToken(app.tok), nil
	}

	return nil, errors.New("the LanguageLabelInstruction is invalid")
}
