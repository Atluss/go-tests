package main

import (
	"fmt"
	"strings"
)

func NewPart(id int, name string) *Part {
	return &Part{id, name}
}

type Part struct {
	Id   int    // агрегирование
	Name string // агрегирование
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}
func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}
func (part Part) String() string {
	return fmt.Sprintf("%d --- %q", part.Id, part.Name)
}
func (part *Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

type Count int

func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

func main() {

	tom := NewPart(33, "Tom")
	fmt.Println(tom)

	part := Part{5, "wrench"}
	part.UpperCase()
	part.Id += 11
	fmt.Println(part, part.HasPrefix("w"))

	var count Count
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	k := int(count)
	fmt.Println(count, i, j, k, count.IsZero())

}
