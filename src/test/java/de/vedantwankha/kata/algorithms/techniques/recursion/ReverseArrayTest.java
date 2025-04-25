package de.vedantwankha.kata.algorithms.techniques.recursion;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ReverseArrayTest {

    @Test
    void reverse() {
        int[] arr1 = {1, 2, 3, 4, 5};
        ReverseArray.reverse(arr1);
        assertArrayEquals(new int[]{5, 4, 3, 2, 1}, arr1);

        int[] arr2 = {1, 2};
        ReverseArray.reverse(arr2);
        assertArrayEquals(new int[]{2, 1}, arr2);

        int[] arr3 = {2};
        ReverseArray.reverse(arr3);
        assertArrayEquals(new int[]{2}, arr3);
    }
}