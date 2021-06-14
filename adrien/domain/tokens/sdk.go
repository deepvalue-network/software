package tokens

import "github.com/deepvalue-network/software/adrien/domain/rules"

// NewAdapterBuilder creates a new adapter builder instance
func NewAdapterBuilder() AdapterBuilder {
	builder := NewBuilder()
	tokenBuilder := NewTokenBuilder()
	linesBuilder := NewLinesBuilder()
	lineBuilder := NewLineBuilder()
	blockBuilder := NewBlockBuilder()
	elementBuilder := NewElementBuilder()
	subElementsBuilder := NewSubElementsBuilder()
	subElementBuilder := NewSubElementBuilder()
	cardinalityBuilder := NewCardinalityBuilder()
	specificCardinalityBuilder := NewSpecificCardinalityBuilder()
	rangeBuilder := NewRangeBuilder()
	contentBuilder := NewContentBuilder()

	tokenPattern := "[a-z]{1}[a-zA-Z]*"
	anythingExcept := "[^%s]+"
	begin := ":"
	or := "|"
	end := ";"
	notDelimiter := "---"
	whiteSpacePattern := "[ \t\r\n]*"
	subElementPrefix := "->"
	subElementSuffix := "<-"
	cardinalityZeroMultiplePattern := "[\\*]{1}"
	cardinalityMultiplePattern := "[\\+]{1}"
	cardinalityRangeBegin := "{"
	cardinalityRangeEnd := "}"
	cardinalityRangeSeparator := ","

	return createAdapterBuilder(
		builder,
		tokenBuilder,
		linesBuilder,
		lineBuilder,
		blockBuilder,
		elementBuilder,
		subElementsBuilder,
		subElementBuilder,
		cardinalityBuilder,
		specificCardinalityBuilder,
		rangeBuilder,
		contentBuilder,
		tokenPattern,
		rules.RulePattern,
		anythingExcept,
		begin,
		or,
		end,
		notDelimiter,
		whiteSpacePattern,
		subElementPrefix,
		subElementSuffix,
		cardinalityZeroMultiplePattern,
		cardinalityMultiplePattern,
		cardinalityRangeBegin,
		cardinalityRangeEnd,
		cardinalityRangeSeparator,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a new token builder instance
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewLinesBuilder creates a new lines builder
func NewLinesBuilder() LinesBuilder {
	return createLinesBuilder()
}

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewBlockBuilder creates a new block builder
func NewBlockBuilder() BlockBuilder {
	return createBlockBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewSubElementsBuilder creates a new sub elements builder
func NewSubElementsBuilder() SubElementsBuilder {
	return createSubElementsBuilder()
}

// NewSubElementBuilder creates a new sub element builder
func NewSubElementBuilder() SubElementBuilder {
	return createSubElementBuilder()
}

// NewSpecificCardinalityBuilder creates a new specific cardinality builder instance
func NewSpecificCardinalityBuilder() SpecificCardinalityBuilder {
	return createSpecificCardinalityBuilder()
}

// NewCardinalityBuilder creates a new cardinality builder
func NewCardinalityBuilder() CardinalityBuilder {
	return createCardinalityBuilder()
}

// NewRangeBuilder creates a new range builder
func NewRangeBuilder() RangeBuilder {
	return createRangeBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithRules(rules rules.Rules) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents a token adapter
type Adapter interface {
	ToTokens(content string) (Tokens, error)
}

// Builder represents a tokens builder
type Builder interface {
	Create() Builder
	WithTokens(tokens []Token) Builder
	Now() (Tokens, error)
}

// Tokens represets tokens
type Tokens interface {
	All() []Token
	Find(name string) (Token, error)
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithBlock(block Block) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Block() Block
}

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMust(must Lines) BlockBuilder
	WithNot(not Lines) BlockBuilder
	Now() (Block, error)
}

// Block represents a token block
type Block interface {
	Must() Lines
	HasNot() bool
	Not() Lines
}

// LinesBuilder represents a lines builder
type LinesBuilder interface {
	Create() LinesBuilder
	WithLines(lines []Line) LinesBuilder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	All() []Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithElements(elements []Element) LineBuilder
	Now() (Line, error)
}

// Line represents a token line
type Line interface {
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithContent(content Content) ElementBuilder
	WithCode(code string) ElementBuilder
	WithSubElements(subElements SubElements) ElementBuilder
	WithCardinality(cardinality Cardinality) ElementBuilder
	Now() (Element, error)
}

// Element represents a token element
type Element interface {
	Content() Content
	Code() string
	HasSubElements() bool
	SubElements() SubElements
	HasCardinality() bool
	Cardinality() Cardinality
}

// SubElementsBuilder represents sub elements builder
type SubElementsBuilder interface {
	Create() SubElementsBuilder
	WithSubElements(subElements []SubElement) SubElementsBuilder
	Now() (SubElements, error)
}

// SubElements represents sub elements
type SubElements interface {
	All() []SubElement
	Find(name string) (SubElement, error)
}

// SubElementBuilder represents a sub element builder
type SubElementBuilder interface {
	Create() SubElementBuilder
	WithContent(content Content) SubElementBuilder
	WithCardinality(cardinality SpecificCardinality) SubElementBuilder
	Now() (SubElement, error)
}

// SubElement represents a sub element
type SubElement interface {
	Content() Content
	Cardinality() SpecificCardinality
}

// SpecificCardinalityBuilder represents a specific cardinality builder
type SpecificCardinalityBuilder interface {
	Create() SpecificCardinalityBuilder
	WithAmount(amount uint) SpecificCardinalityBuilder
	WithRange(rnge Range) SpecificCardinalityBuilder
	Now() (SpecificCardinality, error)
}

// SpecificCardinality represents a specific cardinaltiy
type SpecificCardinality interface {
	IsAmount() bool
	Amount() *uint
	IsRange() bool
	Range() Range
}

// CardinalityBuilder represents a cardinality builder
type CardinalityBuilder interface {
	Create() CardinalityBuilder
	IsNonZeroMultiple() CardinalityBuilder
	IsZeroMultiple() CardinalityBuilder
	WithSpecific(specific SpecificCardinality) CardinalityBuilder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality of an element
type Cardinality interface {
	IsNonZeroMultiple() bool
	IsZeroMultiple() bool
	IsSpecific() bool
	Specific() SpecificCardinality
}

// RangeBuilder represents a range builder
type RangeBuilder interface {
	Create() RangeBuilder
	WithMinimum(min uint) RangeBuilder
	WithMaximum(max uint) RangeBuilder
	Now() (Range, error)
}

// Range represents a range
type Range interface {
	Min() uint
	HasMax() bool
	Max() *uint
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithToken(token string) ContentBuilder
	WithRule(rule rules.Rule) ContentBuilder
	Now() (Content, error)
}

// Content represents an element content
type Content interface {
	Name() string
	IsToken() bool
	Token() string
	IsRule() bool
	Rule() rules.Rule
}
