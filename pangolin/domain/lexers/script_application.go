package lexers

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

type scriptApplication struct {
	grammar grammar.Grammar
	events  map[string]Event
}

type position struct {
	pattern string
	from    int
	to      int
	rule    grammar.Rule
}

func createPosition(pattern string, from int, to int, rule grammar.Rule) (*position, error) {
	if from < 0 {
		str := fmt.Sprintf("the from index must be greater or equal to zero when creating a position (Pattern: %s)", pattern)
		return nil, errors.New(str)
	}

	if to <= from {
		str := fmt.Sprintf("the to (index: %d) must be greater than the from (index: %d) when creating a position for the pattern: %s", from, to, pattern)
		return nil, errors.New(str)
	}

	out := position{
		pattern: pattern,
		from:    from,
		to:      to,
		rule:    rule,
	}

	return &out, nil
}

func createScriptApplication(grammar grammar.Grammar, events map[string]Event) ScriptApplication {
	out := scriptApplication{
		grammar: grammar,
		events:  events,
	}

	return &out
}

// Execute executes the script application
func (app *scriptApplication) Execute(script string) (string, error) {
	positions, err := app.findChannelPositions(script)
	if err != nil {
		return "", nil
	}

	index := 0
	runes := []rune(script)
	originalAmount := len(runes)
	for _, onePosition := range positions {
		if evt, ok := app.events[onePosition.pattern]; ok {
			fn := evt.Fn()

			from := onePosition.from - index
			to := onePosition.to - index
			runes = fn(from, to, runes, onePosition.rule)
			index = originalAmount - len(runes)
		}
	}

	return string(runes), nil
}

func (app *scriptApplication) findChannelPositions(script string) ([]*position, error) {
	positions := []*position{}
	if !app.grammar.HasChannels() {
		return positions, nil
	}

	channels := app.grammar.Channels().Tokens()
	for _, oneChannel := range channels {
		pos, err := app.token(oneChannel, script)
		if err != nil {
			return nil, err
		}

		for _, onePos := range pos {
			positions = append(positions, onePos)
		}
	}

	return positions, nil
}

func (app *scriptApplication) token(tok grammar.Token, script string) ([]*position, error) {
	out := []*position{}
	blcks := tok.Blocks()
	for _, oneBlocks := range blcks {
		positions, err := app.tokenBlocks(tok.Name(), oneBlocks, script)
		if err != nil {
			return nil, err
		}

		for _, onePosition := range positions {
			out = append(out, onePosition)
		}
	}

	return out, nil
}

func (app *scriptApplication) tokenBlocks(tokenName string, blocks grammar.TokenBlocks, script string) ([]*position, error) {
	out := []*position{}
	blcks := blocks.Get()
	for _, oneBlock := range blcks {
		positions, err := app.tokenBlock(tokenName, oneBlock, script)
		if err != nil {
			return nil, err
		}

		for _, onePosition := range positions {
			out = append(out, onePosition)
		}
	}

	return out, nil
}

func (app *scriptApplication) tokenBlock(tokenName string, block grammar.TokenBlock, script string) ([]*position, error) {
	section := block.Section()
	positions, err := app.section(tokenName, section, script)
	if err != nil {
		return nil, err
	}

	if block.HasOptional() {
		return []*position{
			positions[0],
		}, nil
	}

	return positions, nil
}

func (app *scriptApplication) section(tokenName string, section grammar.TokenSection, script string) ([]*position, error) {
	if section.HasRule() {
		rule := section.Rule()
		return app.tokenRule(tokenName, rule, script)
	}

	reference := section.Token().Reference()
	return app.findToken(reference, script)
}

func (app *scriptApplication) tokenRule(tokenName string, tokenRule grammar.TokenRule, script string) ([]*position, error) {
	patternAsString := ""
	rule := tokenRule.Rule()
	sections := rule.Sections()
	for _, oneSection := range sections {
		patternPartAsString := app.findRulePattern(oneSection)
		patternAsString = fmt.Sprintf("%s%s", patternAsString, patternPartAsString)
	}

	out := []*position{}
	pattern := regexp.MustCompile(patternAsString)
	indexes := pattern.FindAllStringIndex(script, -1)
	for _, oneIndex := range indexes {
		pos, err := createPosition(tokenName, oneIndex[0], oneIndex[1], rule)
		if err != nil {
			return nil, err
		}

		out = append(out, pos)
	}

	return out, nil
}

func (app *scriptApplication) findToken(reference string, script string) ([]*position, error) {
	allTokens := app.grammar.Tokens().Tokens()
	if tok, ok := allTokens[reference]; ok {
		return app.token(tok, script)
	}

	str := fmt.Sprintf("the Token (Reference: %s) could not be found", reference)
	return nil, errors.New(str)
}

func (app *scriptApplication) findRulePattern(section grammar.RuleSection) string {
	if section.HasConstant() {
		return section.Constant().Value()
	}

	return section.Pattern().Value()
}
