package err

import (
	"errors"
	"fmt"
	"testing"
)

var ErrorCustom = errors.New("a error")

type MyStructureErr struct{}

func (e MyStructureErr) Error() string {
	return ErrorCustom.Error()
}

func NewMyStructureErr() error {
	//var result MyStructureErr

	result := MyStructureErr{}
	return result
}

func TestErrorInterface(t *testing.T) {
	err := NewMyStructureErr()

	if errors.As(err, &MyStructureErr{}) {
		fmt.Println("LOL")
	}
	if errors.Is(err, ErrorCustom) {
		fmt.Println("LOL2")
	}
	if err == ErrorCustom {
		fmt.Println("LOL3")
	}
}
