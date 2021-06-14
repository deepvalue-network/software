package tokens

type subElement struct {
	content     Content
	cardinality SpecificCardinality
}

func createSubElement(
	content Content,
	cardinality SpecificCardinality,
) SubElement {
	out := subElement{
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Content returns the content
func (obj *subElement) Content() Content {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *subElement) Cardinality() SpecificCardinality {
	return obj.cardinality
}
