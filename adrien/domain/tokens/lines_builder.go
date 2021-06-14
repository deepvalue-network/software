package tokens

import "errors"

type linesBuilder struct {
	list []Line
}

func createLinesBuilder() LinesBuilder {
	out := linesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linesBuilder) Create() LinesBuilder {
	return createLinesBuilder()
}

// WithLines add lines to the builder
func (app *linesBuilder) WithLines(lines []Line) LinesBuilder {
	app.list = lines
	return app
}

// Now builds a new Lines instance
func (app *linesBuilder) Now() (Lines, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Lines instance")
	}

	return createLines(app.list), nil
}
