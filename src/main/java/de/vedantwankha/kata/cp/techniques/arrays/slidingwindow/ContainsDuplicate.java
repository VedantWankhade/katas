package de.vedantwankha.kata.cp.techniques.arrays.slidingwindow;

import java.util.HashSet;
import java.util.Set;

/**
 * <p>https://leetcode.com/problems/contains-duplicate-ii/description/?envType=problem-list-v2&envId=sliding-window</p>
 * <p>Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.</p>
 * <p>Example 1:</p>
 *
 * <p>Input: nums = [1,2,3,1], k = 3</p>
 * <p>Output: true</p>
 *
 * <p>Example 2:</p>
 *
 * <p>Input: nums = [1,0,1,1], k = 1</p>
 * <p>Output: true</p>
 *
 * <p>Example 3:</p>
 *
 * <p>Input: nums = [1,2,3,1,2,3], k = 2</p>
 * <p>Output: false</p>
 *
 *
 *
 * <p>Constraints:</p>
 *
 *     <p>1 <= nums.length <= 10<sup>5</sup></p>
 *     <p>-10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></p>
 *     <p>0 <= k <= 10<sup>5</sup></p>
 */
public class ContainsDuplicate {
    public static boolean containsNearbyDuplicate(int[] nums, int k) {
        if (nums.length < 2 || (nums.length == 2 && k < 1) || k < 1) {
            return false;
        }

        int i = 0, j = 0;
        Set<Integer> set = new HashSet<>();

        while (i < nums.length) {
//            System.out.println("i: " + i);
//            System.out.println("j: " + j);
//            System.out.println("i - j: " + (i - j));
//            System.out.print("set: " + set + " -> ");
            if (i - j > k) {
                set.remove(nums[j]);
                j++;
            }
            else {
                if (set.contains(nums[i])) return true;
                set.add(nums[i]);
                i++;
            }
//            System.out.println(set);
        }

        return false;
    }
}
