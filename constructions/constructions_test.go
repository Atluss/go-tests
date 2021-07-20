package constructions

import (
	"fmt"
	"testing"
)

// makeEvenGenerator функция со счётчиком
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// sliceIndex функция высшего порядка
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// TestFunctionCounter проверка счетчика вызова функции
func TestFunctionCounter(t *testing.T) {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

// TestHigherOrderFunc test `Higher-order function`
func TestHigherOrderFunc(t *testing.T) {
	xs := []int{2, 4, 6, 8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Printf("%d, %d, %d, %d",
		sliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
		sliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
		sliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
		sliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }),
	)
}

// TestOuter out by GOTO
func TestOuter(t *testing.T) {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break outer
		}
		println()
	}
	fmt.Println("DONE")
}
