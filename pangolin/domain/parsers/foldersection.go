package parsers

import (
	"fmt"
	"path/filepath"
)

type folderSection struct {
	isTail bool
	name   FolderName
}

func createFolderSection(name FolderName) FolderSection {
	return createFolderSectionInternally(false, name)
}

func createFolderSectionWithTail(name FolderName) FolderSection {
	return createFolderSectionInternally(true, name)
}

func createFolderSectionInternally(
	isTail bool,
	name FolderName,
) FolderSection {
	out := folderSection{
		isTail: isTail,
		name:   name,
	}

	return &out
}

// IsTail returns true if the section is a tail, false otherwise
func (obj *folderSection) IsTail() bool {
	return obj.isTail
}

// Name returns the section name
func (obj *folderSection) Name() FolderName {
	return obj.name
}

// String returns the name as string
func (obj *folderSection) String() string {
	if obj.IsTail() {
		return obj.name.String()
	}

	return fmt.Sprintf("%s%s", obj.Name().String(), string(filepath.Separator))
}
