package hash

import (
	"bytes"
	"fmt"
)

// Compare returns true if the current hash is the same as the passed one
func (obj Hash) Compare(other Hash) bool {
	currentBytes := []byte{}
	for _, oneByte := range obj {
		currentBytes = append(currentBytes, oneByte)
	}

	otherBytes := []byte{}
	for _, oneByte := range other {
		otherBytes = append(otherBytes, oneByte)
	}

	return bytes.Compare(currentBytes, otherBytes) == 0
}

// Bytes returns the bytes of an hash
func (obj Hash) Bytes() []byte {
	out := []byte{}
	for _, oneByte := range obj {
		out = append(out, oneByte)
	}

	return out
}

// String returns the string of an hash
func (obj Hash) String() string {
	bytes := obj.Bytes()
	return fmt.Sprintf("%x", bytes)
}
