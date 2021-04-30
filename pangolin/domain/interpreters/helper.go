package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	language_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
)

func createMachineFromLabels(
	machineBuilder MachineBuilder,
	stackFrame StackFrame,
	labels labels.Labels,
) (Machine, error) {
	var machine Machine
	callLabelFunc := func(name string) error {
		lbl, err := labels.Fetch(name)
		if err != nil {
			return err
		}

		insList := lbl.Instructions().All()
		for _, oneIns := range insList {
			stop, err := machine.ReceiveLbl(oneIns)
			if err != nil {
				return err
			}

			if stop {
				return nil
			}
		}

		return nil
	}

	fetchStackframeFunc := func() StackFrame {
		return stackFrame
	}

	mach, err := machineBuilder.Create().WithCallLabelFunc(callLabelFunc).WithFetchStackFunc(fetchStackframeFunc).Now()
	if err != nil {
		return nil, err
	}

	machine = mach
	return machine, nil
}

func createMachineFromLanguageLabels(
	machineBuilder MachineBuilder,
	stackFrame StackFrame,
	labels language_labels.Labels,
) (Machine, error) {
	var machineLang MachineLanguage
	callLabelFunc := func(name string) error {
		lbl, err := labels.Fetch(name)
		if err != nil {
			return err
		}

		insList := lbl.Instructions().All()
		for _, oneIns := range insList {
			stop, err := machineLang.ReceiveLbl(oneIns)
			if err != nil {
				return err
			}

			if stop {
				return nil
			}
		}

		return nil
	}

	fetchStackframeFunc := func() StackFrame {
		return stackFrame
	}

	mach, err := machineBuilder.Create().WithCallLabelFunc(callLabelFunc).WithFetchStackFunc(fetchStackframeFunc).Now()
	if err != nil {
		return nil, err
	}

	return mach, nil
}
