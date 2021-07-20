package main

import (
	"fmt"
	"time"
)

func readChanel(i chan int) {
	i <- 10
	fmt.Println("func", len(i), &i)
	close(i)
}

func main() {
	var tChanal = make(chan int)
	fmt.Println(len(tChanal), &tChanal)
	// go readChanel(tChanal)
	// time.Sleep(time.Second * 1)

	tChanal <- 10
	// close(tChanal)
	// close(tChanal)
	a := <-tChanal
	fmt.Println("aaaaa")
	time.Sleep(time.Second * 1)

	// close of closed channel
	// close(tChanal)
	// panic send on closed channel
	// tChanal<-123
	fmt.Println(a)
}
