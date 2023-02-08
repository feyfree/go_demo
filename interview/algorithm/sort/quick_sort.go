package sort

func QuickSort(data []int) {
	sort0(data, 0, len(data)-1)
}

func sort0(data []int, l int, h int) {
	if l >= h {
		return
	}
	m := l
	for i := l + 1; i <= h; i++ {
		if data[i] < data[l] {
			m++
			data[m], data[i] = data[i], data[m]
		}
	}
	data[m], data[l] = data[l], data[m]
	sort0(data, l, m-1)
	sort0(data, m+1, h)
}
