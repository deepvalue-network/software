package parsers

type arythmetic struct {
	add StandardOperation
	sub StandardOperation
	mul StandardOperation
	div RemainingOperation
}

func createArythmeticWithAddition(add StandardOperation) Arythmetic {
	return createArythmeticInternally(add, nil, nil, nil)
}

func createArythmeticWithSubstraction(sub StandardOperation) Arythmetic {
	return createArythmeticInternally(nil, sub, nil, nil)
}

func createArythmeticWithMultiplication(mul StandardOperation) Arythmetic {
	return createArythmeticInternally(nil, nil, mul, nil)
}

func createArythmeticWithDivision(div RemainingOperation) Arythmetic {
	return createArythmeticInternally(nil, nil, nil, div)
}

func createArythmeticInternally(add StandardOperation, sub StandardOperation, mul StandardOperation, div RemainingOperation) Arythmetic {
	out := arythmetic{
		add: add,
		sub: sub,
		mul: mul,
		div: div,
	}

	return &out
}

// IsAdd returns true if the arythmetic operator is an addition, false otherwise
func (obj *arythmetic) IsAdd() bool {
	return obj.add != nil
}

// Add returns the addition operator, if any
func (obj *arythmetic) Add() StandardOperation {
	return obj.add
}

// IsSub returns true if the arythmetic operator is a substraction, false otherwise
func (obj *arythmetic) IsSub() bool {
	return obj.sub != nil
}

// Sub returns the substraction operator, if any
func (obj *arythmetic) Sub() StandardOperation {
	return obj.sub
}

// IsMul returns true if the arythmetic operator is a multiplication, false otherwise
func (obj *arythmetic) IsMul() bool {
	return obj.mul != nil
}

// Mul returns the multiplication operator, if any
func (obj *arythmetic) Mul() StandardOperation {
	return obj.mul
}

// IsDiv returns true if the arythmetic operator is a division, false otherwise
func (obj *arythmetic) IsDiv() bool {
	return obj.div != nil
}

// Div returns the division operator, if any
func (obj *arythmetic) Div() RemainingOperation {
	return obj.div
}
