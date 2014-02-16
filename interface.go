package gopractice

import (
	"fmt"
	"reflect"
)

type BinaryFunc func(int, int) int

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

//adding a method to a function!
func (b BinaryFunc) Debug() string {

	return fmt.Sprintf("Debugging %s", reflect.TypeOf(&b).Name())
}
