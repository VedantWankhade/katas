package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

@DisplayName("Test hash map")
public class HashMapTest {
    @Test
    @DisplayName("Test put and get methods")
    void testPutGet() {
        HashMap<String, String> m = new HashMap<>(4);
        m.put("1", "One");
        m.put("2", "Two");
        m.put("3", "Three");
        System.out.println(m.get("1"));
        System.out.println(m.get("2"));
        System.out.println(m.get("3"));

        System.out.println(m.toStringChain());

        m.put("4", "Hello");
        System.out.println(m.toStringChain());

    }

    @Test
    @DisplayName("Test entrySet")
    void testEntrySet() {
        HashMap<String, String> m = new HashMap<>(4);
        m.put("1", "One");
        m.put("2", "Two");
        m.put("3", "Three");
        System.out.println(m.entrySet());
    }

    @Test
    @DisplayName("Test iterator")
    void testIterator() {
        HashMap<String, String> m = new HashMap<>(4);
        m.put("1", "One");
        m.put("2", "Two");
        m.put("3", "Three");
        m.put("6", "Six");
        m.put("87", "Eight Seven");
        m.put("11", "Eleven");
        System.out.println(m.toStringChain());
        for (var e: m) {
            System.out.println(e);
        }
    }
}
