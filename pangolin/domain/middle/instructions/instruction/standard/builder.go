package standard

import "errors"

type builder struct {
	arythmeticBuilder ArythmeticBuilder
	logicalBuilder    LogicalBuilder
	relationalBuilder RelationalBuilder
	miscBuilder       MiscBuilder
	operationBuilder  OperationBuilder
	operation         Operation
	res               string
	first             string
	second            string
	isConcatenation   bool
	isFrameAssignment bool
	isAdd             bool
	isSub             bool
	isMul             bool
	isLessThan        bool
	isEqual           bool
	isNotEqual        bool
	isAnd             bool
	isOr              bool
}

func createBuilder(
	arythmeticBuilder ArythmeticBuilder,
	logicalBuilder LogicalBuilder,
	relationalBuilder RelationalBuilder,
	miscBuilder MiscBuilder,
	operationBuilder OperationBuilder,
) Builder {
	out := builder{
		arythmeticBuilder: arythmeticBuilder,
		logicalBuilder:    logicalBuilder,
		relationalBuilder: relationalBuilder,
		miscBuilder:       miscBuilder,
		operationBuilder:  operationBuilder,
		operation:         nil,
		res:               "",
		first:             "",
		second:            "",
		isConcatenation:   false,
		isFrameAssignment: false,
		isAdd:             false,
		isSub:             false,
		isMul:             false,
		isLessThan:        false,
		isEqual:           false,
		isNotEqual:        false,
		isAnd:             false,
		isOr:              false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.arythmeticBuilder,
		app.logicalBuilder,
		app.relationalBuilder,
		app.miscBuilder,
		app.operationBuilder,
	)
}

// WithResult adds a result to the builder
func (app *builder) WithResult(result string) Builder {
	app.res = result
	return app
}

// WithFirst adds a first to the builder
func (app *builder) WithFirst(first string) Builder {
	app.first = first
	return app
}

// WithSecond adds a second to the builder
func (app *builder) WithSecond(second string) Builder {
	app.second = second
	return app
}

// WithOperation adds an operation to the builder
func (app *builder) WithOperation(operation Operation) Builder {
	app.operation = operation
	return app
}

// IsConcatenation flags a builder as concatenation
func (app *builder) IsConcatenation() Builder {
	app.isConcatenation = true
	return app
}

// IsFrameAssignment flags a builder as frameAssignment
func (app *builder) IsFrameAssignment() Builder {
	app.isFrameAssignment = true
	return app
}

// IsAdd flags a builder as add
func (app *builder) IsAdd() Builder {
	app.isAdd = true
	return app
}

// IsSub flags a builder as sub
func (app *builder) IsSub() Builder {
	app.isSub = true
	return app
}

// IsMul flags a builder as mul
func (app *builder) IsMul() Builder {
	app.isMul = true
	return app
}

// IsLessThan flags a builder as lessThan
func (app *builder) IsLessThan() Builder {
	app.isLessThan = true
	return app
}

// IsEqual flags a builder as equal
func (app *builder) IsEqual() Builder {
	app.isEqual = true
	return app
}

// IsNotEqual flags a builder as notEqual
func (app *builder) IsNotEqual() Builder {
	app.isNotEqual = true
	return app
}

// IsAnd flags a builder as and
func (app *builder) IsAnd() Builder {
	app.isAnd = true
	return app
}

// IsOr flags a builder as or
func (app *builder) IsOr() Builder {
	app.isOr = true
	return app
}

// Now builds a new Standard instance
func (app *builder) Now() (Standard, error) {
	if app.res == "" {
		return nil, errors.New("the result is mandatory in order to build a Standard instance")
	}

	if app.first == "" {
		return nil, errors.New("the first is mandatory in order to build a Standard instance")
	}

	if app.second == "" {
		return nil, errors.New("the second is mandatory in order to build a Standard instance")
	}

	if app.operation != nil {
		return createStandard(app.operation, app.res, app.first, app.second), nil
	}

	operationBuilder := app.operationBuilder.Create()
	if app.isConcatenation || app.isFrameAssignment {
		builder := app.miscBuilder.Create()
		if app.isConcatenation {
			builder.IsConcatenation()
		}

		if app.isFrameAssignment {
			builder.IsFrameAssignment()
		}

		misc, err := builder.Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithMisc(misc)
	}

	if app.isAdd || app.isSub || app.isMul {
		arythmeticBuilder := app.arythmeticBuilder.Create()
		if app.isAdd {
			arythmeticBuilder.IsAdd()
		}

		if app.isSub {
			arythmeticBuilder.IsSub()
		}

		if app.isMul {
			arythmeticBuilder.IsMul()
		}

		ary, err := arythmeticBuilder.Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithArythmetic(ary)
	}

	if app.isLessThan || app.isEqual || app.isNotEqual {
		relationalBuilder := app.relationalBuilder.Create()
		if app.isLessThan {
			relationalBuilder.IsLessThan()
		}

		if app.isEqual {
			relationalBuilder.IsEqual()
		}

		if app.isNotEqual {
			relationalBuilder.IsNotEqual()
		}

		rel, err := relationalBuilder.Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithRelational(rel)
	}

	if app.isAnd || app.isOr {
		logicalBuilder := app.logicalBuilder.Create()
		if app.isAnd {
			logicalBuilder.IsAnd()
		}

		if app.isOr {
			logicalBuilder.IsOr()
		}

		log, err := logicalBuilder.Now()
		if err != nil {
			return nil, err
		}

		operationBuilder.WithLogical(log)
	}

	op, err := operationBuilder.Now()
	if err != nil {
		return nil, err
	}

	return createStandard(op, app.res, app.first, app.second), nil
}
