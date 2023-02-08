package sort

func BinarySearch(data []int, target int) int {
	start, end := 0, len(data)
	for start < end {
		mid := (start + end) / 2
		if data[mid] == target {
			return mid
		} else if data[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}
