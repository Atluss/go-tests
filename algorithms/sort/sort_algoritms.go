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

// PastSort сортировка вставкой
func PastSort(ar []int) []int {

	length := len(ar)
	b := ar[0]
	j := 0

	for i := 1; i < length; i++ {

		b = ar[i]

		for j = i - 1; j > -1 && b < ar[j]; j-- {
			ar[j+1] = ar[j]
		}

		ar[j+1] = b
	}

	return ar
}
