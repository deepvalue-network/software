package grammars

type file struct {
	root         string
	tokensPath   string
	rulesPath    string
	channelsPath string
}

func createFile(
	root string,
	tokensPath string,
	rulesPath string,
) File {
	return createFileInternally(root, tokensPath, rulesPath, "")
}

func createFileWithChannelsPath(
	root string,
	tokensPath string,
	rulesPath string,
	channelsPath string,
) File {
	return createFileInternally(root, tokensPath, rulesPath, channelsPath)
}

func createFileInternally(
	root string,
	tokensPath string,
	rulesPath string,
	channelsPath string,
) File {
	out := file{
		root:         root,
		tokensPath:   tokensPath,
		rulesPath:    rulesPath,
		channelsPath: channelsPath,
	}

	return &out
}

// Root returns the root
func (obj *file) Root() string {
	return obj.root
}

// TokensPath returns the tokens path
func (obj *file) TokensPath() string {
	return obj.tokensPath
}

// RulesPath returns the rules path
func (obj *file) RulesPath() string {
	return obj.rulesPath
}

// HasChannelsPath returns true if there is a channelsPath, false otherwise
func (obj *file) HasChannelsPath() bool {
	return obj.channelsPath != ""
}

// ChannelsPath returns the channels path, if any
func (obj *file) ChannelsPath() string {
	return obj.channelsPath
}
