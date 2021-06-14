package tokens

type token struct {
	name  string
	block Block
}

func createToken(
	name string,
	block Block,
) Token {
	out := token{
		name:  name,
		block: block,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Block returns the block
func (obj *token) Block() Block {
	return obj.block
}
