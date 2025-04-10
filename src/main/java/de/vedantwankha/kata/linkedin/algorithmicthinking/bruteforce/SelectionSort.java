package de.vedantwankha.kata.linkedin.algorithmicthinking.bruteforce;

import java.util.Comparator;

public class SelectionSort {
    public static <T extends Comparable<T>> void sort(T[] arr) {
        sort(arr, (t1, t2) -> t1.compareTo(t2));
    }

    public static <T> void sort(T[] arr, Comparator<T> c) {
        for (int i = 0; i < arr.length - 1; i++) {
            int smallestIndex = i;
            int j = i;
            while (j < arr.length) {
                if (c.compare(arr[j], arr[smallestIndex]) < 0) {
                    smallestIndex = j;
                }
                j++;
            }
            T temp = arr[i];
            arr[i] = arr[smallestIndex];
            arr[smallestIndex] = temp;
        }
    }
}
