package grammar

type tokenBlock struct {
	section           TokenSection
	optional          RawToken
	multipleOptional  RawToken
	multipleMandatory RawToken
}

func createTokenBlock(section TokenSection) TokenBlock {
	return createTokenBlockInternally(section, nil, nil, nil)
}

func createTokenBlockWithOptional(section TokenSection, optional RawToken) TokenBlock {
	return createTokenBlockInternally(section, optional, nil, nil)
}

func createTokenBlockWithMultipleOptional(section TokenSection, multipleOptional RawToken) TokenBlock {
	return createTokenBlockInternally(section, nil, multipleOptional, nil)
}

func createTokenBlockWithMultipleMandatory(section TokenSection, multipleMandatory RawToken) TokenBlock {
	return createTokenBlockInternally(section, nil, nil, multipleMandatory)
}

func createTokenBlockInternally(section TokenSection, optional RawToken, multipleOptional RawToken, multipleMandatory RawToken) TokenBlock {
	out := tokenBlock{
		section:           section,
		optional:          optional,
		multipleOptional:  multipleOptional,
		multipleMandatory: multipleMandatory,
	}

	return &out
}

// Section returns the section
func (obj *tokenBlock) Section() TokenSection {
	return obj.section
}

// HasOptional returns true if the tokenBlock is optional, false otherwise
func (obj *tokenBlock) HasOptional() bool {
	return obj.optional != nil
}

// Optional returns the optional RawToken, if any
func (obj *tokenBlock) Optional() RawToken {
	return obj.optional
}

// HasMultipleOptional returns true if the tokenBlock is multiple optional, false otherwise
func (obj *tokenBlock) HasMultipleOptional() bool {
	return obj.multipleOptional != nil
}

// MultipleOptional returns the multipleOptional RawToken, if any
func (obj *tokenBlock) MultipleOptional() RawToken {
	return obj.multipleOptional
}

// HasMultipleMandatory returns true if the tokenBlock is multiple mandatory, false otherwise
func (obj *tokenBlock) HasMultipleMandatory() bool {
	return obj.multipleMandatory != nil
}

// MultipleMandatory returns the multipleMandatory RawToken, if any
func (obj *tokenBlock) MultipleMandatory() RawToken {
	return obj.multipleMandatory
}
