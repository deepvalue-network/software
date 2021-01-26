package parsers

type patternLabels struct {
	enter string
	exit  string
}

func createPatternLabelsWithEnterLabel(
	enterLabel string,
) PatternLabels {
	return createPatternLabelsInternally(enterLabel, "")
}

func createPatternLabelsWithExitLabel(
	exitLabel string,
) PatternLabels {
	return createPatternLabelsInternally("", exitLabel)
}

func createPatternLabelsWithEnterLabelAndExitLabel(
	enterLabel string,
	exitLabel string,
) PatternLabels {
	return createPatternLabelsInternally(enterLabel, exitLabel)
}

func createPatternLabelsInternally(
	enter string,
	exit string,
) PatternLabels {
	out := patternLabels{
		enter: enter,
		exit:  exit,
	}

	return &out
}

// HasEnterLabel returns true if there is an enter label, false otherwise
func (obj *patternLabels) HasEnterLabel() bool {
	return obj.enter != ""
}

// EnterLabel returns the enter label, if any
func (obj *patternLabels) EnterLabel() string {
	return obj.enter
}

// HasExitLabel returns true if there is an exit label, false otherwise
func (obj *patternLabels) HasExitLabel() bool {
	return obj.exit != ""
}

// ExitLabel returns the exit label, if any
func (obj *patternLabels) ExitLabel() string {
	return obj.exit
}
