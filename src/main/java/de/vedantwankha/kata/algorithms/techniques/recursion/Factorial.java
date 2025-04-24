package de.vedantwankha.kata.algorithms.techniques.recursion;

/**
 * Find factorial of given number
 * n! = n * n - 1 * n - 2 * ... * 2 * 1
 */
public class Factorial {
    public static long factorial(int n) {
        if (n < 0) throw new IllegalArgumentException("Factorials are not defined for negative numbers");
        long m = (long) n;
        return factorial(m);
    }

    private static long factorial(long n) {
        if (n == 0 || n == 1) return 1;
        return n * factorial(n - 1);
    }
}
