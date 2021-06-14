package lexers

import (
	"errors"
	"strings"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

type builder struct {
	scriptApplicationBuilder ScriptApplicationBuilder
	nodeTreeAdapterBuilder   NodeTreeAdapterBuilder
	grammar                  grammar.Grammar
	tree                     NodeTree
	script                   string
	evts                     []Event
}

func createBuilder(
	scriptApplicationBuilder ScriptApplicationBuilder,
	nodeTreeAdapterBuilder NodeTreeAdapterBuilder,
) Builder {
	out := builder{
		scriptApplicationBuilder: scriptApplicationBuilder,
		nodeTreeAdapterBuilder:   nodeTreeAdapterBuilder,
		grammar:                  nil,
		tree:                     nil,
		script:                   "",
		evts:                     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.scriptApplicationBuilder, app.nodeTreeAdapterBuilder)
}

// WithGrammar add a grammar instance to the builder
func (app *builder) WithGrammar(grammar grammar.Grammar) Builder {
	app.grammar = grammar
	return app
}

// WithNodeTree add a nodeTree instance to the builder
func (app *builder) WithNodeTree(tree NodeTree) Builder {
	app.tree = tree
	return app
}

// WithScript add a script instance to the builder
func (app *builder) WithScript(script string) Builder {
	app.script = script
	return app
}

// WithEvents add events to the builder
func (app *builder) WithEvents(evts []Event) Builder {
	app.evts = evts
	return app
}

// Now builds a new Lexer instance
func (app *builder) Now() (Lexer, error) {
	if app.grammar == nil {
		return nil, errors.New("the Grammar instance is mandatory in order to build a Lexer instance")
	}

	if app.script != "" {
		// replace the semicolon in the script by its ascii characters:
		app.script = strings.ReplaceAll(app.script, secmiColon, semiColonASCII)

		// build the script application:
		scriptApplication, err := app.scriptApplicationBuilder.Create().WithGrammar(app.grammar).WithEvents(app.evts).Now()
		if err != nil {
			return nil, err
		}

		// execute the script events:
		script, err := scriptApplication.Execute(app.script)
		if err != nil {
			return nil, err
		}

		// set the modified script as the current script:
		app.script = script

		// build the nodeTreeAdapter:
		nodeTreeAdapter, err := app.nodeTreeAdapterBuilder.Create().WithGrammar(app.grammar).Now()
		if err != nil {
			return nil, err
		}

		// build the nodeTree:
		nodeTree, err := nodeTreeAdapter.ToNodeTree(app.script)
		if err != nil {
			return nil, err
		}

		app.tree = nodeTree
	}

	if app.tree == nil {
		return nil, errors.New("the NodeTree is mandatory in order to build a Lexer instance")
	}

	return createLexer(app.grammar, app.tree), nil
}
