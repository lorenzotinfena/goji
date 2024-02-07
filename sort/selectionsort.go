package sort

func SelectionSort[T any](slice []T, prior func(a, b T) bool) {
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if prior(slice[j], slice[min]) {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}
}
