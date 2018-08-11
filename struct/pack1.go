package main

import (
	"fmt"
)

var maps = make(map[string]string)

// var mapMaps = make(map[string]map[string]map[string]string)
var mapMaps = make(map[string]map[string]map[string]string)

func main() {
	setMap()
	printMap()
}

func printIt() {
	fmt.Println("1")
}

func setMap() {

	maps["test"] = "test"
	maps["test1"] = "test"
	maps["test2"] = "test"
	maps["test3"] = "test"
	maps["test4"] = "test"
	maps["test5"] = "test"

	mapMaps["test"] = map[string]map[string]string{
		"HE": map[string]string{
			"no": "bo",
		},
	}

}

func printMap() {

	// for i, value := range maps {
	// 	fmt.Println(i, " ", value)
	// }

	for i, value := range mapMaps {
		fmt.Println(i)
		for k, valueE := range value {
			fmt.Println(k, " ")
			for key, val := range valueE {
				fmt.Println(key, "  ", val)
			}
		}
	}
}
