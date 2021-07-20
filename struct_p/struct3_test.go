package struct_p

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo struct {
	A int
	B string
	C interface{}
}

func (f *Foo) a() {}

type Boo struct {
	A int
	B string
	C interface{}
}

func TestCompareStruct(t *testing.T) {
	a := Foo{A: 1, B: "one", C: "two"}
	b := Boo{A: 1, B: "one", C: "two"}

	c := Foo{A: 1, B: "one", C: "two"}

	//fmt.Println(a == b) // error
	fmt.Println(a == c)
	fmt.Println(reflect.DeepEqual(a, b))
}
