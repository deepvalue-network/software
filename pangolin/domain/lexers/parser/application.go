package parsers

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/lexers"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the parser
func (app *application) Execute(parser Parser) (interface{}, error) {
	tree := parser.Lexer().Tree()
	events := parser.Events().Events()

	// call the enter events:
	ins, _, err := app.callEventsOnTree(tree, events, true)
	return ins, err
}

func (app *application) callEventsOnTree(
	tree lexers.NodeTree,
	events []Event,
	isEnter bool,
) (interface{}, []lexers.NodeTree, error) {

	// create the call stack:
	callStack := []lexers.NodeTree{}

	// execute the events:
	var last interface{}
	for _, oneEvent := range events {
		ins, err := app.callEvent(tree, oneEvent, isEnter, events)
		if err != nil {
			return nil, nil, err
		}

		if ins != nil {
			last = ins
		}
	}

	// add the current tree in the call stack:
	callStack = append(callStack, tree)

	// if the state is not enter, return:
	if !isEnter {
		return last, callStack, nil
	}

	nodes := tree.Nodes()
	for _, oneNode := range nodes {
		if oneNode.HasTrees() {
			trees := oneNode.Trees()
			for _, oneTree := range trees {
				ins, retCallStack, err := app.callEventsOnTree(oneTree, events, isEnter)
				if err != nil {
					return nil, nil, err
				}

				if ins != nil {
					last = ins
				}

				// add the returned call stack to the list:
				for _, oneCallStack := range retCallStack {
					callStack = append(callStack, oneCallStack)
				}
			}
		}
	}

	ins, err := app.callEventsOnStack(events, callStack)
	if err != nil {
		return nil, nil, err
	}

	return ins, []lexers.NodeTree{}, nil
}

func (app *application) callEventsOnStack(
	events []Event,
	callStack []lexers.NodeTree,
) (interface{}, error) {
	// reverse the call stack:
	amount := len(callStack)
	reversed := make([]lexers.NodeTree, amount)
	for index, oneTree := range callStack {
		newIndex := (amount - 1) - index
		reversed[newIndex] = oneTree
	}

	// call the stack:
	var last interface{}
	for _, oneTree := range reversed {
		ins, _, err := app.callEventsOnTree(oneTree, events, false)
		if err != nil {
			return nil, err
		}

		last = ins
	}

	return last, nil
}

// callEvent skips the event if the name does not match of the event does not contain the right func
func (app *application) callEvent(tree lexers.NodeTree, event Event, isEnter bool, allEvents []Event) (interface{}, error) {
	name := tree.Token().Name()
	tokenName := event.Token()
	if name != tokenName {
		return nil, nil
	}

	fn := event.Fn(isEnter)
	if fn == nil {
		return nil, nil
	}

	if event.HasRetrieveReplacement() {
		replFn := event.RetrieveReplacement()
		replacements, err := replFn(tree)
		if err != nil {
			return nil, err
		}

		for token, mp := range replacements {
			for code, ins := range mp {
				if ins == nil {
					continue
				}

				for _, oneEvent := range allEvents {
					if oneEvent.Token() != token {
						continue
					}

					if !oneEvent.HasSet() {
						str := fmt.Sprintf("while parsing the %s token (Grammar: %s), a replacement (Token: %s) was retrieved, however there is no Set func in this token's event", tree.Token(), tree.Grammar().Name(), token)
						return nil, errors.New(str)
					}

					setFn := oneEvent.Set()
					setFn(code, ins)
					continue
				}
			}
		}
	}

	ins, err := fn(tree)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
