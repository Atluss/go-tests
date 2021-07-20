package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapMap(t *testing.T) {
	elements := map[string]map[string]map[string]string{
		"Helium": {
			"gas": {
				"A": "asd",
			},
		},
		"Lithium": {
			"solid": {
				"A": "asd",
			},
		},
		"Beryllium": {
			"solid": {
				"A": "asd",
			},
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	for i, value := range elements {
		fmt.Println("Short code: ", i)
		fmt.Println("\t value: ", value)
	}
}

func TestCompareMap(t *testing.T) {
	a := map[string]string{"A": "B"}
	b := map[string]string{"A": "B"}
	c := map[string]func(){"A": func() {}}
	d := map[string]interface{}{"A": func() {}}

	c["C"] = func() {
		fmt.Println("ASD")
	}
	if f, ok := c["C"]; ok {
		f()
	}

	if f, ok := d["A"]; ok {
		if ff, ok := f.(func()); ok {
			ff()
		}
	}

	//fmt.Println(a == b) // error
	fmt.Println(reflect.DeepEqual(a, b))
	fmt.Println(reflect.DeepEqual(c, d))

	e := map[string][]int{"A": {1, 2, 3}, "B": {1, 2, 3}, "C": {1, 2, 3}}
	f := map[string][]int{"A": {1, 2, 3}, "B": {1, 2, 3}, "C": {1, 2, 3}}
	//fmt.Println(e == f) // error
	fmt.Println(reflect.DeepEqual(e, f))
}

func TestGetMapValue(t *testing.T) {
	a := map[string]int{
		"A": 1,
	}

	fmt.Printf("%T, %d\n", a["A"], a["A"])

	if c, ok := a["A"]; ok {
		fmt.Println(c)
		fmt.Printf("%T, %d\n", c, c)
	}
}

func TestEmptyInterfaceKey(t *testing.T) {
	m := map[interface{}]int{}

	var a interface{}
	var b interface{}

	a = 2
	b = "A"

	m[a] = 1
	m[b] = 2

	fmt.Println(m, m[2], m["A"])
	for k, v := range m {
		switch k.(type) {
		case string:
			fmt.Printf("%T, %s, value: %d\n", k, k, v)
		case int:
			fmt.Printf("%T, %d, value: %d\n", k, k, v)
		}
		//fmt.Printf("%T, %d, value: %d\n", k, k, v)
	}
}

func TestCreateMap(t *testing.T) {
	var m map[string]int
	m["a"] = 1
	if v, ok := m["b"]; ok { //B
		println(v)
	}
}

func TestChangeStructInMap(t *testing.T) {
	type S struct {
		name string
	}
	// если нужно изменять значение через ключ и точку то делаем так:
	mPtr := map[string]*S{"x": &S{"one"}}
	mPtr["x"].name = "two" // change by point get

	// обычная работа с map
	m := map[string]S{"x": S{"one"}}
	t.Log(m["x"].name)
	s := m["x"]
	s.name = "two" // значение у map не поменяется
	t.Logf("%+v", m)
}
