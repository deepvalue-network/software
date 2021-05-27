package token

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	codeMatchBuilder := NewCodeMatchBuilder()
	codeBuilder := NewCodeBuilder()
	return createAdapter(builder, codeMatchBuilder, codeBuilder)
}

// NewBuilder creates a new token builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewCodeMatchBuilder creates a new codeMatchBuilder instance
func NewCodeMatchBuilder() CodeMatchBuilder {
	return createCodeMatchBuilder()
}

// NewCodeBuilder creates a new code builder
func NewCodeBuilder() CodeBuilder {
	return createCodeBuilder()
}

// Adapter represents a token adapter
type Adapter interface {
	ToToken(parsed parsers.Token) (Token, error)
}

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithCodeMatch(codeMatch CodeMatch) Builder
	WithCode(code Code) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	IsCodeMatch() bool
	CodeMatch() CodeMatch
	IsCode() bool
	Code() Code
}

// CodeMatchBuilder represents a codeMatch builder
type CodeMatchBuilder interface {
	Create() CodeMatchBuilder
	WithReturn(ret string) CodeMatchBuilder
	WithSectionName(sectionName string) CodeMatchBuilder
	WithPatterns(patterns []string) CodeMatchBuilder
	Now() (CodeMatch, error)
}

// CodeMatch represents a token code match
type CodeMatch interface {
	Return() string
	SectionName() string
	Patterns() []string
}

// CodeBuilder represents a code builder
type CodeBuilder interface {
	Create() CodeBuilder
	WithReturn(ret string) CodeBuilder
	WithPattern(pattern string) CodeBuilder
	WithAmount(amount string) CodeBuilder
	Now() (Code, error)
}

// Code represents a token code
type Code interface {
	Return() string
	HasPattern() bool
	Pattern() string
	HasAmount() bool
	Amount() string
}
