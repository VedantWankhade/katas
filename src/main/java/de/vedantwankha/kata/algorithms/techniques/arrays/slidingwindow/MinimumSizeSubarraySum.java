package de.vedantwankha.kata.algorithms.techniques.arrays.slidingwindow;

/**
 * <p>https://leetcode.com/problems/minimum-size-subarray-sum/description/</p>
 * <p>Given an array of positive integers nums and a positive integer target, return the minimal length of a</p>
 *
 * <p>whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.</p>
 *
 *
 *
 * <p>Example 1:</p>
 *
 * <p>Input: target = 7, nums = [2,3,1,2,4,3]</p>
 * <p>Output: 2</p>
 * <p>Explanation: The subarray [4,3] has the minimal length under the problem constraint.</p>
 *
 * <p>Example 2:</p>
 *
 * <p>Input: target = 4, nums = [1,4,4]</p>
 * <p>Output: 1</p>
 *
 * <p>Example 3:</p>
 *
 * <p>Input: target = 11, nums = [1,1,1,1,1,1,1,1]</p>
 * <p>Output: 0</p>
 *
 *
 *
 * <p>Constraints:</p>
 *
 *     <p>1 <= target <= 10<sup>9</sup></p>
 *     <p>1 <= nums.length <= 10<sup>5</sup></p>
 *     <p>1 <= nums[i] <= 10<sup>4</sup></p>
 *
 *
 * <p>Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).</p>
 */
public class MinimumSizeSubarraySum {
    public static int minSubArrayLen(int target, int[] nums) {
        if (target == 1 && nums.length == 1) {
            return nums[0] == target ? 1 : 0;
        }

        int i = 0, j = 0;
        int sum = nums[i], len = 1;
        int minLen = Integer.MAX_VALUE;
        if (sum == target) return 1;

        while (j <= i) {
//            System.out.println("i: " + i + "; j: " + j);
            len = i - j + 1;
//            System.out.println("len: " + len);
//            System.out.println("sum: " + sum);
//            System.out.println("minLen: " + minLen);
            if (sum >= target) {
                minLen = Math.min(len, minLen);
                sum = sum - nums[j++];
            } else {
                i++;
                if (i >= nums.length) break;
                sum = sum + nums[i];
            }
        }
        if (minLen >= Integer.MAX_VALUE) return 0;
        return minLen;
    }
}