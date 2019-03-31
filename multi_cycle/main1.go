// example how to use gorutines and channels for parallel computing
package main

import (
	"log"
	"sync"
)

func main() {

	sums := []uint{1, 2, 4, 6, 6}

	log.Println(multiCalc(sums))

}

func multiCalc(sumc []uint) (total uint) {

	sizes := make(chan uint)
	var wg sync.WaitGroup

	for _, f := range sumc {
		// add wait
		wg.Add(1)

		// calculate something
		go func(f uint) {
			// don't forget reduce wg
			defer wg.Done()
			sizes <- f * 2
		}(f)
	}

	// wait sums, and close channel
	go func() {
		wg.Wait()
		close(sizes)
	}()

	// read from chanel
	for size := range sizes {
		total += size
	}

	return total
}
