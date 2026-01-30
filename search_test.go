package dsa

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(BinarySearch(1, list, func(t1, t2 int) int {
		return t1 - t2
	}))
	fmt.Println(BinarySearch(2, list, func(t1, t2 int) int {
		return t1 - t2
	}))
	fmt.Println(BinarySearch(5, list, func(t1, t2 int) int {
		return t1 - t2
	}))
	fmt.Println(BinarySearch(6, list, func(t1, t2 int) int {
		return t1 - t2
	}))
	fmt.Println(BinarySearch(99, list, func(t1, t2 int) int {
		return t1 - t2
	}))
}
