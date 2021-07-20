package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("A: ", e)
		}
	}()
	pan1()
	fmt.Println("It is write")
}

func pan1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("A: ", e) // recover this panic
		}
	}()
	pan2()
}

func pan2() {
	panic("ASD")
}
