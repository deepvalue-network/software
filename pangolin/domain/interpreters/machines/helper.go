package machines

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	language_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
)

func fromLabelsToCallLabelByNameFunc(stackframe stackframes.StackFrame, labels labels.Labels) (CallLabelByNameFn, error) {
	var insMachine Instruction
	fn := func(name string) error {
		lbl, err := labels.Fetch(name)
		if err != nil {
			return err
		}

		lblInsList := lbl.Instructions().All()
		for _, oneLblIns := range lblInsList {
			stop, err := insMachine.ReceiveLbl(oneLblIns)
			if err != nil {
				return err
			}

			if stop {
				break
			}
		}

		return nil
	}

	machine, err := NewInstructionBuilder().Create().WithCallLabelFn(fn).WithStackFrame(stackframe).Now()
	if err != nil {
		return nil, err
	}

	insMachine = machine
	return fn, nil
}

func fromLanguageLabelsToCallLabelByNameFunc(stackframe stackframes.StackFrame, labels language_labels.Labels) (CallLabelByNameFn, error) {
	var insMachine LanguageInstruction
	fn := func(name string) error {
		lbl, err := labels.Fetch(name)
		if err != nil {
			return err
		}

		lblInsList := lbl.Instructions().All()
		for _, oneLblIns := range lblInsList {
			stop, err := insMachine.ReceiveLbl(oneLblIns)
			if err != nil {
				return err
			}

			if stop {
				break
			}
		}

		return nil
	}

	/*machine, err := NewLanguageInstructionBuilder().Create().WithCallLabelFn(fn).WithStackFrame(stackframe).Now()
	if err != nil {
		return nil, err
	}

	insMachine = machine*/
	return fn, nil
}
