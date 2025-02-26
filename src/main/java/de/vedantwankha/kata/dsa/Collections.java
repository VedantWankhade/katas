package de.vedantwankha.kata.dsa;

public class Collections {
    public static <T extends Comparable<T>> int binarySearch(List<? extends T> list, T key) {
        // somehow assert sorted order

        if (list instanceof LinkedList<? extends T>)
            System.getLogger("log").log(System.Logger.Level.WARNING, "Binary Search on Linked List does not make any sense.");

        int lo = 0, hi = list.size() - 1;
        int mid;

        while (lo <= hi) {
            mid = (lo + hi) / 2;
            if (key.compareTo(list.get(mid)) == 0) {
                return mid;
            } else if (key.compareTo(list.get(mid)) < 0) {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return -1;
    }

    public static <T extends Comparable<T>> List<T> selectionSort(List<T> list) {
        for (int i = 0; i < list.size() - 1; i++) {
            int maxIdx = i;
            for (int j = i + 1; j < list.size(); j++) {
                if (list.get(maxIdx).compareTo(list.get(j)) < 0) {
                    maxIdx = j;
                }
            }
            T temp = list.get(maxIdx);
            list.set(maxIdx, list.get(i));
            list.set(i, temp);
        }
        return list;
    }
}
