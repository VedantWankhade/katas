package de.vedantwankha.kata.algorithms.techniques.recursion;

/**
 * Check if the string is palindrome
 */
public class Palindrome {
    public static boolean palindrome(String s) {
        if (s.length() < 2) return true;
        return palindrome(s, 0, s.length() - 1);
    }

    private static boolean palindrome(String s, int lo, int hi) {
        if (lo > hi) {
            return true;
        }
        if (lo == hi) return true;
        if (s.charAt(lo) == s.charAt(hi)) {
            return palindrome(s, lo + 1, hi - 1);
        }
        return false;
    }
}
