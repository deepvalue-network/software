package parsers

type relativePaths struct {
	list []RelativePath
}

func createRelativePaths(
	list []RelativePath,
) RelativePaths {
	out := relativePaths{
		list: list,
	}

	return &out
}

// All returns the relative paths
func (obj *relativePaths) All() []RelativePath {
	return obj.list
}
