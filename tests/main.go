package main

import (
	"fmt"
)

type testSctruct struct {
	x int
}

func (test *testSctruct) test() {

}

func main() {

	x := [5]int{1, 3, 5, 6, 8}

	for i, value := range x {
		fmt.Println(i, " -- ", value)
	}
}

func zeroExptionTest() {
	for i := 0; i < 10; i++ {

		fmt.Println("i: ", i, " ", i%2)

	}
}
