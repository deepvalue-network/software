package parsers

type targetSingle struct {
	evts []Event
	path RelativePath
}

func createTargetSingleWithEvents(
	evts []Event,
) TargetSingle {
	return createTargetSingleInternally(evts, nil)
}

func createTargetSingleWithPath(
	path RelativePath,
) TargetSingle {
	return createTargetSingleInternally(nil, path)
}

func createTargetSingleInternally(
	evts []Event,
	path RelativePath,
) TargetSingle {
	out := targetSingle{
		evts: evts,
		path: path,
	}

	return &out
}

// IsEvents returns true if there is events, false otherwise
func (obj *targetSingle) IsEvents() bool {
	return obj.evts != nil
}

// Events returns the events, if any
func (obj *targetSingle) Events() []Event {
	return obj.evts
}

// IsPath returns true if there is a path, false otherwise
func (obj *targetSingle) IsPath() bool {
	return obj.path != nil
}

// Path returns the path, if any
func (obj *targetSingle) Path() RelativePath {
	return obj.path
}
