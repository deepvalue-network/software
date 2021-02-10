package genesis

// CreateGenesisForTests creates a new genesis instance for tests
func CreateGenesisForTests() Genesis {
	blockBaseDiff := uint(2)
	incrPerHashDiff := float64(0.03)
	linkDiff := uint(8)
	miningValue := uint8(DefaultMiningValue)

	ins, err := NewBuilder().Create().
		WithBlockBaseDifficulty(blockBaseDiff).
		WithBlockIncreasePerHashDifficulty(incrPerHashDiff).
		WithLinkDifficulty(linkDiff).
		WithMiningValue(miningValue).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
