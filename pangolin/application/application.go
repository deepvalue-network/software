package application

type application struct {
	lexer       Lexer
	parser      Parser
	linker      Linker
	interpreter Interpreter
}

func createApplication(
	lexer Lexer,
	parser Parser,
	linker Linker,
	interpreter Interpreter,
) Application {
	out := application{
		lexer:       lexer,
		parser:      parser,
		linker:      linker,
		interpreter: interpreter,
	}

	return &out
}

// Lexer returns the lexer application
func (obj *application) Lexer() Lexer {
	return obj.lexer
}

// Parser returns the parser application
func (obj *application) Parser() Parser {
	return obj.parser
}

// Linker returns the linker application
func (obj *application) Linker() Linker {
	return obj.linker
}

// Interpreter returns the interpreter application
func (obj *application) Interpreter() Interpreter {
	return obj.interpreter
}
