package sort

import (
	"fmt"
	"testing"
)

var data = []int{1, 9, 7, 8, 5, 11, 2, 0, 33, 3, 4, 6}

func TestSelectionSort(t *testing.T) {
	SelectionSort(data)
	fmt.Println(data)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(data)
	fmt.Println(data)
}

func TestQuickSort(t *testing.T) {
	QuickSort(data)
	fmt.Println(data)
}

func TestBinarySearch(t *testing.T) {

	fmt.Println(BinarySearch([]int{1, 2}, 2))
}

func TestMergeSortTopDown(t *testing.T) {
	MergeSortTopDown(data)
	fmt.Println(data)
}

func TestMergeSortBottomUp(t *testing.T) {
	MergeSortBottomUp(data)
	fmt.Println(data)
}

func TestHeapSort(t *testing.T) {
	HeapSort(data)
	fmt.Println(data)
}
