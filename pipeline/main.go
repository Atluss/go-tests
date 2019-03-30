// example how to work with pipelines
package main

import (
	"fmt"
)

func main() {

	nat := make(chan int)
	sqr := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			nat <- x
		}
		close(nat)
	}()

	go func() {
		for x := range nat {
			//fmt.Println("read from 1")
			sqr <- x * x
		}
		close(sqr)
	}()

	for x := range sqr {
		fmt.Println(x)
	}

}
