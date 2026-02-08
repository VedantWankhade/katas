package arraystring_test

import (
	"fmt"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	arraystring.RemoveDuplicatesFromSortedArray(nums1)
	arraystring.RemoveDuplicatesFromSortedArray(nums2)

	fmt.Println(nums1)
	fmt.Println(nums2)
}

func TestRemoveDuplicatesFromSortedArrayFast(t *testing.T) {
	nums1 := []int{1, 1, 2}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	arraystring.RemoveDuplicatesFromSortedArrayFast(nums1)
	arraystring.RemoveDuplicatesFromSortedArrayFast(nums2)

	fmt.Println(nums1)
	fmt.Println(nums2)
}
