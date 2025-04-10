package de.vedantwankha.kata.linkedin.algorithmicthinking.bruteforce;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LinearSearchTest {

    @Test
    void search() {
        assertEquals(LinearSearch.search(new Integer[]{2, 4, 6, 1, 0}, 22), -1);
        assertEquals(LinearSearch.search(new Integer[]{2, 4, 6, 1, 0}, 2), 0);
        assertEquals(LinearSearch.search(new Integer[]{2, 4, 6, 1, 0}, 0), 4);
        assertEquals(LinearSearch.search(new Integer[]{2, 4, 6, 1, 0}, 1), 3);
    }
}