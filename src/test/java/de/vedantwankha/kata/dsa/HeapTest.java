package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

@DisplayName("Test heap")
public class HeapTest {
    @Test
    @DisplayName("Test heap")
    void testHeap() {
        Heap<Integer> h = new Heap<>();
        h.add(5);
        System.out.println(h);
        h.add(3);
        System.out.println(h);
        h.add(9);
        System.out.println(h);
        h.add(2);
        System.out.println(h);
        h.add(1);
        while(!h.isEmpty()) {
            System.out.println(h.extract());
            System.out.println(h);
        }
    }
}
