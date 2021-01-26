package token

type code struct {
	ret     string
	token   string
	pattern string
	amount  string
}

func createCode(
	ret string,
	token string,
) Code {
	return createCodeInternally(ret, token, "", "")
}

func createCodeWithPattern(
	ret string,
	token string,
	pattern string,
) Code {
	return createCodeInternally(ret, token, pattern, "")
}

func createCodeWithAmount(
	ret string,
	token string,
	amount string,
) Code {
	return createCodeInternally(ret, token, "", amount)
}

func createCodeWithPatternAndAmount(
	ret string,
	token string,
	pattern string,
	amount string,
) Code {
	return createCodeInternally(ret, token, pattern, amount)
}

func createCodeInternally(
	ret string,
	token string,
	pattern string,
	amount string,
) Code {
	out := code{
		ret:     ret,
		token:   token,
		pattern: pattern,
		amount:  amount,
	}

	return &out
}

// Return returns the return variable
func (obj *code) Return() string {
	return obj.ret
}

// Token returns the token variable
func (obj *code) Token() string {
	return obj.token
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
