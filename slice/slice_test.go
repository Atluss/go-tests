package slice

import (
	"fmt"
	"log"
	"reflect"
	"sort"
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

// TestSlicePointers данный пример показывает что когда присваивается диапазон из другого слайса то заполняются адреса памяти слайса
func TestSlicePointers(t *testing.T) {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss[0] = 1233
	s[0] = 12

	fmt.Println(ss)
	fmt.Println(s)

	fmt.Println(&ss[0], &s[1])

	fmt.Println(len(ss), cap(ss))
}

func TestSliceCompare(t *testing.T) {
	a := []int{1}
	b := []int{1}
	//fmt.Println(a == b) NO!
	fmt.Println(reflect.DeepEqual(a, b))
}

func TestSliceSortReverse(t *testing.T) {
	a := []int{12, -23, 23, 234, 23}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
}

// TestRef answer: [3 4] [3 4] [1 2]
func TestRef(t *testing.T) {
	x := []int{1, 2}
	y := []int{3, 4}
	ref := x
	x = y
	fmt.Println(x, y, ref)
}

// TestCopySlice [4] 1 2 (a[3:4:5] 5 - for cap)
func TestCopySlice(t *testing.T) {
	a := [6]int{1, 2, 3, 4, 5, 6}
	c := a[3:4:5]
	fmt.Println(c, len(c), cap(c))
}

func TestReverseSlice(t *testing.T) {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}

	l := len(a)

	for i := 1; i < l/2; i++ {
		a[i-1], a[l-i] = a[l-i], a[i-1]
	}
	fmt.Println(a)

	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}

func TestSliceChangeArr(t *testing.T) {
	a := [6]int{1, 2, 3, 4, 5, 6}
	b := a[1:3]
	b = append(b, 234)
	b[0] = 123
	fmt.Println(a, b)
}

// TestSliceMemoryLen примере выделения памяти в слайсе
func TestSliceMemoryLen(t *testing.T) {
	s := make([]int, 1)
	s = append(s, 12, 23)
	t.Logf("len: %d, cap: %d, st: %+v", len(s), cap(s), s)
}
