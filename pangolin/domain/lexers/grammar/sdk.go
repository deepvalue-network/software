package grammar

import "io/ioutil"

const (
	begin                = ":"
	or                   = "|"
	end                  = ";"
	optional             = "?"
	multipleOptional     = "*"
	multipleMandatory    = "+"
	constantBegin        = "'"
	constantEnd          = "'"
	localNameDelimiter   = "@"
	grammarNamePattern   = "[a-z]{1}[a-zA-Z]*"
	rulePattern          = "[A-Z\\_]+"
	tokenPattern         = "[a-z]{1}[a-zA-Z]*"
	channelPattern       = "_[a-z]{1}[a-zA-Z]*"
	potentialWhitespaces = "[ \t\r\n]*"
	anythingExcept       = "[^%s]+"
	singleFormat         = "\\%s"
	lowercaseLetters     = "abcdefghijklmnopqrstuvwxyz"
	delimiter            = "."
)

// FileFetcher is a func that fetches daa from file
type FileFetcher func(path string) ([]byte, error)

func defaultFetch(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

var defaultName = "default"

// NewRuleTreeAdapter creates a new RuleTreeAdapter instance
func NewRuleTreeAdapter() RuleTreeAdapter {
	repetitiveRuleNodeBuilder := createRepetitiveRuleNodeBuilder()
	ruleNodeBuilder := createRuleNodeBuilder()
	ruleTreeBuilder := createRuleTreeBuilder(ruleNodeBuilder)
	return createRuleTreeAdapter(repetitiveRuleNodeBuilder, ruleNodeBuilder, ruleTreeBuilder)
}

// NewRepositoryBuilder creates a new repository builder
func NewRepositoryBuilder() RepositoryBuilder {
	// create the repository:
	rawTokenBuilder := createRawTokenBuilder()
	builder := NewBuilder()

	// rules:
	ruleSectionBuilder := createRuleSectionBuilder()
	ruleBuilder := createRuleBuilder()
	ruleAdapter := createRuleAdapter(rawTokenBuilder, ruleSectionBuilder, ruleBuilder)

	// tokens:
	tokenRuleBuilder := createTokenRuleBuilder()
	tokenSectionBuilder := createTokenSectionBuilder()
	tokenBuilder := createTokenBuilder()
	tokensBuilder := createTokensBuilder()
	tokenBlockBuilder := createTokenBlockBuilder()
	tokenBlocksBuilder := createTokenBlocksBuilder(tokenBlockBuilder, tokenSectionBuilder)
	replacementTokenBuilder := createReplacementTokenBuilder()
	replacementTokenAdapter := createReplacementTokenAdapter(replacementTokenBuilder)
	tokensAdapter := createTokensAdapter(
		rawTokenBuilder,
		tokenRuleBuilder,
		tokenBlocksBuilder,
		tokenBuilder,
		tokensBuilder,
		replacementTokenAdapter,
	)

	retrieverCriteriaRepositoryBuilder := NewRetrieverCriteriaRepositoryBuilder()
	return createRepositoryBuilder(
		retrieverCriteriaRepositoryBuilder,
		builder,
		ruleAdapter,
		tokensAdapter,
	)
}

// NewRetrieverCriteriaRepositoryBuilder returns a new NewRetrieverCriteriaRepository builder
func NewRetrieverCriteriaRepositoryBuilder() RetrieverCriteriaRepositoryBuilder {
	builder := NewRetrieverCriteriaBuilder()
	return createRetrieverCriteriaRepositoryBuilder(builder)
}

// NewRetrieverCriteriaBuilder returns a new RetrieverCriteria builder
func NewRetrieverCriteriaBuilder() RetrieverCriteriaBuilder {
	return createRetrieverCriteriaBuilder()
}

// NewBuilder creates a new grammar builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// RuleTreeAdapter represents a RuleTree adapter
type RuleTreeAdapter interface {
	ToRuleTree(grammar Grammar) (RuleTree, error)
}

// RuleTreeBuilder represents a RuleTree builder
type RuleTreeBuilder interface {
	Create() RuleTreeBuilder
	WithToken(token Token) RuleTreeBuilder
	WithNodes(nodes [][]RepetitiveRuleNode) RuleTreeBuilder
	WithGrammar(grammar Grammar) RuleTreeBuilder
	WithChildren(children map[string]RuleTree) RuleTreeBuilder
	Now() (RuleTree, error)
}

// RuleTree represents a rule tree
type RuleTree interface {
	Token() Token
	Nodes() [][]RepetitiveRuleNode
	Grammar() Grammar
	ResetNodeAtIndex(i int, j int, newNode RepetitiveRuleNode) error
}

// RepetitiveRuleNodeBuilder represents a repetitiveRuleNode builder
type RepetitiveRuleNodeBuilder interface {
	Create() RepetitiveRuleNodeBuilder
	WithNode(node RuleNode) RepetitiveRuleNodeBuilder
	WithRecursiveName(recursiveName string) RepetitiveRuleNodeBuilder
	IsMultipleMandatory() RepetitiveRuleNodeBuilder
	IsMultipleOptional() RepetitiveRuleNodeBuilder
	IsOptional() RepetitiveRuleNodeBuilder
	Now() (RepetitiveRuleNode, error)
}

// RepetitiveRuleNode represents a repetitive rule node
type RepetitiveRuleNode interface {
	HasRecursiveName() bool
	RecursiveName() string
	HasNode() bool
	Node() RuleNode
	IsMultipleMandatory() bool
	IsMultipleOptional() bool
	IsOptional() bool
	SetNode(node RuleNode)
}

// RuleNodeBuilder represents a RuleNode builder
type RuleNodeBuilder interface {
	Create() RuleNodeBuilder
	WithRule(rule Rule) RuleNodeBuilder
	WithLeaf(leaf RuleTree) RuleNodeBuilder
	Now() (RuleNode, error)
}

// RuleNode represents a rule node
type RuleNode interface {
	HasRule() bool
	Rule() Rule
	HasLeaf() bool
	Leaf() RuleTree
	ResetLeaf(leaf RuleTree)
}

// PositionApplicationBuilder represents a positionApplication builder
type PositionApplicationBuilder interface {
	Create() PositionApplicationBuilder
	WithGrammar(grammar Grammar) PositionApplicationBuilder
	Now() (PositionApplication, error)
}

// PositionApplication represents a position application
type PositionApplication interface {
	FindChannels(script string) ([]Position, error)
}

// PositionBuilder represents a position builder
type PositionBuilder interface {
	Create() PositionBuilder
	From(from int) PositionBuilder
	To(to int) PositionBuilder
	Now() (Position, error)
}

// Position represents a script position
type Position interface {
	From() int
	To() int
}

// RetrieverCriteriaBuilder represents the retriever criteria builder
type RetrieverCriteriaBuilder interface {
	Create() RetrieverCriteriaBuilder
	WithName(name string) RetrieverCriteriaBuilder
	WithRoot(root string) RetrieverCriteriaBuilder
	WithBaseDirPath(baseDirPath string) RetrieverCriteriaBuilder
	WithTokensPath(tokensPath string) RetrieverCriteriaBuilder
	WithChannelsPath(channelsPath string) RetrieverCriteriaBuilder
	WithRulesPath(rulesPath string) RetrieverCriteriaBuilder
	WithExtends(extends []RetrieverCriteria) RetrieverCriteriaBuilder
	Now() (RetrieverCriteria, error)
}

// RetrieverCriteria represents a repository retriever criteria
type RetrieverCriteria interface {
	Name() string
	Root() string
	BaseDirPath() string
	TokensPath() string
	RulesPath() string
	HasChannelsPath() bool
	ChannelsPath() string
	HasExtends() bool
	Extends() map[string]RetrieverCriteria
}

// RetrieverCriteriaRepositoryBuilder represents a retrieverCriteriaRepository builder
type RetrieverCriteriaRepositoryBuilder interface {
	Create() RetrieverCriteriaRepositoryBuilder
	WithName(name string) RetrieverCriteriaRepositoryBuilder
	WithRoot(root string) RetrieverCriteriaRepositoryBuilder
	WithFileFetcher(fileFetcher FileFetcher) RetrieverCriteriaRepositoryBuilder
	Now() (RetrieverCriteriaRepository, error)
}

// RetrieverCriteriaRepository represents a retrieverRetrieverCriteria repository
type RetrieverCriteriaRepository interface {
	Retrieve(filePath string) (RetrieverCriteria, error)
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithFileFetcher(fileFetcher FileFetcher) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a grammar repository
type Repository interface {
	Retrieve(criteria RetrieverCriteria) (Grammar, error)
	RetrieveFromFile(rootPattern string, name string, filePath string) (Grammar, error)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithRoot(root string) Builder
	WithChannels(channels Tokens) Builder
	WithTokens(tokens Tokens) Builder
	WithRules(rules []Rule) Builder
	WithSubGrammars(grammars map[string]Grammar) Builder
	WithGrammar(gr Grammar) Builder
	Now() (Grammar, error)
}

// Grammar represents the grammar
type Grammar interface {
	Name() string
	Root() string
	RootToken() Token
	Rules() map[string]Rule
	HasTokens() bool
	Tokens() Tokens
	HasChannels() bool
	Channels() Tokens
	HasSubGrammars() bool
	SubGrammars() map[string]Grammar
	FetchByName(name string) (Grammar, error)
}

// TokensAdapter represents a tokens adapter
type TokensAdapter interface {
	ToTokens(script string, tokenPattern string, replacementTokenPattern string, grammarName string, extends map[string]Grammar, rules []Rule) (Tokens, error)
}

// TokensBuilder represents a tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithTokens(tokens map[string]Token) TokensBuilder
	WithReplacements(replacements []ReplacementToken) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	Tokens() map[string]Token
	Replace(tok Token) Tokens
	HasReplacements() bool
	Replacements() []ReplacementToken
}

// ReplacementTokenAdapter represents a replacement token adapter
type ReplacementTokenAdapter interface {
	ToReplacementTokens(script string, pattern string) ([]ReplacementToken, error)
}

// ReplacementTokenBuilder represents a replacement token builder
type ReplacementTokenBuilder interface {
	Create() ReplacementTokenBuilder
	WithToGrammar(toGrammar string) ReplacementTokenBuilder
	WithFromToken(toToken string) ReplacementTokenBuilder
	Now() (ReplacementToken, error)
}

// ReplacementToken represents a replacement token
type ReplacementToken interface {
	ToGrammar() string
	FromToken() string
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithBlocks(blocks []TokenBlocks) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Blocks() []TokenBlocks
	SubTokenNames() []string
	SetName(name string)
}

// TokenBlocksBuilder represents a tokenBlocks builder
type TokenBlocksBuilder interface {
	Create() TokenBlocksBuilder
	WithRules(rules []TokenRule) TokenBlocksBuilder
	WithTokens(tokens []RawToken) TokenBlocksBuilder
	WithOptionals(optionals []RawToken) TokenBlocksBuilder
	WithMultipleOptionals(multipleOptionals []RawToken) TokenBlocksBuilder
	WithMultipleMandatories(multipleMandatories []RawToken) TokenBlocksBuilder
	Now() (TokenBlocks, error)
}

// TokenBlocks represents a tokenBlocks
type TokenBlocks interface {
	Get() []TokenBlock
	SubTokenNames() []string
	RulesToken() ([]Rule, string)
}

// TokenBlockBuilder represents a tokenBlock builder
type TokenBlockBuilder interface {
	Create() TokenBlockBuilder
	WithSection(section TokenSection) TokenBlockBuilder
	WithOptional(optional RawToken) TokenBlockBuilder
	WithMultipleOptional(multipleOptional RawToken) TokenBlockBuilder
	WithMultipleMandatory(multipleMandatory RawToken) TokenBlockBuilder
	Now() (TokenBlock, error)
}

// TokenBlock represents a tokenBlock
type TokenBlock interface {
	Section() TokenSection
	HasOptional() bool
	Optional() RawToken
	HasMultipleOptional() bool
	MultipleOptional() RawToken
	HasMultipleMandatory() bool
	MultipleMandatory() RawToken
}

// TokenRuleBuilder represents a tokenRule builder
type TokenRuleBuilder interface {
	Create() TokenRuleBuilder
	WithRule(rule Rule) TokenRuleBuilder
	WithRawToken(rawToken RawToken) TokenRuleBuilder
	Now() (TokenRule, error)
}

// TokenRule represents a token rule
type TokenRule interface {
	Rule() Rule
	RawToken() RawToken
}

// TokenSectionBuilder represents a tokenSection builder
type TokenSectionBuilder interface {
	Create() TokenSectionBuilder
	WithRule(rule TokenRule) TokenSectionBuilder
	WithToken(token RawToken) TokenSectionBuilder
	Now() (TokenSection, error)
}

// TokenSection represents a token section
type TokenSection interface {
	HasRule() bool
	Rule() TokenRule
	HasToken() bool
	Token() RawToken
	NextRuleToken() (Rule, string)
}

// RuleAdapter represents a rule adapter
type RuleAdapter interface {
	ToRules(script string, grammarName string) ([]Rule, error)
}

// RuleBuilder represents a rule builder
type RuleBuilder interface {
	Create() RuleBuilder
	WithName(name string) RuleBuilder
	WithSections(sections []RuleSection) RuleBuilder
	Now() (Rule, error)
}

// Rule represents a rule
type Rule interface {
	FindFirst(str string) (string, bool, error)
	FindConsecutives(str string) (string, bool, error)
	Name() string
	Sections() []RuleSection
	Index() int
}

// RuleSectionBuilder represents a ruleSection builder
type RuleSectionBuilder interface {
	Create() RuleSectionBuilder
	WithConstant(constant RawToken) RuleSectionBuilder
	WithPattern(pattern RawToken) RuleSectionBuilder
	Now() (RuleSection, error)
}

// RuleSection represents a rule section
type RuleSection interface {
	FindFirst(str string) (string, error)
	HasConstant() bool
	Constant() RawToken
	HasPattern() bool
	Pattern() RawToken
}

// RawTokenBuilder represents a rawToken builder
type RawTokenBuilder interface {
	Create() RawTokenBuilder
	WithName(name string) RawTokenBuilder
	WithValue(value string) RawTokenBuilder
	WithCode(code string) RawTokenBuilder
	WithIndex(index int) RawTokenBuilder
	WithGrammar(grammar string) RawTokenBuilder
	Now() (RawToken, error)
}

// RawToken represents a raw token
type RawToken interface {
	Name() string
	Value() string
	Reference() string
	Code() string
	Index() int
	Grammar() string
}
