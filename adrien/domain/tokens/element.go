package tokens

type element struct {
	content     Content
	code        string
	subElements SubElements
	cardinality Cardinality
}

func createElement(
	content Content,
	code string,
) Element {
	return createElementInternally(content, code, nil, nil)
}

func createElementWithSubElements(
	content Content,
	code string,
	subElements SubElements,
) Element {
	return createElementInternally(content, code, subElements, nil)
}

func createElementWithCardinality(
	content Content,
	code string,
	cardinality Cardinality,
) Element {
	return createElementInternally(content, code, nil, cardinality)
}

func createElementWithSubElementsAndCardinality(
	content Content,
	code string,
	subElements SubElements,
	cardinality Cardinality,
) Element {
	return createElementInternally(content, code, subElements, cardinality)
}

func createElementInternally(
	content Content,
	code string,
	subElements SubElements,
	cardinality Cardinality,
) Element {
	out := element{
		content:     content,
		code:        code,
		subElements: subElements,
		cardinality: cardinality,
	}

	return &out
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}

// Code returns the code
func (obj *element) Code() string {
	return obj.code
}

// HasSubElements returns true if there is subElements, false otherwise
func (obj *element) HasSubElements() bool {
	return obj.subElements != nil
}

// SubElements returns the subElements, if any
func (obj *element) SubElements() SubElements {
	return obj.subElements
}

// HasCardinality returns true if there is cardinality, false otherwise
func (obj *element) HasCardinality() bool {
	return obj.cardinality != nil
}

// Cardinality returns the cardinality, if any
func (obj *element) Cardinality() Cardinality {
	return obj.cardinality
}
