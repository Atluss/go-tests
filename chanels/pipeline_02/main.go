// previous example how to work pipelines, but in func
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func sqr(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	nat := make(chan int)
	sr := make(chan int)

	go counter(nat)
	go sqr(sr, nat)
	printer(sr)
}
