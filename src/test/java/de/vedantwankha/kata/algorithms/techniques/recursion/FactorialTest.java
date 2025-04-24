package de.vedantwankha.kata.algorithms.techniques.recursion;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class FactorialTest {

    @Test
    void factorial() {
        assertEquals(Factorial.factorial(5), 120);
        assertEquals(Factorial.factorial(5), 120);
        assertEquals(Factorial.factorial(5), 120);
        assertThrows(IllegalArgumentException.class, () -> Factorial.factorial(-1));
    }
}