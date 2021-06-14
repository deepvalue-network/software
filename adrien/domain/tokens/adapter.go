package tokens

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/deepvalue-network/software/adrien/domain/rules"
)

type adapter struct {
	builder                        Builder
	tokenBuilder                   TokenBuilder
	linesBuilder                   LinesBuilder
	lineBuilder                    LineBuilder
	blockBuilder                   BlockBuilder
	elementBuilder                 ElementBuilder
	subElementsBuilder             SubElementsBuilder
	subElementBuilder              SubElementBuilder
	cardinalityBuilder             CardinalityBuilder
	specificCardinalityBuilder     SpecificCardinalityBuilder
	rangeBuilder                   RangeBuilder
	contentBuilder                 ContentBuilder
	tokenPattern                   string
	rulePattern                    string
	anythingExcept                 string
	begin                          string
	or                             string
	end                            string
	notDelimiter                   string
	whiteSpacePattern              string
	subElementPrefix               string
	subElementSuffix               string
	cardinalityZeroMultiplePattern string
	cardinalityMultiplePattern     string
	cardinalityRangeBegin          string
	cardinalityRangeEnd            string
	cardinalityRangeSeparator      string
	rules                          rules.Rules
}

func createAdapter(
	builder Builder,
	tokenBuilder TokenBuilder,
	linesBuilder LinesBuilder,
	lineBuilder LineBuilder,
	blockBuilder BlockBuilder,
	elementBuilder ElementBuilder,
	subElementsBuilder SubElementsBuilder,
	subElementBuilder SubElementBuilder,
	cardinalityBuilder CardinalityBuilder,
	specificCardinalityBuilder SpecificCardinalityBuilder,
	rangeBuilder RangeBuilder,
	contentBuilder ContentBuilder,
	tokenPattern string,
	rulePattern string,
	anythingExcept string,
	begin string,
	or string,
	end string,
	notDelimiter string,
	whiteSpacePattern string,
	subElementPrefix string,
	subElementSuffix string,
	cardinalityZeroMultiplePattern string,
	cardinalityMultiplePattern string,
	cardinalityRangeBegin string,
	cardinalityRangeEnd string,
	cardinalityRangeSeparator string,
	rules rules.Rules,
) Adapter {
	out := adapter{
		builder:                        builder,
		tokenBuilder:                   tokenBuilder,
		rulePattern:                    rulePattern,
		linesBuilder:                   linesBuilder,
		lineBuilder:                    lineBuilder,
		blockBuilder:                   blockBuilder,
		elementBuilder:                 elementBuilder,
		subElementsBuilder:             subElementsBuilder,
		subElementBuilder:              subElementBuilder,
		cardinalityBuilder:             cardinalityBuilder,
		specificCardinalityBuilder:     specificCardinalityBuilder,
		rangeBuilder:                   rangeBuilder,
		contentBuilder:                 contentBuilder,
		tokenPattern:                   tokenPattern,
		anythingExcept:                 anythingExcept,
		begin:                          begin,
		or:                             or,
		end:                            end,
		notDelimiter:                   notDelimiter,
		whiteSpacePattern:              whiteSpacePattern,
		subElementPrefix:               subElementPrefix,
		subElementSuffix:               subElementSuffix,
		cardinalityZeroMultiplePattern: cardinalityZeroMultiplePattern,
		cardinalityMultiplePattern:     cardinalityMultiplePattern,
		cardinalityRangeBegin:          cardinalityRangeBegin,
		cardinalityRangeEnd:            cardinalityRangeEnd,
		cardinalityRangeSeparator:      cardinalityRangeSeparator,
		rules:                          rules,
	}

	return &out
}

// ToTokens converts content to a Tokens instance
func (app *adapter) ToTokens(content string) (Tokens, error) {
	list, err := app.tokens(content)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithTokens(list).Now()
}

