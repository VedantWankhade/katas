package katas

/*
Question:

https://leetcode.com/problems/intersection-of-two-arrays-ii/description/?envType=problem-list-v2&envId=two-pointers

	Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.

Constraints:

	1 <= nums1.length, nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 1000

Follow up:

	What if the given array is already sorted? How would you optimize your algorithm?
	What if nums1's size is small compared to nums2's size? Which algorithm is better?
	What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

// hashmap approach
func ArrayIntersection1(nums1, nums2 []int) []int {
	m := make(map[int][]int)
	for _, n := range nums1 {
		if _, ok := m[n]; !ok {
			m[n] = make([]int, 2)
		}
		m[n][0]++
	}

	for _, n := range nums2 {
		if _, ok := m[n]; !ok {
			m[n] = make([]int, 2)
		}
		m[n][1]++
	}

	ans := make([]int, 0)

	for k, v := range m {
		times := min(v[0], v[1])
		for range times {
			ans = append(ans, k)
		}
	}

	return ans
}

// hasmap approach optimized
func ArrayIntersection2(nums1, nums2 []int) []int {
	nums1Len, nums2Len := len(nums1), len(nums2)
	// make sure nums1 is the smaller one (to optimizse hashmap size)
	if nums1Len > nums2Len {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]int)
	for _, n := range nums1 {
		m[n]++
	}

	ans := make([]int, 0)

	for _, n := range nums2 {
		if nn := m[n]; nn != 0 {
			ans = append(ans, n)
			m[n]--
		}
	}

	return ans
}
