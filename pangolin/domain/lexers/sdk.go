package lexers

import (
	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

const defaultLexerName = "default"

const (
	patternInParenthesis = "(%s)"
	pattern              = "%s"
	semiColonASCII       = "&#59"
	secmiColon           = ";"
)

// EventFn represents the event func
type EventFn func(from int, to int, script []rune, rule grammar.Rule) []rune

// NewEventBuilder creates a new EventBuilder instance
func NewEventBuilder() EventBuilder {
	return createEventBuilder()
}

// NewAdapterBuilder creates a new lexer adapter builder
func NewAdapterBuilder() AdapterBuilder {
	builder := NewBuilder()
	grammarBuilder := grammar.NewBuilder()
	grammarRepositoryBuilder := grammar.NewRepositoryBuilder()
	return createAdapterBuilder(builder, grammarBuilder, grammarRepositoryBuilder, defaultLexerName)
}

// NewBuilder creates a new lexer builder
func NewBuilder() Builder {
	ruleTreeAdapter := grammar.NewRuleTreeAdapter()
	scriptApplicationBuilder := createScriptApplicationBuilder()
	elementBuilder := createElementBuilder()
	nodeBuilder := createNodeBuilder()
	nodeTreeBuilder := createNodeTreeBuilder(elementBuilder, nodeBuilder)
	nodeTreeAdapterBuilder := createNodeTreeAdapterBuilder(ruleTreeAdapter, nodeTreeBuilder)
	return createBuilder(scriptApplicationBuilder, nodeTreeAdapterBuilder)
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithName(name string) AdapterBuilder
	WithRoot(root string) AdapterBuilder
	WithFileFetcher(fileFetcher grammar.FileFetcher) AdapterBuilder
	WithGrammarFilePath(grammarFilePath string) AdapterBuilder
	WithGrammarRetrieverCriteria(grammarRetrieverCriteria grammar.RetrieverCriteria) AdapterBuilder
	WithGrammar(grammar grammar.Grammar) AdapterBuilder
	WithEvents(events []Event) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents the lexer adapter
type Adapter interface {
	ToLexer(script string) (Lexer, error)
}

// Builder represents a lexer builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammar.Grammar) Builder
	WithNodeTree(tree NodeTree) Builder
	WithScript(script string) Builder
	WithEvents(evts []Event) Builder
	Now() (Lexer, error)
}

// Lexer represents a lexer
type Lexer interface {
	Grammar() grammar.Grammar
	Tree() NodeTree
}

// NodeTreeAdapterBuilder represents a nodeTreeAdapter builder
type NodeTreeAdapterBuilder interface {
	Create() NodeTreeAdapterBuilder
	WithGrammar(grammar grammar.Grammar) NodeTreeAdapterBuilder
	Now() (NodeTreeAdapter, error)
}

// NodeTreeAdapter represents a nodeTree adapter
type NodeTreeAdapter interface {
	ToNodeTree(script string) (NodeTree, error)
}

// NodeTreeBuilder represents a nodeTree builder
type NodeTreeBuilder interface {
	Create() NodeTreeBuilder
	WithRuleTree(tree grammar.RuleTree) NodeTreeBuilder
	WithScript(script string) NodeTreeBuilder
	Now() (NodeTree, error)
}

// NodeTree represents a nodeTree
type NodeTree interface {
	Grammar() grammar.Grammar
	Token() grammar.Token
	Nodes() []Node
	HasSubNodeTrees() bool
	Code() string
	CodeFromName(name string) string
	CodesFromName(name string) []string
	BestMatchFromNames(names []string) (string, string)
	NextNodeTree() NodeTree
	NextNodeTrees() []NodeTree
	SubTreeFromName(name string) NodeTree
}

// NodeBuilder represents a node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithTrees(trees []NodeTree) NodeBuilder
	WithElements(elements []Element) NodeBuilder
	WithRecursiveName(recursiveName string) NodeBuilder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	HasRecursiveName() bool
	RecursiveName() string
	HasTrees() bool
	Trees() []NodeTree
	HasElements() bool
	Elements() []Element
	Code() string
	CodeFromName(name string) string
	CodesFromName(name string) []string
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule grammar.Rule) ElementBuilder
	WithCode(code string) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Rule() grammar.Rule
	Code() string
}

// ScriptApplicationBuilder represents a script application builder
type ScriptApplicationBuilder interface {
	Create() ScriptApplicationBuilder
	WithGrammar(grammar grammar.Grammar) ScriptApplicationBuilder
	WithEvents(evts []Event) ScriptApplicationBuilder
	Now() (ScriptApplication, error)
}

// ScriptApplication represents a script application
type ScriptApplication interface {
	Execute(script string) (string, error)
}

// EventBuilder represents an event builder
type EventBuilder interface {
	Create() EventBuilder
	WithToken(token string) EventBuilder
	WithFn(fn EventFn) EventBuilder
	Now() (Event, error)
}

// Event represents a lexer event
type Event interface {
	Token() string
	Fn() EventFn
}
