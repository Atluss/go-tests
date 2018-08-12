package main

import (
	"encoding/json"
	"fmt"
)

// тестовая структура
type objStruct struct {
	X, Y float64 // координаты
	Name string  `json:"structName"` // имя
}

// передвижение объекта
func (obj *objStruct) move(lenX, lenY float64) {
	obj.X += lenX
	obj.X += lenY
}

// метод имя
func (obj *objStruct) sayName() {
	fmt.Println(obj.Name)
}

// json объекта
func (obj objStruct) toJSON() string {

	a := &objStruct{obj.X, obj.Y, obj.Name}
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	// fmt.Println(out)

	return string(out)
}

func main() {

	mover := objStruct{X: 10.0, Y: 10.0}

	mover.Name = "Jack"
	mover.move(2.0, 1)

	jsonStr := mover.toJSON()
	fmt.Println(jsonStr)

	// fmt.Println(string(mover.toJSON()))

	// for field, value := range mover {
	// 	fmt.Println(field, ":  ", value)
	// }

}
