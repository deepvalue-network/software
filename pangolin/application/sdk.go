package application

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/lexers"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

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
	Execute(linkedProg linkers.Program, input map[string]computable.Value) (stackframes.StackFrame, error)
	TestsAll(linkedProg linkers.Program) error
	TestByNames(linkedProg linkers.Program, names []string) error
	TestByName(linkedProg linkers.Program, name string) error
}
