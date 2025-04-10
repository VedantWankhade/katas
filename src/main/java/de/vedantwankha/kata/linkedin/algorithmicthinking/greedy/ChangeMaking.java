package de.vedantwankha.kata.linkedin.algorithmicthinking.greedy;

import java.util.ArrayList;
import java.util.List;

/**
* Find minimum number of coins from a set of denominations that add up to a given amount of money.
 * Assume following denominations: 200, 100, 50, 20, 5, 2, 1
 */
public class ChangeMaking {
    public static List<Integer> solve(int targetAmt) {
        var denoms = new int[] {200, 100, 50, 20, 5, 2, 1};
        var coins = new ArrayList<Integer>();

        for (int d: denoms) {
            while (d <= targetAmt) {
                coins.add(d);
                targetAmt -= d;
            }
        }
        return coins;
    }
}
