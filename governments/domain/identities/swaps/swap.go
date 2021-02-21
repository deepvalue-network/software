package swaps

type swap struct {
	in  Incoming
	out Complete
}

func createSwapWithIncoming(
	in Incoming,
) Swap {
	return createSwapInternally(in, nil)
}

func createSwapWithOutgoing(
	out Complete,
) Swap {
	return createSwapInternally(nil, out)
}

func createSwapInternally(
	in Incoming,
	compl Complete,
) Swap {
	out := swap{
		in:  in,
		out: compl,
	}

	return &out
}

// IsIncoming returns true if there is an incoming, false otherwise
func (obj *swap) IsIncoming() bool {
	return obj.in != nil
}

// Incoming returns the incoming, if any
func (obj *swap) Incoming() Incoming {
	return obj.in
}

// IsOutgoing returns true if there is an outgoing, false otherwise
func (obj *swap) IsOutgoing() bool {
	return obj.out != nil
}

// Outgoing returns the outgoing, if any
func (obj *swap) Outgoing() Complete {
	return obj.out
}
