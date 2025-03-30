package de.vedantwankha.kata.algorithms.techniques.arrays.slidingwindow;

import org.junit.jupiter.api.Test;

class MinimumSizeSubarraySumTest {

    @Test
    void minSubArrayLen() {
        System.out.println(MinimumSizeSubarraySum.minSubArrayLen(7, new int[]{2,3,1,2,4,3}));
        System.out.println(MinimumSizeSubarraySum.minSubArrayLen(11, new int[]{1,2,3,4,5}));
        System.out.println(MinimumSizeSubarraySum.minSubArrayLen(0, new int[]{1,1,1,1,1,1,1,1}));
    }
}