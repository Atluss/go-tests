package struct_p

import (
	"fmt"
	"testing"
)

type StructT struct{}

func (s StructT) prt() {
	fmt.Println("LOL")
}

type StructA struct {
	s StructT
}

func (s StructA) prt() {
	fmt.Println("LOL2")
}

func TestStructFunc(t *testing.T) {
	var s StructA
	s.prt()
	s.s.prt()

	var ar []int
	fmt.Println(ar, ar == nil)
}
