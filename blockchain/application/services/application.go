package services

type application struct {
	block      Block
	minedBlock MinedBlock
	link       Link
	minedLink  MinedLink
	chain      Chain
}

func createApplication(
	block Block,
	minedBlock MinedBlock,
	link Link,
	minedLink MinedLink,
	chain Chain,
) Application {
	out := application{
		block:      block,
		minedBlock: minedBlock,
		link:       link,
		minedLink:  minedLink,
		chain:      chain,
	}

	return &out
}

// Block returns the block application
func (obj application) Block() Block {
	return obj.block
}

// MinedBlock returns the mined block application
func (obj application) MinedBlock() MinedBlock {
	return obj.minedBlock
}

// Link returns the link application
func (obj application) Link() Link {
	return obj.link
}

// MinedLink returns the mined link application
func (obj application) MinedLink() MinedLink {
	return obj.minedLink
}

// Chain returns the chain application
func (obj application) Chain() Chain {
	return obj.chain
}
