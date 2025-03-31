package de.vedantwankha.kata.algorithms.techniques.arrays.twopointers;

/**
 * <h1>Sort Colors</h1>
 * <a href="https://leetcode.com/problems/sort-colors/description/">https://leetcode.com/problems/sort-colors/description/</a>
 *
 * <p>Given an array nums with <code>n</code> objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.</p>
 *
 * <p>We will use the integers <code>0</code>, <code>1</code>, and <code>2</code> to represent the color red, white, and blue, respectively.</p>
 *
 * <p>You must solve this problem without using the library's sort function.</p>
 *
 *
 *
 * <h3>Example 1:</h3>
 *
 * <p>Input: <code>nums = [2,0,2,1,1,0]</code></p>
 * <p>Output: <code>[0,0,1,1,2,2]</code></p>
 *
 * <h3>Example 2:</h3>
 *
 * <p>Input: <code>nums = [2,0,1]</code></p>
 * <p>Output: <code>[0,1,2]</code></p>
 *
 *
 *
 * <h3>Constraints:</h3>
 *
 *     <p><code>n == nums.length<code></p>
 *     <p><code>1 <= n <= 300<code></p>
 *     <p><code>nums[i] is either 0, 1, or 2.<code></p>
 *
 *
 *
 * <p>Follow up: Could you come up with a one-pass algorithm using only constant extra space?</p>
 */
public class SortColors {
    public static void sortColors(int[] nums) {
        if (nums.length == 1) return;

        int lo = 0, mid = 0, hi = nums.length - 1;
        while (mid <= hi) {
            switch (nums[mid]) {
                case 0:
                    int temp = nums[lo];
                    nums[lo] = nums[mid];
                    nums[mid] = temp;
                    lo++;
                    mid++;
                    break;
                case 1:
                    mid++;
                    break;
                case 2:
                    int temp1 = nums[hi];
                    nums[hi] = nums[mid];
                    nums[mid] = temp1;
                    hi--;
                    break;
            }
        }
    }
}
