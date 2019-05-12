package main

import (
	"fmt"
	"log"
)

func main() {
	var arMount = [...]string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
		"not Here", "again not", "three not"}
	summer := arMount[5:8]
	endSummber := summer[len(summer)-1:]
	copy(summer, []string{"123", "321", "NEW", "23"})

	fmt.Println("Mounts: ", arMount, len(arMount), cap(arMount))
	fmt.Println("Summer: ", summer, len(summer), cap(summer))
	fmt.Println("End summer: ", endSummber, len(endSummber), cap(endSummber))
	fmt.Println(arMount[1:4], summer, endSummber, len(arMount))
	for range arMount {
		log.Printf("amm")
	}
}
