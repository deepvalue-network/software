package lexers

import "errors"

type nodeBuilder struct {
	trees         []NodeTree
	elements      []Element
	recursiveName string
}

func createNodeBuilder() NodeBuilder {
	out := nodeBuilder{
		trees:         nil,
		elements:      nil,
		recursiveName: "",
	}

	return &out
}

// Create initializes the builder
func (app *nodeBuilder) Create() NodeBuilder {
	return createNodeBuilder()
}

// WithTrees add []NodeTree to the builder
func (app *nodeBuilder) WithTrees(trees []NodeTree) NodeBuilder {
	app.trees = trees
	return app
}

// WithElements add []Element to the builder
func (app *nodeBuilder) WithElements(elements []Element) NodeBuilder {
	app.elements = elements
	return app
}

// WithRecursiveName add a recursiveName to the builder
func (app *nodeBuilder) WithRecursiveName(recursiveName string) NodeBuilder {
	app.recursiveName = recursiveName
	return app
}

// Now builds a new Node instance
func (app *nodeBuilder) Now() (Node, error) {
	if app.trees != nil && len(app.trees) > 0 {
		return createNodeWithNodeTrees(app.trees)
	}

	if app.elements != nil && len(app.elements) > 0 {
		return createNodeWithElements(app.elements)
	}

	if app.recursiveName != "" {
		return createNodeWithRecursiveName(app.recursiveName)
	}

	return nil, errors.New("the Node instance is invalid")
}
