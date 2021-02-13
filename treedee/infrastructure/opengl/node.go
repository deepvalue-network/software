package opengl

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type node struct {
	id          *uuid.UUID
	pos         Position
	orientation Orientation
	content     NodeContent
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
	content NodeContent,
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
	content NodeContent,
	children []Node,
) Node {
	return createNodeInternally(id, pos, orientation, content, children)
}

func createNodeInternally(
	id *uuid.UUID,
	pos Position,
	orientation Orientation,
	content NodeContent,
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
func (obj *node) Content() NodeContent {
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

// Render renders the node:
func (obj *node) Render(
	delta time.Duration,
	activeCamera WorldCamera,
	activeScene Scene,
) error {
	// render the model of the node, if any:
	if obj.HasContent() {
		content := obj.Content()
		if content.IsModel() {
			model := content.Model()
			err := model.Render(delta, activeCamera, activeScene)
			if err != nil {
				return err
			}
		}
	}

	// render the children nodes, if any:
	if obj.HasChildren() {
		children := obj.Children()
		for _, oneChildNode := range children {
			// add the relative position:
			pos := obj.Position().Add(oneChildNode.Position())

			// add the relative orientation:
			orientation := obj.Orientation().Add(oneChildNode.Orientation())

			// update the camera:
			childActiveCamera := activeCamera.Update(pos, orientation)

			// render the child node:
			err := oneChildNode.Render(delta, childActiveCamera, activeScene)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
