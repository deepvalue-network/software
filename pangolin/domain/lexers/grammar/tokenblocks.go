package grammar

type tokenBlocks struct {
	blocks []TokenBlock
}

func createTokenBlocks(blocks []TokenBlock) TokenBlocks {
	out := tokenBlocks{
		blocks: blocks,
	}

	return &out
}

// Get returns the tokenBlocks
func (obj *tokenBlocks) Get() []TokenBlock {
	return obj.blocks
}

// SubTokenNames returns the subToken names
func (obj *tokenBlocks) SubTokenNames() []string {
	names := []string{}
	for _, oneBlock := range obj.blocks {
		section := oneBlock.Section()
		if section.HasToken() {
			isUnique := true
			val := section.Token().Name()
			for _, oneName := range names {
				if oneName == val {
					isUnique = false
					break
				}
			}

			if isUnique {
				names = append(names, section.Token().Name())
			}
		}
	}

	return names
}

// RuleTokenAtIndex returns the RuleSection or Token name at index
func (obj *tokenBlocks) RulesToken() ([]Rule, string) {
	out := []Rule{}
	for _, oneTokenBlock := range obj.blocks {
		rule, tokenName := oneTokenBlock.Section().NextRuleToken()
		if tokenName != "" {
			if len(out) > 0 {
				return out, ""
			}

			return nil, tokenName
		}

		out = append(out, rule)
	}

	return out, ""
}
