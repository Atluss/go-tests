// example how to use multi cycle in one func
package main

import "fmt"

func main() {
	strs := []string{"one", "two", "three", "four"}
	var err error

	strs, err = doSomething(strs)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range strs {
		fmt.Println(v)
	}
}

func doSomething(strs []string) (nstr []string, err error) {
	type item struct {
		mes string
		err error
	}

	// make channel for read results
	ch := make(chan item, len(strs))

	for _, f := range strs {
		go func(f string) {
			var it item
			it.mes, it.err = newStr(f)
			// send result to channel
			ch <- it
		}(f)
	}

	// read channel and return err or calculate result
	for range strs {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		nstr = append(nstr, it.mes)
	}

	return nstr, nil
}

func newStr(str string) (string, error) {
	return str + "_NEW", nil
}
