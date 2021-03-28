package token

type code struct {
	ret     string
	pattern string
	amount  string
}

func createCode(
	ret string,
) Code {
	return createCodeInternally(ret, "", "")
}

func createCodeWithPattern(
	ret string,
	pattern string,
) Code {
	return createCodeInternally(ret, pattern, "")
}

func createCodeWithAmount(
	ret string,
	amount string,
) Code {
	return createCodeInternally(ret, "", amount)
}

func createCodeWithPatternAndAmount(
	ret string,
	pattern string,
	amount string,
) Code {
	return createCodeInternally(ret, pattern, amount)
}

func createCodeInternally(
	ret string,
	pattern string,
	amount string,
) Code {
	out := code{
		ret:     ret,
		pattern: pattern,
		amount:  amount,
	}

	return &out
}

// Return returns the return variable
func (obj *code) Return() string {
	return obj.ret
}

// HasPattern returns true if there is a pattern, false otherwise
func (obj *code) HasPattern() bool {
	return obj.pattern != ""
}

// Pattern returns the pattern, if any
func (obj *code) Pattern() string {
	return obj.pattern
}

// HasAmount returns true if there is an amount, false otherwise
func (obj *code) HasAmount() bool {
	return obj.amount != ""
}

// Amount returns the amount, if any
func (obj *code) Amount() string {
	return obj.amount
}
