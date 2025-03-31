package de.vedantwankha.kata.algorithms.techniques.arrays.twopointers;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SortColorsTest {

    @Test
    void sortColors() {
        var nums = new int[]{2,0,2,1,1,0};
        System.out.println(Arrays.toString(nums));
        SortColors.sortColors(nums);
        System.out.println(Arrays.toString(nums));

        var nums1 = new int[]{2,0,1};
        System.out.println(Arrays.toString(nums1));
        SortColors.sortColors(nums1);
        System.out.println(Arrays.toString(nums1));
    }
}