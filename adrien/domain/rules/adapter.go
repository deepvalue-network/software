package rules

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type adapter struct {
	builder                     Builder
	ruleBuilder                 RuleBuilder
	elementBuilder              ElementBuilder
	contentBuilder              ContentBuilder
	patternsBuilder             PatternsBuilder
	patternBuilder              PatternBuilder
	possibilityBuilder          PossibilityBuilder
	amountBuilder               AmountBuilder
	intervalBuilder             IntervalBuilder
	rulesPossibilitiesDelimiter string
	anythingExcept              string
	begin                       string
	end                         string
	possibilityAmountDelimiter  string
	possibilityDelimiter        string
	amountDelimiter             string
	rulePattern                 string
	constantDelimiter           string
}

func createAdapter(
	builder Builder,
	ruleBuilder RuleBuilder,
	elementBuilder ElementBuilder,
	contentBuilder ContentBuilder,
	patternsBuilder PatternsBuilder,
	patternBuilder PatternBuilder,
	possibilityBuilder PossibilityBuilder,
	amountBuilder AmountBuilder,
	intervalBuilder IntervalBuilder,
	rulesPossibilitiesDelimiter string,
	anythingExcept string,
	begin string,
	end string,
	possibilityAmountDelimiter string,
	possibilityDelimiter string,
	amountDelimiter string,
	rulePattern string,
	constantDelimiter string,
) Adapter {
	out := adapter{
		builder:                     builder,
		ruleBuilder:                 ruleBuilder,
		elementBuilder:              elementBuilder,
		contentBuilder:              contentBuilder,
		patternsBuilder:             patternsBuilder,
		patternBuilder:              patternBuilder,
		possibilityBuilder:          possibilityBuilder,
		amountBuilder:               amountBuilder,
		intervalBuilder:             intervalBuilder,
		rulesPossibilitiesDelimiter: rulesPossibilitiesDelimiter,
		anythingExcept:              anythingExcept,
		begin:                       begin,
		end:                         end,
		possibilityAmountDelimiter:  possibilityAmountDelimiter,
		possibilityDelimiter:        possibilityDelimiter,
		amountDelimiter:             amountDelimiter,
		rulePattern:                 rulePattern,
		constantDelimiter:           constantDelimiter,
	}

	return &out
}

// ToRules converts content to a Rules instance
func (app *adapter) ToRules(content string) (Rules, error) {
	sections := strings.Split(content, app.rulesPossibilitiesDelimiter)
	if len(sections) != 2 {
		str := fmt.Sprintf("the rules was expected to contain Rules and Patterns (2 sections), %d sections provided", len(sections))
		return nil, errors.New(str)
	}

	patterns, err := app.patterns(sections[1])
	if err != nil {
		return nil, err
	}

	rules, err := app.rules(sections[0], patterns)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithRules(rules).Now()
}

func (app *adapter) rules(content string, patterns Patterns) ([]Rule, error) {
	// find the rules:
	anythingExceptionEnd := fmt.Sprintf(app.anythingExcept, app.end)
	patternStr := fmt.Sprintf(
		"(%s)%s(%s)%s",
		app.rulePattern,
		app.begin,
		anythingExceptionEnd,
		app.end,
	)

	rules := []Rule{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		elements, err := app.elements(oneMatch[2], patterns)
		if err != nil {
			return nil, err
		}

		name := strings.TrimSpace(oneMatch[1])
		rule, err := app.ruleBuilder.Create().WithName(name).WithCode(oneMatch[0]).WithElements(elements).Now()
		if err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (app *adapter) elements(content string, patterns Patterns) ([]Element, error) {
	anythingExceptConstantDelimiter := fmt.Sprintf(app.anythingExcept, app.constantDelimiter)
	patternStr := fmt.Sprintf(
		"(%s)?(%s(%s)%s)?",
		anythingExceptConstantDelimiter,
		app.constantDelimiter,
		anythingExceptConstantDelimiter,
		app.constantDelimiter,
	)

	elements := []Element{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		patternName := strings.TrimSpace(oneMatch[1])
		if patternName != "" {
			pattern, err := patterns.Find(patternName)
			if err != nil {
				return nil, err
			}

			content, err := app.contentBuilder.Create().WithPattern(pattern).Now()
			if err != nil {
				return nil, err
			}

			element, err := app.elementBuilder.Create().WithCode(oneMatch[0]).WithContent(content).Now()
			if err != nil {
				return nil, err
			}

			elements = append(elements, element)
		}

		constant := strings.TrimSpace(oneMatch[3])
		if constant != "" {
			content, err := app.contentBuilder.Create().WithConstant(constant).Now()
			if err != nil {
				return nil, err
			}

			element, err := app.elementBuilder.Create().WithCode(oneMatch[0]).WithContent(content).Now()
			if err != nil {
				return nil, err
			}

			elements = append(elements, element)
		}
	}

	return elements, nil
}

func (app *adapter) patterns(content string) (Patterns, error) {
	anythingExceptBegin := fmt.Sprintf(app.anythingExcept, app.begin)
	anythingExceptAmountDelimiter := fmt.Sprintf(app.anythingExcept, app.possibilityAmountDelimiter)
	anythingExceptionEnd := fmt.Sprintf(app.anythingExcept, app.end)

	patternStr := fmt.Sprintf(
		"(%s)%s(%s)%s(%s)%s",
		anythingExceptBegin,
		app.begin,
		anythingExceptAmountDelimiter,
		app.possibilityAmountDelimiter,
		anythingExceptionEnd,
		app.end,
	)

	out := []Pattern{}
	pattern := regexp.MustCompile(patternStr)
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, oneMatch := range matches {
		amountsAsString := strings.Split(oneMatch[3], app.amountDelimiter)
		amountElements := len(amountsAsString)

		if amountElements > 2 {
			str := fmt.Sprintf("the amount (%s) was expecting only at most 2 elements: a minimum (int) followed by a maximum (optional int)", oneMatch[3])
			return nil, errors.New(str)
		}

		possibilities := strings.Split(oneMatch[2], app.possibilityDelimiter)
		possibilityBuilder := app.possibilityBuilder.Create().WithList(possibilities)

		min := -1
		if amountElements >= 1 {
			el, err := strconv.Atoi(strings.TrimSpace(amountsAsString[0]))
			if err != nil {
				return nil, err
			}

			min = el
		}

		amountBuilder := app.amountBuilder.Create()
		if amountElements == 1 {
			amountBuilder.WithExactly(min)
		}

		if amountElements >= 2 {
			intervalBuilder := app.intervalBuilder.Create().WithMin(min)
			amountStr := strings.TrimSpace(amountsAsString[1])
			if amountStr != "" {
				el, err := strconv.Atoi(amountStr)
				if err != nil {
					return nil, err
				}

				intervalBuilder.WithMax(el)
			}

			interval, err := intervalBuilder.Now()
			if err != nil {
				return nil, err
			}

			amountBuilder.WithInterval(interval)
		}

		amount, err := amountBuilder.Now()
		if err != nil {
			return nil, err
		}

		possibility, err := possibilityBuilder.WithAmount(amount).Now()
		if err != nil {
			return nil, err
		}

		patternString := strings.TrimSpace(oneMatch[1])
		pattern, err := app.patternBuilder.Create().WithCode(oneMatch[0]).WithPattern(patternString).WithPossibility(possibility).Now()
		if err != nil {
			return nil, err
		}

		out = append(out, pattern)
	}

	return app.patternsBuilder.Create().WithPatterns(out).Now()
}
