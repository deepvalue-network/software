package parsers

import "errors"

type applicationBuilder struct {
	head      HeadSection
	label     LabelSection
	main      MainSection
	test      TestSection
	def DefinitionSection
}

func createApplicationBuilder() ApplicationBuilder {
	out := applicationBuilder{
		head:      nil,
		label:     nil,
		main:      nil,
		test:      nil,
		def: nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder()
}

// WithHead adds an headSection to the builder
func (app *applicationBuilder) WithHead(head HeadSection) ApplicationBuilder {
	app.head = head
	return app
}

// WithLabel adds a labelSection to the builder
func (app *applicationBuilder) WithLabel(label LabelSection) ApplicationBuilder {
	app.label = label
	return app
}

// WithMain adds a mainSection to the builder
func (app *applicationBuilder) WithMain(main MainSection) ApplicationBuilder {
	app.main = main
	return app
}

// WithTest adds a testSection to the builder
func (app *applicationBuilder) WithTest(test TestSection) ApplicationBuilder {
	app.test = test
	return app
}

// WithVariables add definition to the builder
func (app *applicationBuilder) WithDefinition(def DefinitionSection) ApplicationBuilder {
	app.def = def
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.head == nil {
		return nil, errors.New("the HeadSection is mandatory in order to build a Application instance")
	}

	if app.main == nil {
		return nil, errors.New("the MainSection is mandatory in order to build a Application instance")
	}

	if app.test != nil && app.def != nil && app.label != nil {
		return createApplicationWithDefinitionAndLabelAndTest(app.head, app.main, app.def, app.label, app.test), nil
	}

	if  app.def != nil && app.label != nil {
		return createApplicationWithDefinitionAndLabel(app.head, app.main, app.def, app.label), nil
	}

	if app.test != nil && app.label != nil {
		return createApplicationWithLabelAndTest(app.head, app.main, app.label, app.test), nil
	}

	if app.test != nil && app.def != nil {
		return createApplicationWithDefinitionAndTest(app.head, app.main, app.def, app.test), nil
	}

	if app.test != nil {
		return createApplicationWithTest(app.head, app.main, app.test), nil
	}

	if app.def != nil {
		return createApplicationWithDefinition(app.head, app.main, app.def), nil
	}

	if app.label != nil {
		return createApplicationWithLabel(app.head, app.main, app.label), nil
	}

	return createApplication(app.head, app.main), nil
}
