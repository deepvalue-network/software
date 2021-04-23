package value

type operation struct {
	isPrint bool
}

func createOperationWithPrint() Operation {
	return createOperationInternally(true)
}

func createOperationInternally(isPrint bool) Operation {
	out := operation{
		isPrint: isPrint,
	}

	return &out
}

// IsPrint returns true if the operation is print, false otherwise
func (obj *operation) IsPrint() bool {
	return obj.isPrint
}
