package sort

func InsertionSort(data []int) {
	n := len(data)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}
