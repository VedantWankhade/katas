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

    public static <T extends Comparable<T>> List<T> quickSort(List<T> list) {
        return quickSort(list, 0, list.size() - 1);
    }

    private static <T extends Comparable<T>> List<T> quickSort(List<T> list, int lo, int hi) {
        if (hi - lo + 1 < 2) return list;
        int pivotIndex = partition(list, lo, hi);
        quickSort(list, lo, pivotIndex - 1);
        quickSort(list, pivotIndex + 1, hi);
        return list;
    }

    private static <T extends Comparable<T>> int partition(List<T> list, int lo, int hi) {
        int pivotIndex = hi;
        T pivot = list.get(hi);
        int i = 0, j = 0;
        while (j < hi) {
            if (list.get(j).compareTo(pivot) <= 0) {
                T temp = list.get(i);
                list.set(i, list.get(j));
                list.set(j, temp);
                i++;
            }
            j++;
        }
        list.set(pivotIndex, list.get(i));
        list.set(i, pivot);
        return i;
    }

    public static <T extends Comparable<T>> List<T> mergeSort(List<T> list) {
        if (list.size() < 2) return list;

        List<T> left = new ArrayList<>();
        List<T> right = new ArrayList<>();

        for (int i = 0; i < list.size() / 2; i++) {
            left.add(list.get(i));
        }
        for (int i = list.size() / 2; i < list.size(); i++) {
            right.add(list.get(i));
        }
        left = mergeSort(left);
        right = mergeSort(right);
        return merge(left, right);
    }

    private static <T extends Comparable<T>> List<T> merge(List<T> left, List<T> right) {
        int m = 0, l = 0, r = 0;
        List<T> main = new ArrayList<>();
        while (l < left.size() && r < right.size()) {
            main.add(left.get(l).compareTo(right.get(r)) <= 0 ? left.get(l++) : right.get(r++));
        }

        if (l >= left.size()) {
            for (int i = r; i < right.size(); i++) {
                main.add(right.get(i));
            }
        } else {
            for (int i = l; i < left.size(); i++) {
                main.add(left.get(i));
            }
        }
        return main;
    }
}
