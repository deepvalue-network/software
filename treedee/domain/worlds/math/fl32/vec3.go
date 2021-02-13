package fl32

import "fmt"

// X returns the value at index 0
func (obj Vec3) X() float32 {
	return obj[0]
}

// Y returns the value at index 1
func (obj Vec3) Y() float32 {
	return obj[1]
}

// Z returns the value at index 2
func (obj Vec3) Z() float32 {
	return obj[2]
}

// String returns the string representation of the vector
func (obj Vec3) String() string {
	return fmt.Sprintf("[x: %f, y: %f, z: %f]", obj[0], obj[1], obj[2])
}
