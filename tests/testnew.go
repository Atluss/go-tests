package main

import "fmt"

func main() {

	var x int = 4
	var p *int = &x            // указатель получает адрес переменной
	fmt.Println("Address:", p) // значение указателя - адрес переменной x

	// *p = 1

	fmt.Println("Value:", *p) // значение переменной x
}
