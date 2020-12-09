// команда для запуска профилирования
// go tool pprof /tmp/cpuProfile.out
package main

import (
	"fmt"
	"go-sandbox/algorithms/sort"
	"os"
	"runtime/pprof"
)

func main() {

	cpuProfFile, err := os.Create("/tmp/cpuProfile.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cpuProfFile.Close()

	err = pprof.StartCPUProfile(cpuProfFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	array := []int{12, -23, 0, 0, 123, 2, 4, -222, 123123}

	sortBubble := make([]int, len(array), cap(array))
	copy(sortBubble, array)
	sort.BubbleSort(sortBubble)
	fmt.Printf("%+v\n", sortBubble)

	sortSelect := make([]int, len(array), cap(array))
	copy(sortSelect, array)
	sort.SelectionSort(sortSelect)
	fmt.Printf("%+v\n", sortSelect)

	sortPast := make([]int, len(array), cap(array))
	copy(sortPast, array)
	sort.PastSort(sortPast)
	fmt.Printf("%+v\n", sortPast)
}
