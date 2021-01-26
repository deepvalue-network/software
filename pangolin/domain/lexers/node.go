package lexers

import (
	"errors"
	"strings"
)

type node struct {
	trees         []NodeTree
	elements      []Element
	recursiveName string
}

func createNodeWithNodeTrees(trees []NodeTree) (Node, error) {
	return createNodeInternally(trees, nil, "")
}

func createNodeWithElements(elements []Element) (Node, error) {
	return createNodeInternally(nil, elements, "")
}

func createNodeWithRecursiveName(recursiveName string) (Node, error) {
	return createNodeInternally(nil, nil, recursiveName)
}

func createNodeInternally(trees []NodeTree, elements []Element, recursiveName string) (Node, error) {

	if elements != nil {
		for _, oneElement := range elements {
			if oneElement == nil {
				return nil, errors.New("the Node contains at least 1 Element that is invalid (nil)")
			}
		}
	}

	if trees != nil {
		for _, oneTree := range trees {
			if oneTree == nil {
				return nil, errors.New("the Node contains at least 1 NodeTree that is invalid (nil)")
			}
		}
	}

	out := node{
		trees:         trees,
		elements:      elements,
		recursiveName: recursiveName,
	}

	return &out, nil
}

// HasRecursiveName returns true if there is a recursive name, false otherwise
func (obj *node) HasRecursiveName() bool {
	return obj.recursiveName != ""
}

// RecursiveName returns the recursiveName, if any
func (obj *node) RecursiveName() string {
	return obj.recursiveName
}

// HasTrees return true if there is []NodeTree, false otherwise
func (obj *node) HasTrees() bool {
	return obj.trees != nil
}

// Trees returns the []NodeTree, if any
func (obj *node) Trees() []NodeTree {
	return obj.trees
}

// HasElements return true if there is []Element, false otherwise
func (obj *node) HasElements() bool {
	return obj.elements != nil
}

// Elements return the []Element, if any
func (obj *node) Elements() []Element {
	return obj.elements
}

// Code returns the code that the Node matches
func (obj *node) Code() string {
	if obj.HasTrees() {
		out := []string{}
		for _, oneTree := range obj.trees {
			out = append(out, oneTree.Code())
		}

		return strings.Join(out, "")
	}

	if obj.HasRecursiveName() {
		return ""
	}

	out := []string{}
	for _, oneElement := range obj.elements {
		out = append(out, oneElement.Code())
	}

	return strings.Join(out, "")
}

// CodeFromName returns the code from name, the returned code is empty if the name does not exists
func (obj *node) CodeFromName(name string) string {
	codes := obj.CodesFromName(name)
	if len(codes) <= 0 {
		return ""
	}

	return codes[0]
}

// CodesFromName returns the codes from name, the returned code is empty if the name does not exists
func (obj *node) CodesFromName(name string) []string {
	if obj.HasRecursiveName() {
		return []string{}
	}

	out := []string{}
	for _, oneElement := range obj.elements {
		if oneElement.Rule().Name() == name {
			out = append(out, oneElement.Code())
		}
	}

	return out
}
