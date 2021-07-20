package main

import (
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

func main() {
	timeStart := time.Now()

	_, _ = <-worker(), <-worker()

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
	// выведет 6, если надо 3 то сначала присвоить переменной канал, а потом читать
}