func (app *adapter) tokens(content string) ([]Token, error) {
	// find the tokens
	anythingExceptEnd := fmt.Sprintf(app.anythingExcept, app.end)
	patternStr := fmt.Sprintf(
		"%s(%s)%s%s%s(%s)%s",
		app.whiteSpacePattern,
		app.tokenPattern,
		app.whiteSpacePattern,
		app.begin,
		app.whiteSpacePattern,
		anythingExceptEnd,
		app.end,
	)

	tokens := []Token{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		name := oneMatch[1]
		block, err := app.block(name, oneMatch[2])
		if err != nil {
			return nil, err
		}

		token, err := app.tokenBuilder.Create().WithBlock(block).WithName(name).Now()
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (app *adapter) block(tokenName string, content string) (Block, error) {
	sections := strings.Split(content, app.notDelimiter)
	if len(sections) <= 0 {
		str := fmt.Sprintf("there must be at least must Lines in order to create the Block instance of Token: %s", tokenName)
		return nil, errors.New(str)
	}

	must, err := app.lines(sections[0])
	if err != nil {
		return nil, err
	}

	builder := app.blockBuilder.Create().WithMust(must)
	if len(sections) > 1 {
		trimmedContent := strings.TrimSpace(sections[1])
		if !strings.HasPrefix(trimmedContent, app.begin) {
			str := fmt.Sprintf("the not content (%s) of Token (%s) was expecting the begin (%s) character as prefix", sections[1], tokenName, app.begin)
			return nil, errors.New(str)
		}

		notLines, err := app.lines(sections[1][1:])
		if err != nil {
			return nil, err
		}

		builder.WithNot(notLines)
	}

	return builder.Now()
}

func (app *adapter) lines(content string) (Lines, error) {
	list := []Line{}
	sections := strings.Split(content, app.or)
	for _, oneSection := range sections {
		line, err := app.line(oneSection)
		if err != nil {
			return nil, err
		}

		list = append(list, line)
	}

	return app.linesBuilder.Create().WithLines(list).Now()
}

func (app *adapter) line(content string) (Line, error) {
	elements, err := app.elements(content)
	if err != nil {
		return nil, err
	}

	return app.lineBuilder.Create().WithElements(elements).Now()
}

func (app *adapter) elements(content string) ([]Element, error) {
	cardinalityRangePattern := app.cardinalityRangePattern()
	cardinalityPattern := fmt.Sprintf(
		"(%s)?(%s)?(%s)?",
		cardinalityRangePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityZeroMultiplePattern,
	)

	anythingExceptSubElementSuffix := fmt.Sprintf(app.anythingExcept, app.subElementSuffix)
	subElementPattern := fmt.Sprintf(
		"%s(%s)%s",
		app.subElementPrefix,
		anythingExceptSubElementSuffix,
		app.subElementSuffix,
	)

	elementSuffixPattern := fmt.Sprintf(
		"(%s)?(%s)?",
		cardinalityPattern,
		subElementPattern,
	)

	tokenPattern := fmt.Sprintf(
		"(%s)%s(%s)",
		app.tokenPattern,
		app.whiteSpacePattern,
		elementSuffixPattern,
	)

	rulePattern := fmt.Sprintf(
		"(%s)%s(%s)",
		app.rulePattern,
		app.whiteSpacePattern,
		elementSuffixPattern,
	)

	patternStr := fmt.Sprintf(
		"(%s)?(%s)?",
		tokenPattern,
		rulePattern,
	)

	out := []Element{}
	pattern := regexp.MustCompile(patternStr)
	trimmedContent := strings.TrimSpace(content)
	matches := pattern.FindAllStringSubmatch(trimmedContent, -1)
	for _, oneMatch := range matches {
		amount := (len(oneMatch) - 1) / 2
		index := amount + 1
		tokenMatches := oneMatch[1:index]
		ruleMatches := oneMatch[index:]

		if tokenMatches[0] != "" {
			tokenElement, err := app.tokenElement(tokenMatches)
			if err != nil {
				return nil, err
			}

			out = append(out, tokenElement)
			continue
		}

		if ruleMatches[0] != "" {
			ruleElement, err := app.ruleElement(ruleMatches)
			if err != nil {
				return nil, err
			}

			out = append(out, ruleElement)
		}
	}

	return out, nil
}

func (app *adapter) tokenElement(matches []string) (Element, error) {
	contentIns, err := app.contentBuilder.Create().WithToken(matches[1]).Now()
	if err != nil {
		return nil, err
	}

	return app.buildElement(contentIns, matches)
}

func (app *adapter) ruleElement(matches []string) (Element, error) {
	rule, err := app.rules.Find(matches[1])
	if err != nil {
		return nil, err
	}

	contentIns, err := app.contentBuilder.Create().WithRule(rule).Now()
	if err != nil {
		return nil, err
	}

	return app.buildElement(contentIns, matches)
}

func (app *adapter) buildElement(contentIns Content, matches []string) (Element, error) {
	elementBuilder := app.elementBuilder.Create().WithCode(matches[0]).WithContent(contentIns)
	subElements, err := app.subElements(matches[9])
	if err != nil {
		return nil, err
	}

	if subElements != nil {
		elementBuilder.WithSubElements(subElements)
	}

	cardinality, err := app.cardinality(matches[5], matches[6], matches[7])
	if err != nil {
		return nil, err
	}

	if cardinality != nil {
		elementBuilder.WithCardinality(cardinality)
	}

	return elementBuilder.Now()
}

func (app *adapter) subElements(content string) (SubElements, error) {
	trimmedContent := strings.TrimSpace(content)
	if trimmedContent == "" {
		return nil, nil
	}

	patternStr := fmt.Sprintf(
		"(%s)?(%s)?%s(%s)",
		app.tokenPattern,
		app.rulePattern,
		app.whiteSpacePattern,
		app.cardinalityRangePattern(),
	)

	list := []SubElement{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(trimmedContent, -1)
	for _, oneMatch := range matches {
		builder := app.subElementBuilder.Create()
		if oneMatch[1] != "" {
			contentIns, err := app.contentBuilder.Create().WithToken(oneMatch[1]).Now()
			if err != nil {
				return nil, err
			}

			builder.WithContent(contentIns)
		}

		if oneMatch[2] != "" {
			rule, err := app.rules.Find(oneMatch[2])
			if err != nil {
				return nil, err
			}

			contentIns, err := app.contentBuilder.Create().WithRule(rule).Now()
			if err != nil {
				return nil, err
			}

			builder.WithContent(contentIns)
		}

		specificCardinality, err := app.specificCardinality(oneMatch[4])
		if err != nil {
			return nil, err
		}

		if specificCardinality != nil {
			builder.WithCardinality(specificCardinality)
		}

		subElement, err := builder.Now()
		if err != nil {
			return nil, err
		}

		list = append(list, subElement)
	}

	return app.subElementsBuilder.Create().WithSubElements(list).Now()
}

func (app *adapter) cardinality(rangeContent string, multipleContent string, zeroMultipleContent string) (Cardinality, error) {
	trimmedRangeContent := strings.TrimSpace(rangeContent)
	trimmedMultiple := strings.TrimSpace(multipleContent)
	trimmedZeroMultiple := strings.TrimSpace(zeroMultipleContent)
	if trimmedRangeContent == "" && trimmedMultiple == "" && trimmedZeroMultiple == "" {
		return nil, nil
	}

	builder := app.cardinalityBuilder.Create()
	specificCardinality, err := app.specificCardinality(rangeContent)
	if err != nil {
		return nil, err
	}

	if specificCardinality != nil {
		builder.WithSpecific(specificCardinality)
	}

	if trimmedMultiple != "" {
		builder.IsNonZeroMultiple()
	}

	if trimmedZeroMultiple != "" {
		builder.IsZeroMultiple()
	}

	return builder.Now()
}

func (app *adapter) specificCardinality(content string) (SpecificCardinality, error) {
	trimmedContent := strings.TrimSpace(content)
	if trimmedContent == "" {
		return nil, nil
	}

	builder := app.specificCardinalityBuilder.Create()
	sections := strings.Split(content, app.cardinalityRangeSeparator)
	if len(sections) == 1 {
		amount, err := strconv.Atoi(sections[0])
		if err != nil {
			return nil, err
		}

		builder.WithAmount(uint(amount))
	}

	if len(sections) == 2 {
		min, err := strconv.Atoi(sections[0])
		if err != nil {
			return nil, err
		}

		rangeBuilder := app.rangeBuilder.Create().WithMinimum(uint(min))
		trimmedSection := strings.TrimSpace(sections[1])
		if trimmedSection != "" {
			max, err := strconv.Atoi(trimmedSection)
			if err != nil {
				return nil, err
			}

			rangeBuilder.WithMaximum(uint(max))
		}

		rnge, err := rangeBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithRange(rnge)
	}

	return builder.Now()
}

func (app *adapter) cardinalityRangePattern() string {
	anythingExceptCardinalityRangeEnd := fmt.Sprintf(app.anythingExcept, app.cardinalityRangeEnd)
	return fmt.Sprintf(
		"%s(%s)%s",
		app.cardinalityRangeBegin,
		anythingExceptCardinalityRangeEnd,
		app.cardinalityRangeEnd,
	)
}
