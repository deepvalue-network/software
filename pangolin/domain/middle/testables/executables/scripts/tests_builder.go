package scripts

import "errors"

type testsBuilder struct {
	list []Test
}

func createTestsBuilder() TestsBuilder {
	out := testsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *testsBuilder) Create() TestsBuilder {
	return createTestsBuilder()
}

// WithTests add tests to the builder
func (app *testsBuilder) WithTests(tests []Test) TestsBuilder {
	app.list = tests
	return app
}

// WithTests add tests to the builder
func (app *testsBuilder) Now() (Tests, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Test in order to build a Tests instance")
	}

	mp := map[string]Test{}
	for _, oneTest := range app.list {
		name := oneTest.Name()
		mp[name] = oneTest
	}

	return createTests(app.list, mp), nil
}
