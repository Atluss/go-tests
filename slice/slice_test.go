package slice

import (
	"fmt"
	"log"
	"testing"
)

func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

func TestCopySlice1(t *testing.T) {
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

// TestSlicePointers данный пример показывает что когда присваивается диапозон из другого слайса то заполняются адреса памяти слайса
func TestSlicePointers(t *testing.T) {
	s := []int{1, 2, 3}
	ss := s[1:]
	s[0] = 12

	fmt.Println(ss)
	fmt.Println(s)

	fmt.Println(&ss[0], &s[1])

	fmt.Println(len(ss), cap(ss))
}
