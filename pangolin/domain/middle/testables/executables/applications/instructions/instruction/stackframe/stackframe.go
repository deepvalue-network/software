package stackframe

type stackframe struct {
	isPush bool
	isPop  bool
	index  string
	skip   Skip
	save   Save
	swtch  string
}

func createStackframeWithPush() Stackframe {
	return createStackframeInternally(true, false, "", nil, nil, "")
}

func createStackframeWithPop() Stackframe {
	return createStackframeInternally(false, true, "", nil, nil, "")
}

func createStackframeWithIndex(index string) Stackframe {
	return createStackframeInternally(false, false, index, nil, nil, "")
}

func createStackframeWithSkip(skip Skip) Stackframe {
	return createStackframeInternally(false, false, "", skip, nil, "")
}

func createStackframeWithSave(save Save) Stackframe {
	return createStackframeInternally(false, false, "", nil, save, "")
}

func createStackframeWithSwitch(swtch string) Stackframe {
	return createStackframeInternally(false, false, "", nil, nil, swtch)
}

func createStackframeInternally(
	isPush bool,
	isPop bool,
	index string,
	skip Skip,
	save Save,
	swtch string,
) Stackframe {
	out := stackframe{
		isPush: isPush,
		isPop:  isPop,
		index:  index,
		skip:   skip,
		save:   save,
		swtch:  swtch,
	}

	return &out
}

// IsPush returns true if the instruction is push
func (obj *stackframe) IsPush() bool {
	return obj.isPush
}

// IsPop returns true if the instruction is pop
func (obj *stackframe) IsPop() bool {
	return obj.isPop
}

// IsIndex returns true if there is an index, false otherwise
func (obj *stackframe) IsIndex() bool {
	return obj.index != ""
}

// Index returns the index, if any
func (obj *stackframe) Index() string {
	return obj.index
}

// IsSkip returns true if there is a skip, false otherwise
func (obj *stackframe) IsSkip() bool {
	return obj.skip != nil
}

// Skip returns the skip, if any
func (obj *stackframe) Skip() Skip {
	return obj.skip
}

// IsSave retruns true if there is a save instance, false otherwise
func (obj *stackframe) IsSave() bool {
	return obj.save != nil
}

// Save returns the save instance, if any
func (obj *stackframe) Save() Save {
	return obj.save
}

// IsSwitch retruns true if there is a switch instance, false otherwise
func (obj *stackframe) IsSwitch() bool {
	return obj.swtch != ""
}

// Switch returns the switch instance, if any
func (obj *stackframe) Switch() string {
	return obj.swtch
}
