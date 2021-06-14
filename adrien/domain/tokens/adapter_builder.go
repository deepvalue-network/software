package tokens

import (
	"errors"

	"github.com/deepvalue-network/software/adrien/domain/rules"
)

type adapterBuilder struct {
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

func createAdapterBuilder(
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
) AdapterBuilder {
	out := adapterBuilder{
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
		rules:                          nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.builder,
		app.tokenBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.blockBuilder,
		app.elementBuilder,
		app.subElementsBuilder,
		app.subElementBuilder,
		app.cardinalityBuilder,
		app.specificCardinalityBuilder,
		app.rangeBuilder,
		app.contentBuilder,
		app.tokenPattern,
		app.rulePattern,
		app.anythingExcept,
		app.begin,
		app.or,
		app.end,
		app.notDelimiter,
		app.whiteSpacePattern,
		app.subElementPrefix,
		app.subElementSuffix,
		app.cardinalityZeroMultiplePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityRangeBegin,
		app.cardinalityRangeEnd,
		app.cardinalityRangeSeparator,
	)
}

// WithRules add rules to the builder
func (app *adapterBuilder) WithRules(rules rules.Rules) AdapterBuilder {
	app.rules = rules
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.rules == nil {
		return nil, errors.New("the rules are mandatory in order to build an Adapter instance")
	}

	return createAdapter(
		app.builder,
		app.tokenBuilder,
		app.linesBuilder,
		app.lineBuilder,
		app.blockBuilder,
		app.elementBuilder,
		app.subElementsBuilder,
		app.subElementBuilder,
		app.cardinalityBuilder,
		app.specificCardinalityBuilder,
		app.rangeBuilder,
		app.contentBuilder,
		app.tokenPattern,
		app.rulePattern,
		app.anythingExcept,
		app.begin,
		app.or,
		app.end,
		app.notDelimiter,
		app.whiteSpacePattern,
		app.subElementPrefix,
		app.subElementSuffix,
		app.cardinalityZeroMultiplePattern,
		app.cardinalityMultiplePattern,
		app.cardinalityRangeBegin,
		app.cardinalityRangeEnd,
		app.cardinalityRangeSeparator,
		app.rules,
	), nil
}
