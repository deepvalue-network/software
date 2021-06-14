package tokens

import "errors"

type subElementsBuilder struct {
	list []SubElement
}

func createSubElementsBuilder() SubElementsBuilder {
	out := subElementsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *subElementsBuilder) Create() SubElementsBuilder {
	return createSubElementsBuilder()
}

// WithSubElements add subElements to the builder
func (app *subElementsBuilder) WithSubElements(subElements []SubElement) SubElementsBuilder {
	app.list = subElements
	return app
}

// Now builds a new SubElement instance
func (app *subElementsBuilder) Now() (SubElements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 SubElemnt in order to build a SubElements instance")
	}

	mp := map[string]SubElement{}
	for _, oneSubElement := range app.list {
		name := oneSubElement.Content().Name()
		mp[name] = oneSubElement
	}

	return createSubElements(app.list, mp), nil
}
