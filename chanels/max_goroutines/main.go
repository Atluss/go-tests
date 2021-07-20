package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case id, ok := <-ch:
			if ok {
				do(id)
			} else {
				wg.Done()
				fmt.Println("worker done")
				return
			}
		}
	}
}

func do(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("%d done\n", id)
}

func main() {
	ch := make(chan int, 2)

	wg := &sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(ch, wg)
	}

	for id := 0; id < 100; id++ {
		ch <- id
	}
	close(ch)
	wg.Wait()
}
