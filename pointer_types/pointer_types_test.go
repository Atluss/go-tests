package pointer_types

import (
	"fmt"
	"testing"
)

var patternStrInit = "Type: %T, len: %d, init value: %+v\n"
var patternStrAppend = "After append: len: %d, %+v\n------------------------\n\n"
var i = 1

// TestSlice1 differences from make, new and create var
func TestSlice1(t *testing.T) {

	fmt.Println("Create value: []int")
	var st1 []int
	fmt.Printf(patternStrInit, st1, len(st1), st1)
	st1 = append(st1, i)
	fmt.Printf(patternStrAppend, len(st1), st1)

	fmt.Println("Create value: []*int")
	var st2 []*int
	fmt.Printf(patternStrInit, st2, len(st2), st2)
	st2 = append(st2, &i)
	fmt.Printf(patternStrAppend, len(st2), st2)

	fmt.Println("Create value: := &[]*int{}")
	st3 := &[]*int{}
	fmt.Printf(patternStrInit, st3, len(*st3), st3)
	*st3 = append(*st3, &i)
	fmt.Printf(patternStrAppend, len(*st3), st3)

	// value
	fmt.Println("Make value: make([]int, 1))")
	var sm1 = make([]int, 1)
	fmt.Printf(patternStrInit, sm1, len(sm1), sm1)
	sm1 = append(sm1, i)
	fmt.Printf(patternStrAppend, len(sm1), sm1)

	fmt.Println("Make value: make([]*int, 1)")
	var sm2 = make([]*int, 1)
	fmt.Printf(patternStrInit, sm2, len(sm2), sm2)
	sm2 = append(sm2, &i)
	fmt.Printf(patternStrAppend, len(sm2), sm2)

	// pointer
	fmt.Println("New pointer: new([]int)")
	var s1 = new([]int)
	fmt.Printf(patternStrInit, s1, len(*s1), s1)
	*s1 = append(*s1, i)
	fmt.Printf(patternStrAppend, len(*s1), s1)

	fmt.Println("New pointer pointer: new([]*int)")
	var s2 = new([]*int)
	fmt.Printf(patternStrInit, s2, len(*s2), s2)
	*s2 = append(*s2, &i)
	fmt.Printf(patternStrAppend, len(*s2), s2)

	// High not recommended use this!!!
	fmt.Println("New pointer pointer: new(*[]*int)")
	var s3 = new(*[]*int)
	ts := &[]*int{}
	s3 = &ts
	fmt.Printf(patternStrInit, s3, len(**s3), s3)
	**s3 = append(**s3, &i)
	fmt.Printf(patternStrAppend, len(**s3), s3)

}
