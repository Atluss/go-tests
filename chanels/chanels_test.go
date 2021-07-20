package chanels

import (
	"fmt"
	"runtime"
	"testing"
)

func TestChnStr(t *testing.T) {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")

}
