package parsers

type application struct {
	head      HeadSection
	label     LabelSection
	main      MainSection
	test      TestSection
	def DefinitionSection
}

func createApplication(
	head HeadSection,
	main MainSection,
) Application {
	return createApplicationInternally(head, main, nil, nil, nil)
}

func createApplicationWithTest(
	head HeadSection,
	main MainSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, nil, nil)
}

func createApplicationWithDefinition(
	head HeadSection,
	main MainSection,
	def DefinitionSection,
) Application {
	return createApplicationInternally(head, main, nil, def, nil)
}

func createApplicationWithLabel(
	head HeadSection,
	main MainSection,
	label LabelSection,
) Application {
	return createApplicationInternally(head, main, nil, nil, label)
}

func createApplicationWithDefinitionAndLabel(
	head HeadSection,
	main MainSection,
	def DefinitionSection,
	label LabelSection,
) Application {
	return createApplicationInternally(head, main, nil, def, label)
}

func createApplicationWithDefinitionAndTest(
	head HeadSection,
	main MainSection,
	def DefinitionSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, def, nil)
}

func createApplicationWithLabelAndTest(
	head HeadSection,
	main MainSection,
	label LabelSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, nil, label)
}

func createApplicationWithDefinitionAndLabelAndTest(
	head HeadSection,
	main MainSection,
	def DefinitionSection,
	label LabelSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, def, label)
}

func createApplicationInternally(
	head HeadSection,
	main MainSection,
	test TestSection,
	def DefinitionSection,
	label LabelSection,
) Application {
	out := application{
		head:      head,
		main:      main,
		test:      test,
		def: def,
		label:     label,
	}

	return &out
}

// Head returns the head section
func (obj *application) Head() HeadSection {
	return obj.head
}

// Main returns the main section
func (obj *application) Main() MainSection {
	return obj.main
}

// HasLabel returns true if there is a label, false otherwise
func (obj *application) HasLabel() bool {
	return obj.label != nil
}

// Label returns the label section
func (obj *application) Label() LabelSection {
	return obj.label
}

// HasDefinition returns true if there is definition, false otherwise
func (obj *application) HasDefinition() bool {
	return obj.def != nil
}

// Definition returns the definition, if any
func (obj *application) Definition() DefinitionSection {
	return obj.def
}


// HasTest returns true if there is tests, false otherwise
func (obj *application) HasTest() bool {
	return obj.test != nil
}

// Test returns the test section, if any
func (obj *application) Test() TestSection {
	return obj.test
}
