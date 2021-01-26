package interpreters

import (
	"github.com/steve-care-software/products/pangolin/domain/linkers"
)

func execute(machine Machine, application linkers.Application) (StackFrame, error) {
	ins := application.Instructions().All()
	for _, oneIns := range ins {
		err := machine.Receive(oneIns)
		if err != nil {
			return nil, err
		}
	}

	return machine.StackFrame(), nil
}
