package de.vedantwankha.kata.cp.techniques.arrays.twopointers;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class TwoSumSortedTest {

    @Test
    void twoSum() {
        System.out.println(Arrays.toString(TwoSumSorted.twoSum(new int[]{2,7,11,15}, 9)));
    }
}