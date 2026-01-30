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
