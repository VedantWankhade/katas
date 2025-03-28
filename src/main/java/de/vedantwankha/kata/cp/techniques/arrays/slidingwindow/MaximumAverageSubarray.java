package de.vedantwankha.kata.cp.techniques.arrays.slidingwindow;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

/**
 * <p>https://leetcode.com/problems/maximum-average-subarray-i/description/?envType=problem-list-v2&envId=sliding-window</p>
 *
 * <p>You are given an integer array nums consisting of n elements, and an integer k.</p>
 *
 * <p>Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.</p>
 *
 *
 *
 * <p>Example 1:</p>
 *
 * <p>Input: nums = [1,12,-5,-6,50,3], k = 4</p>
 * <p>Output: 12.75000</p>
 * E<p>xplanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75</p>
 *
 * <p>Example 2:</p>
 *
 * <p>Input: nums = [5], k = 1</p>
 * <p>Output: 5.00000</p>
 *
 *
 *
 * <p>Constraints:</p>
 *
 *     <p>n == nums.length</p>
 *     <p>1 <= k <= n <= 10<sup>5</sup></p>
 *     <p>-10<sup>4</sup> <= nums[i] <= 10<sup>4</sup></p>
 */
public class MaximumAverageSubarray {
    public static double findMaxAverage(int[] nums, int k) {
        if (nums.length == 1) {
            return nums[0];
        }
        if (k == 1)
            return Arrays.stream(nums).max().getAsInt();
        int j = 0, i = k - 1;
        int sum = 0;
        for (int x = j; x <= i; x++) {
            sum += nums[x];
        }
        i++;
        int maxSum = sum;
//        System.out.println("maxAvg: " + (double)maxSum / k);
        while (i < nums.length) {
//            System.out.println("i: " + i + " j: " + j);
            sum = sum - nums[j] + nums[i];
            maxSum = Math.max(sum, maxSum);
            j++;
            i++;
        }
        return (double)maxSum / k;
    }
}
