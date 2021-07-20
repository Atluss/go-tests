package struct_p

import (
	"fmt"
	"reflect"
	"testing"
)

type S struct {
	m string
}

func f() *S {
	return &S{"foo"} // A
}

func TestCStruct(t *testing.T) {
	p := f() // B
	fmt.Printf("%T\n", p)
	fmt.Println(p.m) // print "foo"
}

func TestCompareInterface(t *testing.T) {
	x := interface{}(&S{m: "lol"})
	y := interface{}(&S{m: "lol"})

	z := x

	fmt.Println(x == y)                    // false
	fmt.Println(&x == &y)                  // false
	fmt.Println(reflect.DeepEqual(x, y))   // true
	fmt.Println(reflect.DeepEqual(&x, &y)) // true

	fmt.Println(z == x, z, x)
}

func TestEmptyInterface(t *testing.T) {
	var a interface{}
	var b interface{}
	fmt.Println(a == b)
	fmt.Println(a, a == nil)
	a = "a"
	fmt.Println(a, a == nil)
}
