package parsers

import (
	"strings"
)

type relativePath struct {
	head []FolderSection
	tail FolderSection
}

func createRelativePath(
	head []FolderSection,
) RelativePath {
	return createRelativePathInternally(head, nil)
}

func createRelativePathWithTail(
	head []FolderSection,
	tail FolderSection,
) RelativePath {
	return createRelativePathInternally(head, tail)
}

func createRelativePathInternally(
	head []FolderSection,
	tail FolderSection,
) RelativePath {
	out := relativePath{
		head: head,
		tail: tail,
	}

	return &out
}

// All returns all the sections
func (obj *relativePath) All() []FolderSection {
	if !obj.HasTail() {
		return obj.Head()
	}

	return append(obj.Head(), obj.Tail())
}

// Head returns the head
func (obj *relativePath) Head() []FolderSection {
	return obj.head
}

// HasTail returns true if there is a tail, false otherwise
func (obj *relativePath) HasTail() bool {
	return obj.tail != nil
}

// Tail returns the tail, if any
func (obj *relativePath) Tail() FolderSection {
	return obj.tail
}

// String returns the path as string
func (obj *relativePath) String() string {
	folders := []string{}
	sections := obj.All()
	for _, oneSection := range sections {
		folders = append(folders, oneSection.String())
	}

	return strings.Join(folders, "")
}
