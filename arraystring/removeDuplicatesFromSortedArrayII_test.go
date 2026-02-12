package arraystring_test

import (
	"fmt"
	"testing"

	"github.com/vedantwankhade/katas/leetcode/top-interview-150/arraystring"
)

func TestRemoveDuplicatesFromSortedArrayII(t *testing.T) {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums3 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	arraystring.RemoveDuplicatesFromSortedArrayII(nums1)
	arraystring.RemoveDuplicatesFromSortedArrayII(nums2)
	arraystring.RemoveDuplicatesFromSortedArrayII(nums3)

	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)
}
