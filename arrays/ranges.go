package main

import "fmt"

func main() {
	z := [5]float64{98, 93, 77, 82, 83}

	x := make([]float64, 5, 10)

	x = append(x, 0.1)

	copy(x, z[1:3])

	fmt.Println(x)
}
