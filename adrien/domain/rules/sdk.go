package rules

import "regexp"

// RulePattern represents the rule pattern
const RulePattern = "[A-Z\\_]+"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	ruleBuilder := NewRuleBuilder()
	elementBuilder := NewElementBuilder()
	contentBuilder := NewContentBuilder()
	patternsBuilder := NewPatternsBuilder()
	patternBuilder := NewPatternBuilder()
	possibilityBuilder := NewPossibilityBuilder()
	amountBuilder := NewAmountBuilder()
	intervalBuilder := NewIntervalBuilder()
	rulesPossibilitiesDelimiter := "----"
	anythingExcept := "[^%s]+"
	begin := ":"
	end := ";"
	possibilityAmountDelimiter := "->"
	possibilityDelimiter := ","
	amountDelimiter := ","
	constantDelimiter := "'"
	return createAdapter(
		builder,
		ruleBuilder,
		elementBuilder,
		contentBuilder,
		patternsBuilder,
		patternBuilder,
		possibilityBuilder,
		amountBuilder,
		intervalBuilder,
		rulesPossibilitiesDelimiter,
		anythingExcept,
		begin,
		end,
		possibilityAmountDelimiter,
		possibilityDelimiter,
		amountDelimiter,
		RulePattern,
		constantDelimiter,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewRuleBuilder creates a new rule builder instance
func NewRuleBuilder() RuleBuilder {
	return createRuleBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	return createContentBuilder()
}

// NewPatternsBuilder creates a new patterns builder
func NewPatternsBuilder() PatternsBuilder {
	return createPatternsBuilder()
}

// NewPatternBuilder creates a new pattern builder instance
func NewPatternBuilder() PatternBuilder {
	return createPatternBuilder()
}

// NewPossibilityBuilder creates a new possibility builder instance
func NewPossibilityBuilder() PossibilityBuilder {
	return createPossibilityBuilder()
}

// NewAmountBuilder creates a new amount builder
func NewAmountBuilder() AmountBuilder {
	return createAmountBuilder()
}

// NewIntervalBuilder creates a new interval builder instance
func NewIntervalBuilder() IntervalBuilder {
	return createIntervalBuilder()
}

// Adapter represents a rule adapter
type Adapter interface {
	ToRules(content string) (Rules, error)
}

// Builder represents a rules builder
type Builder interface {
	Create() Builder
	WithRules(rules []Rule) Builder
	Now() (Rules, error)
}

// Rules represents rules
type Rules interface {
	All() []Rule
	Find(name string) (Rule, error)
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithName(name string) RuleBuilder
	WithCode(code string) RuleBuilder
	WithElements(elements []Element) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	Name() string
	Code() string
	Elements() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithContent(content Content) ElementBuilder
	WithCode(code string) ElementBuilder
	Now() (Element, error)
}

// Element represents a rule element
type Element interface {
	Content() Content
	Code() string
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithConstant(constant string) ContentBuilder
	WithPattern(pattern Pattern) ContentBuilder
	Now() (Content, error)
}

// Content represents an element content
type Content interface {
	IsConstant() bool
	Constant() string
	IsPattern() bool
	Pattern() Pattern
}

// PatternsBuilder represents a patterns builder
type PatternsBuilder interface {
	Create() PatternsBuilder
	WithPatterns(patterns []Pattern) PatternsBuilder
	Now() (Patterns, error)
}

// Patterns represents a patterns
type Patterns interface {
	All() []Pattern
	Find(name string) (Pattern, error)
}

// PatternBuilder represents a pattern builder
type PatternBuilder interface {
	Create() PatternBuilder
	WithCode(code string) PatternBuilder
	WithPattern(pattern string) PatternBuilder
	WithPossibility(possibility Possibility) PatternBuilder
	Now() (Pattern, error)
}

// Pattern represents a rule pattern
type Pattern interface {
	Name() string
	Code() string
	Pattern() *regexp.Regexp
	Possibility() Possibility
}

// PossibilityBuilder represents a possibility builder
type PossibilityBuilder interface {
	Create() PossibilityBuilder
	WithList(list []string) PossibilityBuilder
	WithAmount(amount Amount) PossibilityBuilder
	Now() (Possibility, error)
}

// Possibility represents a possibility
type Possibility interface {
	List() []string
	Amount() Amount
}

// AmountBuilder represents an amount builder
type AmountBuilder interface {
	Create() AmountBuilder
	WithExactly(exactly int) AmountBuilder
	WithInterval(interval Interval) AmountBuilder
	Now() (Amount, error)
}

// Amount represents a possibility amount
type Amount interface {
	IsExactly() bool
	Exactly() int
	IsInterval() bool
	Interval() Interval
}

// IntervalBuilder represents an interval builder
type IntervalBuilder interface {
	Create() IntervalBuilder
	WithMin(min int) IntervalBuilder
	WithMax(max int) IntervalBuilder
	Now() (Interval, error)
}

// Interval represents an interval
type Interval interface {
	Min() int
	HasMax() bool
	Max() int
}
