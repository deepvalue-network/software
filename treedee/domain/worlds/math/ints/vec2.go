package ints

import "fmt"

// X returns the value at index 0
func (obj Vec2) X() int {
	return obj[0]
}

// Y returns the value at index 1
func (obj Vec2) Y() int {
	return obj[1]
}

// Compare returns true if the two vectors are the same, false otehrwise
func (obj *Vec2) Compare(input Vec2) bool {
	return obj.X() == input.X() && obj.Y() == input.Y()
}

// String returns the string representation of the vector
func (obj Vec2) String() string {
	return fmt.Sprintf("[x: %d, y: %d]", obj.X(), obj.Y())
}
