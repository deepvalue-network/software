package nodes

import uuid "github.com/satori/go.uuid"

type node struct {
	id          *uuid.UUID
	pos         Position
	orientation Orientation
	content     Content
	children    []Node
}

func createNode(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
) Node {
	return createNodeInternally(id, pos, orientation, nil, nil)
}

func createNodeWithContent(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
	content Content,
) Node {
	return createNodeInternally(id, pos, orientation, content, nil)
}

func createNodeWithChildren(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
	children []Node,
) Node {
	return createNodeInternally(id, pos, orientation, nil, children)
}

func createNodeWithContentAndChildren(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
	content Content,
	children []Node,
) Node {
	return createNodeInternally(id, pos, orientation, content, children)
}

func createNodeInternally(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
	content Content,
	children []Node,
) Node {
	out := node{
		id:          id,
		pos:         pos,
		orientation: orientation,
		content:     content,
		children:    children,
	}

	return &out
}

// ID returns the id
func (obj *node) ID() *uuid.UUID {
	return obj.id
}

// Position returns the position
func (obj *node) Position() Position {
	return obj.pos
}

// Orientation returns the orientation
func (obj *node) Orientation() Orientation {
	return obj.orientation
}

// HasContent returns true if there is content, false otherwise
func (obj *node) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *node) Content() Content {
	return obj.content
}

// HasChildren returns true if there is children, false otherwise
func (obj *node) HasChildren() bool {
	return obj.children != nil
}

// Children returns the children, if any
func (obj *node) Children() []Node {
	return obj.children
}
