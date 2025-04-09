package de.vedantwankha.kata.linkedin.algorithmicthinking;

/**
 * <h1>Game of two players</h1>
 * <p>Take in turns to count from 1 to 100, but each time you are encounter a multiple of 3 say "fizz" instead of 3 and when you encounter a multiple of 5 say "buzz". When you encounter a number that is a multiple of both 3 and 5 then say "fizz, buzz"</p>
 */
public class FizzBuzz {
    public static void solve() {
        for (int i = 1; i < 101; i++) {
            System.out.print(i + " ");
            if (i % 3 == 0) {
                System.out.print("fizz");
            }
            if (i % 5 == 0) {
                if (i % 3 == 0) {
                    System.out.print(", ");
                }
                System.out.print("buzz");
            }
            System.out.println();
        }
    }
}
