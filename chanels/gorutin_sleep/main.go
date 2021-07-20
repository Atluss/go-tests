package main

import (
	"fmt"
	"time"
)

func testGoRoutin() {
	go func() {
		for {
			fmt.Println("tik tak")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("END FUNC")
	return
}

func main() {
	go testGoRoutin()
	time.Sleep(time.Second * 10)
}
