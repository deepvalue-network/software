package grammar

import (
	"errors"
	"fmt"
)

type ruleTreeAdapter struct {
	repetitiveRuleNodeBuilder RepetitiveRuleNodeBuilder
	ruleNodeBuilder           RuleNodeBuilder
	ruleTreeBuilder           RuleTreeBuilder
}

func createRuleTreeAdapter(
	repetitiveRuleNodeBuilder RepetitiveRuleNodeBuilder,
	ruleNodeBuilder RuleNodeBuilder,
	ruleTreeBuilder RuleTreeBuilder,
) RuleTreeAdapter {
	out := ruleTreeAdapter{
		repetitiveRuleNodeBuilder: repetitiveRuleNodeBuilder,
		ruleNodeBuilder:           ruleNodeBuilder,
		ruleTreeBuilder:           ruleTreeBuilder,
	}

	return &out
}

// ToRuleTree converts a Grammar instance to a RuleTree
func (app *ruleTreeAdapter) ToRuleTree(grammar Grammar) (RuleTree, error) {
	tokens := grammar.Tokens()
	children := map[string]map[string]RuleTree{}
	if tokens.HasReplacements() {
		replacements := tokens.Replacements()
		for _, oneReplacement := range replacements {
			toGrammarName := oneReplacement.ToGrammar()
			fromTokenName := oneReplacement.FromToken()
			if _, ok := children[toGrammarName][fromTokenName]; ok {
				continue
			}

			tokens := grammar.Tokens().Tokens()
			if childTok, ok := tokens[fromTokenName]; ok {
				child, err := app.ruleTree(childTok, grammar, []string{}, nil)
				if err != nil {
					return nil, err
				}

				if _, ok := children[toGrammarName]; !ok {
					children[toGrammarName] = map[string]RuleTree{}
				}

				children[toGrammarName][fromTokenName] = child
			} else {
				str := fmt.Sprintf("the Token (%s) is used in a replacement instruction, but it is not declared", fromTokenName)
				return nil, errors.New(str)
			}
		}
	}

	tok := grammar.RootToken()
	return app.ruleTree(tok, grammar, []string{}, children)
}

func (app *ruleTreeAdapter) ruleTree(tok Token, grammar Grammar, callStack []string, children map[string]map[string]RuleTree) (RuleTree, error) {
	subBlocks := tok.Blocks()
	grammarName := grammar.Name()
	callName := fmt.Sprintf("%s.%s", grammarName, tok.Name())
	callStack = append(callStack, callName)
	repetitiveRuleNodes := [][]RepetitiveRuleNode{}
	for _, oneSubBlocks := range subBlocks {
		oneRepetitiveRuleNodes, err := app.tokenBlocks(oneSubBlocks, grammar, callStack, children)
		if err != nil {
			return nil, err
		}

		repetitiveRuleNodes = append(repetitiveRuleNodes, oneRepetitiveRuleNodes)
	}

	builder := app.ruleTreeBuilder.Create().WithToken(tok).WithNodes(repetitiveRuleNodes).WithGrammar(grammar)
	if chd, ok := children[grammarName]; ok {
		builder.WithChildren(chd)
	}

	tree, err := builder.Now()
	if err != nil {
		return nil, err
	}

	// returns
	return tree, nil
}

func (app *ruleTreeAdapter) tokenBlocks(blocks TokenBlocks, grammar Grammar, callStack []string, children map[string]map[string]RuleTree) ([]RepetitiveRuleNode, error) {
	nodes := []RepetitiveRuleNode{}
	tokenBlocks := blocks.Get()
	for _, oneTokenBlock := range tokenBlocks {
		node, err := app.tokenBlock(oneTokenBlock, grammar, callStack, children)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (app *ruleTreeAdapter) tokenBlock(tokenBlock TokenBlock, grammar Grammar, callStack []string, children map[string]map[string]RuleTree) (RepetitiveRuleNode, error) {
	section := tokenBlock.Section()
	node, err := app.section(section, grammar, callStack, children)
	if err != nil {
		return nil, err
	}

	builder := app.repetitiveRuleNodeBuilder.Create()
	if node != nil {
		builder.WithNode(node)
	}

	if node == nil {
		recursiveName := fmt.Sprintf("%s.%s", grammar.Name(), section.Token().Value())
		builder.WithRecursiveName(recursiveName)
	}

	if tokenBlock.HasMultipleMandatory() {
		builder.IsMultipleMandatory()
	}

	if tokenBlock.HasMultipleOptional() {
		builder.IsMultipleOptional()
	}

	if tokenBlock.HasOptional() {
		builder.IsOptional()
	}

	return builder.Now()
}

func (app *ruleTreeAdapter) section(section TokenSection, grammar Grammar, callStack []string, children map[string]map[string]RuleTree) (RuleNode, error) {
	ruleNodeBuilder := app.ruleNodeBuilder.Create()
	if section.HasRule() {
		rule := section.Rule().Rule()
		ruleNodeBuilder.WithRule(rule)
	}

	if section.HasToken() {
		tok := section.Token()
		leaf, err := app.ruleTreesByToken(tok, grammar, callStack, children)
		if err != nil {
			return nil, err
		}

		if leaf == nil {
			return nil, nil
		}

		ruleNodeBuilder.WithLeaf(leaf)
	}

	return ruleNodeBuilder.Now()
}

func (app *ruleTreeAdapter) ruleTreesByToken(tok RawToken, grammar Grammar, callStack []string, children map[string]map[string]RuleTree) (RuleTree, error) {
	gr, err := grammar.FetchByName(tok.Grammar())
	if err != nil {
		return nil, err
	}

	tokenName := tok.Value()
	callName := fmt.Sprintf("%s.%s", gr.Name(), tokenName)
	for _, oneCallStack := range callStack {
		if oneCallStack == callName {
			return nil, nil
		}
	}

	tokens := gr.Tokens()
	allTokens := tokens.Tokens()
	if tok, ok := allTokens[tokenName]; ok {
		return app.ruleTree(tok, gr, callStack, children)
	}

	str := fmt.Sprintf("the RawToken (reference: %s) could not be found", tok.Reference())
	return nil, errors.New(str)
}
