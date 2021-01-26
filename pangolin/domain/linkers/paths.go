package linkers

type paths struct {
	baseDir  string
	tokens   string
	rules    string
	logics   string
	channels string
}

func createPaths(
	baseDir string,
	tokens string,
	rules string,
	logics string,
) Paths {
	return createPathsInternally(baseDir, tokens, rules, logics, "")
}

func createPathsWithChannels(
	baseDir string,
	tokens string,
	rules string,
	logics string,
	channels string,
) Paths {
	return createPathsInternally(baseDir, tokens, rules, logics, channels)
}

func createPathsInternally(
	baseDir string,
	tokens string,
	rules string,
	logics string,
	channels string,
) Paths {
	out := paths{
		baseDir:  baseDir,
		tokens:   tokens,
		rules:    rules,
		logics:   logics,
		channels: channels,
	}

	return &out
}

// BaseDir returns the base dir
func (obj *paths) BaseDir() string {
	return obj.baseDir
}

// Tokens returns the tokens
func (obj *paths) Tokens() string {
	return obj.tokens
}

// Rules returns the rules
func (obj *paths) Rules() string {
	return obj.rules
}

// Logics returns the logics
func (obj *paths) Logics() string {
	return obj.logics
}

// HasChannels returns true if there is channels, false otherwise
func (obj *paths) HasChannels() bool {
	return obj.channels != ""
}

// Channels returns the chanels, if any
func (obj *paths) Channels() string {
	return obj.channels
}
