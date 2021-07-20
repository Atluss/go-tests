package struct_p

import (
	"testing"
)

type Count int

func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

func TestChangeByPointer(t *testing.T) {
	var count Count
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	k := int(count)
	t.Log(count, i, j, k, count.IsZero())
}
