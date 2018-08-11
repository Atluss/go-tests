package main

import (
	"fmt"
	"math"
)

// Circle is a struct
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.x * c.x
}

func main() {

	// c := Circle{0.0, 0.1, 5.0}

	c := new(Circle)
	c.x = 5.0
	fmt.Println(c.area())

	setMap()
	printMap()

}
