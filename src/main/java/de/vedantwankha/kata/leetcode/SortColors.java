package de.vedantwankha.kata.leetcode;

/**
 * <h1>Sort Colors</h1>
 * <p><a href="https://leetcode.com/problems/sort-colors/description/">https://leetcode.com/problems/sort-colors/description/</a></p>
 * <p>Given an array nums with <code>n</code> objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.</p>
 * <p>We will use the integers <code>0</code>, <code>1</code>, and <code>2</code> to represent the color red, white, and blue, respectively.</p>
 * <p>You must solve this problem without using the library's sort function.</p>
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
 *     <p><code>n == nums.length</code></p>
 *     <p><code>1 <= n <= 30<sup>0</sup></code></p>
 *     <p><code>nums[i] is either 0, 1, or 2.</code></p>
 *  <hr/>
 *
 * <p>Follow up: Could you come up with a one-pass algorithm using only constant extra space?</p>
 *  <hr/>
 * <small>
 *     <p><b>Hint1</b>: A rather straight forward solution is a two-pass algorithm using counting sort.</p>
 *     <p><b>Hint2</b>: Iterate the array counting number of 0's, 1's, and 2's.</p>
 *     <p><b>Hint3</b>: Overwrite array with the total number of 0's, then 1's and followed by 2's.</p>
 * </small>
 */
public class SortColors {
    public static void sortColors(int[] nums) {
        if (nums.length == 0) return;


    }
}
