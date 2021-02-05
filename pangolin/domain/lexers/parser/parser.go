package parsers

import "github.com/deepvalue-network/software/pangolin/domain/lexers"

type parser struct {
	lexer  lexers.Lexer
	events Events
}

func createParser(lexer lexers.Lexer, events Events) Parser {
	out := parser{
		lexer:  lexer,
		events: events,
	}

	return &out
}

// Lexer returns the lexer
func (obj *parser) Lexer() lexers.Lexer {
	return obj.lexer
}

// Events returns the events
func (obj *parser) Events() Events {
	return obj.events
}
