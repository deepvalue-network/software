package parsers

type application struct {
	head  HeadSection
	label LabelSection
	main  MainSection
	test  TestSection
}

func createApplication(
	head HeadSection,
	main MainSection,
) Application {
	return createApplicationInternally(head, main, nil, nil)
}

func createApplicationWithTest(
	head HeadSection,
	main MainSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, nil)
}

func createApplicationWithLabel(
	head HeadSection,
	main MainSection,
	label LabelSection,
) Application {
	return createApplicationInternally(head, main, nil, label)
}

func createApplicationWithLabelAndTest(
	head HeadSection,
	main MainSection,
	label LabelSection,
	test TestSection,
) Application {
	return createApplicationInternally(head, main, test, label)
}

func createApplicationInternally(
	head HeadSection,
	main MainSection,
	test TestSection,
	label LabelSection,
) Application {
	out := application{
		head:  head,
		main:  main,
		test:  test,
		label: label,
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

// HasTest returns true if there is tests, false otherwise
func (obj *application) HasTest() bool {
	return obj.test != nil
}

// Test returns the test section, if any
func (obj *application) Test() TestSection {
	return obj.test
}
