package grammar

type retrieverCriteria struct {
	name         string
	root         string
	baseDirPath  string
	tokensPath   string
	rulesPath    string
	channelsPath string
	extends      map[string]RetrieverCriteria
}

func createRetrieverCriteria(
	name string,
	root string,
	baseDirPath string,
	tokensPath string,
	rulesPath string,
) RetrieverCriteria {
	return createRetrieverCriteriaInternally(name, root, baseDirPath, tokensPath, rulesPath, "", nil)
}

func createRetrieverCriteriaWithChannelsPath(
	name string,
	root string,
	baseDirPath string,
	tokensPath string,
	rulesPath string,
	channelsPath string,
) RetrieverCriteria {
	return createRetrieverCriteriaInternally(name, root, baseDirPath, tokensPath, rulesPath, channelsPath, nil)
}

func createRetrieverCriteriaWithExtends(
	name string,
	root string,
	baseDirPath string,
	tokensPath string,
	rulesPath string,
	extends map[string]RetrieverCriteria,
) RetrieverCriteria {
	return createRetrieverCriteriaInternally(name, root, baseDirPath, tokensPath, rulesPath, "", extends)
}

func createRetrieverCriteriaWithChannelsPathAndExtends(
	name string,
	root string,
	baseDirPath string,
	tokensPath string,
	rulesPath string,
	channelsPath string,
	extends map[string]RetrieverCriteria,
) RetrieverCriteria {
	return createRetrieverCriteriaInternally(name, root, baseDirPath, tokensPath, rulesPath, channelsPath, extends)
}

func createRetrieverCriteriaInternally(
	name string,
	root string,
	baseDirPath string,
	tokensPath string,
	rulesPath string,
	channelsPath string,
	extends map[string]RetrieverCriteria,
) RetrieverCriteria {
	out := retrieverCriteria{
		name:         name,
		root:         root,
		baseDirPath:  baseDirPath,
		tokensPath:   tokensPath,
		rulesPath:    rulesPath,
		channelsPath: channelsPath,
		extends:      extends,
	}

	return &out
}

// Name returns the name
func (obj *retrieverCriteria) Name() string {
	return obj.name
}

// Root returns the root
func (obj *retrieverCriteria) Root() string {
	return obj.root
}

// BaseDirPath returns the base dir path
func (obj *retrieverCriteria) BaseDirPath() string {
	return obj.baseDirPath
}

// TokensPath returns the tokens path
func (obj *retrieverCriteria) TokensPath() string {
	return obj.tokensPath
}

// RulesPath returns the rules path
func (obj *retrieverCriteria) RulesPath() string {
	return obj.rulesPath
}

// HasChannelsPath returns true if there is a channels path, false otherwise
func (obj *retrieverCriteria) HasChannelsPath() bool {
	return obj.channelsPath != ""
}

// ChannelsPath returns the channels path
func (obj *retrieverCriteria) ChannelsPath() string {
	return obj.channelsPath
}

// HasExtends returns true if the grammar extends another grammar, false otherwise
func (obj *retrieverCriteria) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *retrieverCriteria) Extends() map[string]RetrieverCriteria {
	return obj.extends
}
