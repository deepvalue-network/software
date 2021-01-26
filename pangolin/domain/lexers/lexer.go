package lexers

import "github.com/steve-care-software/products/pangolin/domain/lexers/grammar"

type lexer struct {
	grammar  grammar.Grammar
	nodeTree NodeTree
}

func createLexer(grammar grammar.Grammar, nodeTree NodeTree) Lexer {
	out := lexer{
		grammar:  grammar,
		nodeTree: nodeTree,
	}

	return &out
}

// Grammar returns the grammar
func (obj *lexer) Grammar() grammar.Grammar {
	return obj.grammar
}

// Tree returns the nodeTree
func (obj *lexer) Tree() NodeTree {
	return obj.nodeTree
}
