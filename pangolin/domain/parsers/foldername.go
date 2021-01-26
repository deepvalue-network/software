package parsers

type folderName struct {
	isCurrent  bool
	isPrevious bool
	name       string
}

func createFolderNameWithCurrent() FolderName {
	return createFolderNameInternally(true, false, "")
}

func createFolderNameWithPrevious() FolderName {
	return createFolderNameInternally(false, true, "")
}

func createFolderNameWithName(name string) FolderName {
	return createFolderNameInternally(false, false, name)
}

func createFolderNameInternally(
	isCurrent bool,
	isPrevious bool,
	name string,
) FolderName {
	out := folderName{
		isCurrent:  isCurrent,
		isPrevious: isPrevious,
		name:       name,
	}

	return &out
}

// IsCurrent returns true if the folder is the current folder, false otherwise
func (obj *folderName) IsCurrent() bool {
	return obj.isCurrent
}

// IsPrevious returns true if the folder is the previous folder, false otherwise
func (obj *folderName) IsPrevious() bool {
	return obj.isPrevious
}

// IsName returns true if there is a folder name, false otherwise
func (obj *folderName) IsName() bool {
	return obj.name != ""
}

// Name returns the name, if any
func (obj *folderName) Name() string {
	return obj.name
}

// String returns the name as string
func (obj *folderName) String() string {
	if obj.IsCurrent() {
		return "."
	}

	if obj.IsPrevious() {
		return ".."
	}

	return obj.Name()
}
