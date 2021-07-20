package read_write_close_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 42
	}()

	return ch
}

// TestWorker wait 6 sec
func TestWorker(t *testing.T) {
	_, _ = <-worker(), <-worker()
}

func TestPipeline(t *testing.T) {
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

// TestDeadLockNoBuf1 - Error: fatal error: all goroutines are asleep - deadlock!
func TestDeadLockNoBuf1(t *testing.T) {
	ch := make(chan int)
	i := <-ch
	fmt.Println("Done", i)
}

// TestDeadLockNoBuf2 - Error: fatal error: all goroutines are asleep - deadlock!
func TestDeadLockNoBuf2(t *testing.T) {
	ch := make(chan int)
	ch <- 2
	fmt.Println("Done")
}

// TestCloseCloseCh - Error: panic: close of closed channel [recovered]
func TestCloseCloseCh(t *testing.T) {
	ch := make(chan int, 1)
	close(ch)
	close(ch)
	fmt.Println("Done")
}

// TestWriteToCloseCh - Error: panic: send on closed channel [recovered]
func TestWriteToCloseCh(t *testing.T) {
	ch := make(chan int, 1)
	close(ch)
	ch <- 1
	fmt.Println("Done")
}

// TestReadFromCloseCh - OK
func TestReadFromCloseCh(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	time.Sleep(time.Second)
	fmt.Println("Done", <-ch)
}

func TestCompareChannels(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 1

	fmt.Println(ch1 == ch2)                  // false
	fmt.Println(reflect.DeepEqual(ch1, ch2)) // false
}

func TestDeclareChan(t *testing.T) {
	ch := make(chan int)
	ch <- 20

	fmt.Println("Done", <-ch)
}
