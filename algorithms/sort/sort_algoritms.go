package sort

// BubbleSort сортировка пузырьком
func BubbleSort(ar []int) []int {
	length := len(ar)

	for i := 1; i < length; i++ {
		for k := 0; k < length-i; k++ {

			if ar[k] > ar[k+1] {
				ar[k], ar[k+1] = ar[k+1], ar[k]
			}
		}
	}
	return ar
}

// SelectionSort сортировка выборкой
func SelectionSort(ar []int) []int {
	length := len(ar)
	nom := 0
	max := 0

	for j := 1; j <= length; j++ {
		max = ar[0]
		nom = 0
		for i := 1; i <= length-j; i++ {
			if ar[i] > max {
				max = ar[i]
				nom = i
			}
		}
		ar[nom], ar[length-j] = ar[length-j], ar[nom]
	}

	return ar
}

// InsertSort сортировка вставкой
func InsertSort(ar []int) []int {

	for i := 1; i < len(ar); i++ {

		current := ar[i]
		j := i - 1

		for j >= 0 && ar[j] > current {
			ar[j+1] = ar[j]
			j--
		}

		ar[j+1] = current
	}

	return ar
}

func MergeSort(in []int) []int {

	if len(in) <= 1 {
		return in
	}

	middleIndex := len(in) / 2

	left := in[:middleIndex]
	right := in[middleIndex:]

	leftSorted := MergeSort(left)
	rightSorted := MergeSort(right)

	return compareMergeSort(leftSorted, rightSorted)
}

func compareMergeSort(left, right []int) (out []int) {
	i, j := 0, 0

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}

	if i < len(left) {
		out = append(out, left[i:]...)
	}

	if j < len(right) {
		out = append(out, right[j:]...)
	}

	return out
}
