// this example shows how to use memoization in Go
package main

import (
	"bytes"
	"fmt"
	"strings"
)

type MemoizeFunction func(int, ...int) interface{}

var RomanForDecimal MemoizeFunction

func init() {
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	RomanForDecimal = Memoize(func(x int, xs ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("RomanForDecimal() only handles integers [0, 3999]")
		}
		var buffer bytes.Buffer
		for i, decimal := range decimals {
			remainder := x / decimal
			x %= decimal
			if remainder > 0 {
				buffer.WriteString(strings.Repeat(romans[i], remainder))
			}
		}

		return buffer.String()
	})
}

func Memoize(function MemoizeFunction) MemoizeFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			fmt.Printf("return from mem: %*s ", 8, value)
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%*d: %s\n", 2, i, RomanForDecimal(i))
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("%*d: %s\n", 2, i, RomanForDecimal(i))
	}
}
