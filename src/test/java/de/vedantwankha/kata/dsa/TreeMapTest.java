package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

@DisplayName("Test TreeMap")
public class TreeMapTest {
    @Test
    @DisplayName("Test put method")
    void testPut() {
        Map<Integer, String> m = new TreeMap<>(TreeMap.BSTType.BST);
        assertTrue(m.isEmpty());
        System.out.println(m);
        assertEquals(0, m.size());
        m.put(3, "Three");
        assertFalse(m.isEmpty());
        assertEquals(1, m.size());
        m.put(5, "Five");
        assertEquals(2, m.size());
        m.put(1, "One");
        assertEquals(3, m.size());
        m.put(2, "Two");
        assertEquals(4, m.size());
        m.put(10, "Ten");
        assertEquals(5, m.size());
        System.out.println(m);
    }

    @Test
    @DisplayName("Test contains and remove method")
    void testContainsAndRemove() {
        Map<Integer, String> m = new TreeMap<>(TreeMap.BSTType.BST);
        m.put(3, "Three");
        m.put(5, "Five");
        m.put(1, "One");
        m.put(2, "Two");
        m.put(10, "Ten");

        assertTrue(m.containsKey(2));
        assertTrue(m.containsKey(5));
        assertTrue(m.containsKey(1));
        assertFalse(m.containsKey(0));
        assertFalse(m.containsKey(99));
        System.out.println(m);
        assertTrue(m.containsKey(10));
        m.remove(10);
        System.out.println(m);
        assertFalse(m.containsKey(10));
    }

    @Test
    @DisplayName("Test keySet, getValues and entrySet methods")
    void testKeysValuesEntries() {
        Map<Integer, String> m = new TreeMap<>(TreeMap.BSTType.BST);
        m.put(3, "Three");
        m.put(5, "Five");
        m.put(1, "One");
        m.put(2, "Two");
        m.put(10, "Ten");

        Collection<String> vals = m.values();
        System.out.println(vals);
        Set<Integer> keys = m.keySet();
        System.out.println(keys);
        Set<Map.Entry<Integer, String>> entries = m.entrySet();
        System.out.println(entries);
    }
}
