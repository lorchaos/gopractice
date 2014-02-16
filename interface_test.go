package gopractice

import (
	"fmt"
	"testing"
)

func ExampleBinaryOperation() {

	b := BinaryFunc(Sum)

	fmt.Printf("%d", b(5, 6))

	//The output is compared to the next line
	//Output: 11
}

func TestSubtraction(t *testing.T) {

	b := BinaryFunc(Sub)

	if b(10, 4) != 6 {
		t.Error("Wrong result")
	}

	t.Log(b.Debug())
}

func TestSum(t *testing.T) {

	b := BinaryFunc(Sum)

	if b(10, 4) != 14 {
		t.Error("Wrong result")
	}

	t.Log(b.Debug())
}
