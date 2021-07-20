package _interface

import (
	"fmt"
	"testing"
)

func TestComInterface(t *testing.T) {
	var a interface{}
	var b interface{}

	a = 2
	b = "asd"

	fmt.Println(a == b)
}
