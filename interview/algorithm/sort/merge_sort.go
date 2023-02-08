package sort

import "math"

// MergeSortBottomUp merge sort 主要是 merge, 合并两个有序数组
func MergeSortBottomUp(data []int) {
	n := len(data)
	// 辅助数组
	aux := make([]int, n)
	for step := 1; step < n; step = step * 2 {
		for lo := 0; lo+step < n; lo = lo + step*2 {
			mid := lo + step - 1
			hi := int(math.Min(float64(lo+2*step-1), float64(n-1)))
			merge(data, aux, lo, mid, hi)
		}
	}
}

func MergeSortTopDown(data []int) {
	sortTd(data, make([]int, len(data)), 0, len(data)-1)
}

func sortTd(data []int, aux []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	sortTd(data, aux, lo, mid)
	sortTd(data, aux, mid+1, hi)
	merge(data, aux, lo, mid, hi)
}

func merge(data []int, aux []int, lo int, mid int, hi int) {
	// 拷贝到辅助集合
	if hi-lo+1 > 0 {
		for i := lo; i <= hi; i++ {
			aux[i] = data[i]
		}
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i <= mid && j <= hi {
			if aux[i] > aux[j] {
				data[k] = aux[j]
				j++
			} else {
				data[k] = aux[i]
				i++
			}
		} else if i > mid {
			data[k] = aux[j]
			j++
		} else {
			data[k] = aux[i]
			i++
		}
	}
}
