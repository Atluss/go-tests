package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 1
	if value := reflect.ValueOf(number); value.CanSet() { // try set
		value.SetInt(2) // panic
	}
	fmt.Printf("Base number: %d\n", number)

	value := reflect.ValueOf(&number)
	fmt.Printf("Type: %v Val: %+v\n", reflect.TypeOf(number), value.Elem())
	ptr := value.Elem()
	ptr.SetInt(3) // OK. Значение, на которое ссылается указатель, заменить можно.

	fmt.Printf("After change: %d\n", number)
	fmt.Println("----SELECT---")

	aim := []interface{}{0, "asd", 2.44, true, 123}
	for i := range aim {
		fmt.Println("===================")
		fmt.Printf("%v\n", reflect.TypeOf(aim[i]))
		switch aim[i].(type) { // shadow variable
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		case int:
			fmt.Println("int")
		case bool:
			fmt.Println("bool")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("Unknow type: %v\n", reflect.TypeOf(aim[i]))
		}
	}
}
