package sort

func SelectionSort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}
