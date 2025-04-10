package de.vedantwankha.kata.linkedin.algorithmicthinking.bruteforce;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SelectionSortTest {

    @Test
    void sort() {
        var arr = new Integer[]{2, 5, 6, 1, 0};
        System.out.println(Arrays.toString(arr));
        SelectionSort.sort(arr);
        System.out.println(Arrays.toString(arr));
    }

    @Test
    void testSort() {
        var arr = new Integer[]{2, 5, 6, 1, 0};
        System.out.println(Arrays.toString(arr));
        SelectionSort.sort(arr, (t1, t2) -> t2.compareTo(t1));
        System.out.println(Arrays.toString(arr));
    }
}