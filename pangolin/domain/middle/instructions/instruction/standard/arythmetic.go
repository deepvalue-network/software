package standard

type arythmetic struct {
	isAdd bool
	isSub bool
	isMul bool
}

func createArythmeticWithAdd() Arythmetic {
	return createArythmeticInternally(true, false, false)
}

func createArythmeticWithSub() Arythmetic {
	return createArythmeticInternally(false, true, false)
}

func createArythmeticWithMul() Arythmetic {
	return createArythmeticInternally(false, false, true)
}

func createArythmeticInternally(isAdd bool, isSub bool, isMul bool) Arythmetic {
	out := arythmetic{
		isAdd: isAdd,
		isSub: isSub,
		isMul: isMul,
	}

	return &out
}

// IsAdd returns true if add, false otherwise
func (obj *arythmetic) IsAdd() bool {
	return obj.isAdd
}

// IsSub returns true if sub, false otherwise
func (obj *arythmetic) IsSub() bool {
	return obj.isSub
}

// IsMul returns true if mul, false otherwise
func (obj *arythmetic) IsMul() bool {
	return obj.isMul
}
