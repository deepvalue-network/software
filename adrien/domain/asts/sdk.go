package asts

import (
	"github.com/deepvalue-network/software/adrien/domain/grammars"
	"github.com/deepvalue-network/software/adrien/domain/tokens"
)

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithGrammar(grammar grammars.Grammar) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents an ast adapter
type Adapter interface {
	ToAST(script string) (AST, error)
}

// Builder represents an ast builder
type Builder interface {
	Create() Builder
	WithHead(head Node) Builder
	Now() (AST, error)
}

// AST represents an abstract syntax tree
type AST interface {
	Head() Node
	IsComplete() bool
}

// Node represents a node element
type Node interface {
	IsNode() bool
	Node() Node
	IsElement() bool
	Element() Element
	IsError() bool
	Error() Error
}

// Element represents a node element
type Element interface {
	Token() tokens.Token
	Code() string
	HasReplacement() bool
	Replacement() string
}

// Error represents an element with error
type Error interface {
	IsInvalid() bool
	Invalid() Invalid
	IsIncomplete() bool
	Incomplete() Incomplete
}

// Invalid represents an invalid element
type Invalid interface {
	Code() string
	Possibilities() tokens.Tokens
}

// Incomplete represents an incomplete element
type Incomplete interface {
	Possibilities() tokens.Tokens
}
