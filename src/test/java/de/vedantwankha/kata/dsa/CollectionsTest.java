package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class CollectionsTest {
    @Test
    @DisplayName("Test binary search")
    void testBinarySearch() {
        List<Integer> l = new ArrayList<>();
        l.add(1);
        l.add(2);
        l.add(3);
        l.add(4);
        l.add(5);
        l.add(6);
        l.add(7);
        assertEquals(0, Collections.binarySearch(l, 1));
        assertEquals(1, Collections.binarySearch(l, 2));
        assertEquals(4, Collections.binarySearch(l, 5));
        assertEquals(5, Collections.binarySearch(l, 6));
        assertEquals(6, Collections.binarySearch(l, 7));
        assertEquals(-1, Collections.binarySearch(l, 99));

        List<Integer> ll = new LinkedList<>();
        ll.add(1);
        ll.add(2);
        assertEquals(1, Collections.binarySearch(ll, 2));
    }

    @Test
    @DisplayName("Test selection sort")
    void testSelectionSort() {
        List<Integer> l = new ArrayList<>();
        l.add(5);
        l.add(2);
        l.add(1);
        l.add(7);
        l.add(5);
        l.add(8);
        l.add(3);

        System.out.println(Collections.selectionSort(l));
    }
}
