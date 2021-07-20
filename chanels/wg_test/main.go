package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-done:
		return

	default:
		fmt.Println("Begin")
		time.Sleep(2 * time.Second)
		fmt.Println("End")

	}
}

func main() {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go worker(done, wg)

	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()
	fmt.Println("Ok")
}
