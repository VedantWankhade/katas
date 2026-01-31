package dsa

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 0}
	fmt.Println(list)
	fmt.Println(BubbleSort(list))
}

func TestSelectionSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 0}
	fmt.Println(list)
	fmt.Println(SelectionSort(list))
}

func TestInsertionSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 0}
	fmt.Println(list)
	fmt.Println(InsertionSort(list))
}

func TestHeapSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 0}
	fmt.Println(list)
	fmt.Println(HeapSort(list))
}

func TestMergeSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 0}
	fmt.Println(list)
	fmt.Println(MergeSort(list))
}

func TestQuickSort(t *testing.T) {
	list := []int{2, 4, 1, 2, 7, 3}
	fmt.Println(list)
	fmt.Println(QuickSort(list))
}
