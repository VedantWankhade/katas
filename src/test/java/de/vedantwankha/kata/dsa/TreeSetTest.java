package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

@DisplayName("Test TreeSet")
public class TreeSetTest {
    @Test
    @DisplayName("Test add")
    void testSearch() {
        Set<Integer> t = new TreeSet<>();
        t.add(5);
        t.add(2);
        t.add(1);
        t.add(6);
        t.add(7);
        t.add(3);
        assertEquals(6, t.size());
        System.out.println(t);
        t.add(3);
        assertEquals(6, t.size());
        System.out.println(t);
        t.add(3);
    }
}
