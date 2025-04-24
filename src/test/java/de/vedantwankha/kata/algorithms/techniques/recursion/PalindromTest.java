package de.vedantwankha.kata.algorithms.techniques.recursion;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class PalindromTest {

    @Test
    void palindrome() {
        assertTrue(Palindrome.palindrome("aba"));
        assertTrue(Palindrome.palindrome("aa"));
        assertTrue(Palindrome.palindrome("a"));
        assertFalse(Palindrome.palindrome("ab"));
        assertFalse(Palindrome.palindrome("abc"));
        assertTrue(Palindrome.palindrome("abba"));
    }
}