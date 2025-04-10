package de.vedantwankha.kata.linkedin.algorithmicthinking.warmup;

/**
 * <h1>100 Doors</h1>
 * <p>There are 100 doors in a row that are initially closed</p>
 * <p>You make 100 passes by the doors</p>
 * <p>On the first pass, you visit every door in sequence and toggle its state</p>
 * <p>On second pass, you visit every second door and toggle its state</p>
 * <p>And so on... unitil you visit only 100th door</p>
 * <p>Find which doors are open at the end</p>
 */
public class HundredDoors {
    public static void solve() {
        var doors = new boolean[101];
        for (int i = 1; i < 101; i++) {
            for (int j = 1; j < 101; j++) {
                if (j % i == 0) {
                    doors[j] = !doors[j];
                }
            }
        }
        for (int i = 1; i < 101; i++) {
            if (doors[i])
                System.out.print(i + " ");
        }
    }
}
