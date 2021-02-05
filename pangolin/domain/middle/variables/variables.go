package variables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
)

type variables struct {
	vr map[string]variable.Variable
}

func createVariables(vr map[string]variable.Variable) Variables {
	out := variables{
		vr: vr,
	}

	return &out
}

// All return all variables
func (obj *variables) All() []variable.Variable {
	out := []variable.Variable{}
	for _, oneVar := range obj.vr {
		out = append(out, oneVar)
	}

	return out
}

// Merge merge the variables
func (obj *variables) Merge(vr Variables) Variables {
	for _, oneVar := range obj.vr {
		name := oneVar.Name()
		obj.vr[name] = oneVar
	}

	return obj
}
