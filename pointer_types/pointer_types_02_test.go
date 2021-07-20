package pointer_types

import (
	"fmt"
	"testing"
)

type ALOL struct {
	g string
}

type TestStruct struct {
	a ALOL
	b bool
}

func (t *TestStruct) ChangeB() {
	t.b = true
}

func TestChangeNoP(t *testing.T) {

	st := TestStruct{
		b: false,
	}

	st.ChangeB()

	fmt.Println(st)

}

func TestDeclare(t *testing.T) {

	var st TestStruct
	//var st = new(TestStruct)

	fmt.Println(st.a)
	fmt.Printf("%+v\n", st)

	var i int
	fmt.Printf("%+v\n", i)

	i2 := 123
	fmt.Printf("%+v\n", i2)

	var i3 = 321
	fmt.Printf("%+v\n", i3)

	//var i4 *int
	//*i4 = 123
	//fmt.Printf("%+v\n", i4)
}

func TestCompareNil(t *testing.T) {
	var s []int
	fmt.Println(s == nil)

	m := map[int]int{1: 1}
	m[2] = 2
	fmt.Println(m == nil, m)

	var c chan int
	fmt.Println(c == nil)
}
