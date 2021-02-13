package fl32

import "fmt"

// X returns the value at index 0
func (obj Vec2) X() float32 {
	return obj[0]
}

// Y returns the value at index 1
func (obj Vec2) Y() float32 {
	return obj[1]
}

// String returns the string representation of the vector
func (obj Vec2) String() string {
	return fmt.Sprintf("[x: %f, y: %f]", obj.X(), obj.Y())
}
