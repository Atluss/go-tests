package err

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type MyStructure struct {
	Message string
}

func (e MyStructure) Error() string {
	return e.Message
}

func Execute() error {
	var result *MyStructure // создаёт указатель и не выделяет помять под значение значит будет равен nil

	if result == nil {
		fmt.Println("TRUE", result)
	}

	return result // приводим к интерфейсу error тип: MyStructure, но значение нил, значит не нил
}

func TestEmptyError(t *testing.T) {
	err := Execute()

	if err == nil {
		fmt.Println("PRINT ?", err)
	}

	if errors.As(err, &MyStructure{}) {
		fmt.Println("LOL")
	}
	if errors.Is(err, errors.New("")) {
		fmt.Println("LOL2")
	}
	if err == errors.New("") {
		fmt.Println("LOL3")
	}
}

func TestUnwrap(t *testing.T) {
	err := errors.New("my error")
	err = fmt.Errorf("1s wrapping my error with Errorf: %w", err)
	err = fmt.Errorf("2nd wrapping my error with Errorf: %w", err)

	ce := errors.Unwrap(err)

	fmt.Println(ce)
	fmt.Println(errors.Unwrap(ce))

	er1 := errors.Unwrap(ce)
	fmt.Println(errors.Unwrap(er1))

	fmt.Println(err)
}

func TestAsError(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func TestIsError(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}
