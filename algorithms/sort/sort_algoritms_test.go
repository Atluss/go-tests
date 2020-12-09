// запуск бенчей
// go test -bench=. sort_algoritms.go sort_algoritms_test.go
package sort

import (
	"fmt"
	"testing"
)

var testSlice = []int{12, -23, 0, 10, 123, 2, 4, -222, 123123, -1000}
var testAsc = []int{-1000, -222, -23, 0, 2, 4, 10, 12, 123, 123123}

func TestBubbleSort(t *testing.T) {

	sortAr := make([]int, len(testSlice), cap(testSlice))
	copy(sortAr, testSlice)

	sortAr = BubbleSort(sortAr)
	fmt.Printf("%+v\n", sortAr)
	if len(sortAr) != len(testSlice) {
		t.Error(`len(sortAr) != len(testAsc)`)
	}

	for i := 0; i < len(testAsc); i++ {
		if sortAr[i] != testAsc[i] {
			t.Errorf(`sortAr (%d) != testAsc (%d)`, sortAr[i], testAsc[i])
		}
	}
}

func TestSelectionSort(t *testing.T) {

	sortAr := make([]int, len(testSlice), cap(testSlice))
	copy(sortAr, testSlice)

	sortAr = SelectionSort(testSlice)
	fmt.Printf("%+v\n", sortAr)

	if len(sortAr) != len(testAsc) {
		t.Error(`len(sortAr) != len(testAsc)`)
	}

	for i := 0; i < len(testAsc); i++ {
		if sortAr[i] != testAsc[i] {
			t.Errorf(`sortAr (%d) != testAsc (%d)`, sortAr[i], testAsc[i])
		}
	}
}

func TestPastSort(t *testing.T) {

	sortAr := make([]int, len(testSlice), cap(testSlice))
	copy(sortAr, testSlice)

	sortAr = PastSort(testSlice)
	fmt.Printf("%+v\n", sortAr)

	if len(sortAr) != len(testSlice) {
		t.Error(`len(sortAr) != len(testAsc)`)
	}

	for i := 0; i < len(testAsc); i++ {
		if sortAr[i] != testAsc[i] {
			t.Errorf(`sortAr (%d) != testAsc (%d)`, sortAr[i], testAsc[i])
		}
	}
}

func benchmarkBubbleSort(b *testing.B, arr []int) {

	sortAr := make([]int, len(arr), cap(arr))

	for i := 0; i < b.N; i++ {
		copy(sortAr, arr)
		BubbleSort(sortAr)
	}
}

func benchmarkSelectionSort(b *testing.B, arr []int) {

	sortAr := make([]int, len(arr), cap(arr))

	for i := 0; i < b.N; i++ {
		copy(sortAr, arr)
		SelectionSort(sortAr)
	}
}

func benchmarkPastSort(b *testing.B, arr []int) {

	sortAr := make([]int, len(arr), cap(arr))

	for i := 0; i < b.N; i++ {
		copy(sortAr, arr)
		PastSort(sortAr)
	}
}

func BenchmarkBubbleSortArrLen10(b *testing.B) {
	benchmarkBubbleSort(b, testSlice)
}

func BenchmarkSelectionSortArrLen10(b *testing.B) {
	benchmarkSelectionSort(b, testSlice)
}

func BenchmarkPastSortArrLen10(b *testing.B) {
	benchmarkPastSort(b, testSlice)
}
