package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters"
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	lexerAdapterBuilder := lexers.NewAdapterBuilder()
	parserBuilder := parsers.NewParserBuilder()
	linkerBuilder := linkers.NewBuilder()
	interpreterBuilder := interpreters.NewBuilder()
	return createBuilder(
		lexerAdapterBuilder,
		parserBuilder,
		linkerBuilder,
		interpreterBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithCurrentDirPath(dirPath string) Builder
	WithGrammarFilePath(grammarFilePath string) Builder
	WithEvents(events []lexers.Event) Builder
	Now() (Application, error)
}

// Application represents a pangolin application
type Application interface {
	Lexer() Lexer
	Parser() Parser
	Linker() Linker
	Interpreter() Interpreter
}

// Lexer represents a pangolin lexer
type Lexer interface {
	Execute(script string) (lexers.Lexer, error)
}

// Parser represents a pangolin parser
type Parser interface {
	Execute(lexer lexers.Lexer) (parsers.Program, error)
}

// Linker represents a pangolin linker
type Linker interface {
	Execute(parsed parsers.Program) (linkers.Program, error)
}

// Interpreter represents a pangolin interpreter
type Interpreter interface {
	Execute(excutable linkers.Executable, input stackframes.StackFrame) (stackframes.StackFrame, error)
	Tests(testable linkers.Testable) error
}
