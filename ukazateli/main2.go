package main

import (
	"fmt"
)

func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)

	var x int = *xPtr

	fmt.Println(x) // x is 1
}
