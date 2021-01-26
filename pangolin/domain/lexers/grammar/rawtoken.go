package grammar

import "fmt"

type rawToken struct {
	value string
	code  string
	index int
	gr    string
}

func createRawToken(value string, code string, index int, gr string) RawToken {
	return createRawTokenInternally(value, code, index, gr)
}

func createRawTokenInternally(value string, code string, index int, gr string) RawToken {
	out := rawToken{
		value: value,
		code:  code,
		index: index,
		gr:    gr,
	}

	return &out
}

// Value returns the value
func (obj *rawToken) Value() string {
	return obj.value
}

// Reference returns the reference
func (obj *rawToken) Reference() string {
	return fmt.Sprintf("%s.%s", obj.gr, obj.Value())
}

// Code returns the code
func (obj *rawToken) Code() string {
	return obj.code
}

// Index returns the index
func (obj *rawToken) Index() int {
	return obj.index
}

// Grammar returns the grammar, if any
func (obj *rawToken) Grammar() string {
	return obj.gr
}
