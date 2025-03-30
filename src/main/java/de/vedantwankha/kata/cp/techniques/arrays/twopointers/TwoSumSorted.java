package de.vedantwankha.kata.cp.techniques.arrays.twopointers;

/**
 * <p>https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/?envType=problem-list-v2&envId=two-pointers</p>
 * <p>Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.</p>
 *
 * <p>Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.</p>
 *
 * <p>The tests are generated such that there is exactly one solution. You may not use the same element twice.</p>
 *
 * <p>Your solution must use only constant extra space.</p>
 *
 * <p>Example 1:</p>
 *
 * <p>Input: numbers = [2,7,11,15], target = 9</p>
 * <p>Output: [1,2]</p>
 * <p>Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].</p>
 *
 * <p>Example 2:</p>
 *
 * <p>Input: numbers = [2,3,4], target = 6</p>
 * <p>Output: [1,3]</p>
 * <p>Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].</p>
 *
 * <p>Example 3:</p>
 *
 * <p>Input: numbers = [-1,0], target = -1</p>
 * <p>Output: [1,2]</p>
 * <p>Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].</p>
 *
 *
 *
 * <p>Constraints:</p>
 *
 *     <p>2 <= numbers.length <= 3 * 104</p>
 *     <p>-1000 <= numbers[i] <= 1000</p>
 *     <p>numbers is sorted in non-decreasing order.</p>
 *     <p>-1000 <= target <= 1000</p>
 *     <p>The tests are generated such that there is exactly one solution.</p>
 */
public class TwoSumSorted {
    public static int[] twoSum(int[] numbers, int target) {
        if (numbers.length == 2) return new int[] {1, 2};

        int lo = 0, hi = numbers.length - 1;
        while (lo < hi) {
            if (target == numbers[lo] + numbers[hi]) {
                return new int[]{lo + 1, hi + 1};
            } else if (target < numbers[lo] + numbers[hi]) {
                hi--;
            } else {
                lo++;
            }
        }

        return new int[] {0, 0};
    }
}
