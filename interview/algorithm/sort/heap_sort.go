package sort

func HeapSort(data []int) {
	n := len(data)
	// 1. 维护堆有序
	// 2. 交换堆顶元素， 下沉， 继续维护堆有序
	for i := n / 2; i >= 1; i-- {
		sink(data, i, n)
	}
	scope := n
	for scope > 1 {
		swap(data, 1, scope)
		scope--
		sink(data, 1, scope)
	}
}

func sink(data []int, k int, n int) {
	for 2*k <= n {
		i := 2 * k
		if i < n && less(data, i, i+1) {
			i = i + 1
		}
		if !less(data, k, i) {
			break
		}
		swap(data, k, i)
		k = i
	}
}

func less(data []int, a int, b int) bool {
	return data[a-1] < data[b-1]
}

func swap(data []int, a int, b int) {
	data[a-1], data[b-1] = data[b-1], data[a-1]
}
