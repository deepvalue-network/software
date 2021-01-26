package parsers

type variableDirection struct {
	incoming   VariableIncoming
	isOutgoing bool
}

func createVariableDirectionWithIncoming(incoming VariableIncoming) VariableDirection {
	return createVariableDirectionInternally(incoming, false)
}

func createVariableDirectionWithOutgoing() VariableDirection {
	return createVariableDirectionInternally(nil, true)
}

func createVariableDirectionWithIncomingAndOutgoing(incoming VariableIncoming) VariableDirection {
	return createVariableDirectionInternally(incoming, true)
}

func createVariableDirectionInternally(
	incoming VariableIncoming,
	isOutgoing bool,
) VariableDirection {
	out := variableDirection{
		incoming:   incoming,
		isOutgoing: isOutgoing,
	}

	return &out
}

// IsIncoming returns true if the variable is incoming, false otherwise
func (obj *variableDirection) IsIncoming() bool {
	return obj.incoming != nil
}

// Incoming returns the VariableIncoming instance, if any
func (obj *variableDirection) Incoming() VariableIncoming {
	return obj.incoming
}

// IsOutgoing returns true if the variable is outgoing, false otherwise
func (obj *variableDirection) IsOutgoing() bool {
	return obj.isOutgoing
}
