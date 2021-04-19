package token

type token struct {
	codeMatch CodeMatch
	code      Code
}

func createTokenWithCodeMatch(
	codeMatch CodeMatch,
) Token {
	return createTokenInternally(codeMatch, nil)
}

func createTokenWithCode(
	code Code,
) Token {
	return createTokenInternally(nil, code)
}

func createTokenInternally(
	codeMatch CodeMatch,
	code Code,
) Token {
	out := token{
		codeMatch: codeMatch,
		code:      code,
	}

	return &out
}

// IsCodeMatch returns true if there is a codeMatch, false otherwise
func (obj *token) IsCodeMatch() bool {
	return obj.codeMatch != nil
}

// CodeMatch returns the codeMatch, if any
func (obj *token) CodeMatch() CodeMatch {
	return obj.codeMatch
}

// IsCode returns true if there is a code, false otherwise
func (obj *token) IsCode() bool {
	return obj.code != nil
}

// Code returns the code, if any
func (obj *token) Code() Code {
	return obj.code
}
