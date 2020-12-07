package maps

import (
	"fmt"
	"testing"
)

func TestMapMap(t *testing.T) {
	elements := map[string]map[string]map[string]string{
		"Helium": {
			"gas": {
				"A": "asd",
			},
		},
		"Lithium": {
			"solid": {
				"A": "asd",
			},
		},
		"Beryllium": {
			"solid": {
				"A": "asd",
			},
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	for i, value := range elements {
		fmt.Println("Short code: ", i)
		fmt.Println("\t value: ", value)
	}
}
