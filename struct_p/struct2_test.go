package struct_p

import (
	"fmt"
	"testing"
)

type A interface {
	Tprint()
}

type C struct {
	g A
}

func q(x interface{}) {
}

func w(x *interface{}) {
}

func z(x A)  {}
func x(x *A) {}

type B struct{}

func (b *B) Tprint() {
	fmt.Println(b)
}

func TestPointerSt(t *testing.T) {
	s := C{g: &B{}}
	p := &s
	fmt.Printf("%T and %T\n", s, p)
	fmt.Printf("%s and %s\n", s, p)
	q(s)
	//w(s) // error
	q(p)
	//w(p) // error

	bb := B{}
	z(&bb) // pointer because Tprint *B
	//x(&bb) // error
}
