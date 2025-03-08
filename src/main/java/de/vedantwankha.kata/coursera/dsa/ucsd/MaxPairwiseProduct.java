package de.vedantwankha.kata.coursera.dsa.ucsd;

import java.util.*;
import java.io.*;

public class MaxPairwiseProduct {
    static int getMaxPairwiseProduct(int[] numbers) {
        int max_product = 0;
        int n = numbers.length;

        for (int first = 0; first < n; ++first) {
            for (int second = first + 1; second < n; ++second) {
                max_product = Math.max(max_product,
                        numbers[first] * numbers[second]);
            }
        }

        return max_product;
    }

    static long getMaxPairwiseProductFast(int[] numbers) {
        int maxIndex1 = 0;
        int maxIndex2 = 0;
        for (int i = 1; i < numbers.length; i++) {
            if (numbers[maxIndex1] < numbers[i]) {
                maxIndex1 = i;
            }
        }

        int temp = numbers[numbers.length - 1];
        numbers[numbers.length - 1] = numbers[maxIndex1];
        numbers[maxIndex1] = temp;

        for (int i = 1; i < numbers.length - 1; i++) {
            if (numbers[maxIndex2] < numbers[i]) {
                maxIndex2 = i;
            }
        }
        return (long) (numbers[numbers.length - 1]) * (long) (numbers[maxIndex2]);
    }

    public static void main(String[] args) {
        FastScanner scanner = new FastScanner(System.in);
        int n = scanner.nextInt();
        int[] numbers = new int[n];
        for (int i = 0; i < n; i++) {
            numbers[i] = scanner.nextInt();
        }
        System.out.println(getMaxPairwiseProductFast(numbers));
    }

    static class FastScanner {
        BufferedReader br;
        StringTokenizer st;

        FastScanner(InputStream stream) {
            try {
                br = new BufferedReader(new
                        InputStreamReader(stream));
            } catch (Exception e) {
                e.printStackTrace();
            }
        }

        String next() {
            while (st == null || !st.hasMoreTokens()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }
    }

}

