package de.vedantwankha.kata.algorithms.techniques.recursion;

/**
 * Reverse an array using recursion
 */
public class ReverseArray {
    public static void reverse(int[] arr) {
        if (arr.length == 0 || arr.length == 1) return;
        reverse(arr, 0, arr.length - 1);
    }

    private static void reverse(int[] arr, int lo, int hi) {
        if (lo >= hi) return;
        int temp = arr[lo];
        arr[lo] = arr[hi];
        arr[hi] = temp;
        reverse(arr, lo + 1, hi - 1);
    }
}
