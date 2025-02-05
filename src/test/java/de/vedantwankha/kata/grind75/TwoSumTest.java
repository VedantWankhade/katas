package de.vedantwankha.kata.grind75;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class TwoSumTest {
    @ParameterizedTest
    @MethodSource("testProvider")
    void solve(int[] nums, int target, int[] expected) {
        assertEquals(toSet(expected), toSet(TwoSum.solve(nums, target)));
    }

    static Stream<Arguments> testProvider() {
        return Stream.of(
                Arguments.arguments(new int[] {2,7,11,15}, 9, new int[]{0, 1}),
                Arguments.arguments(new int[]{3,2,4}, 6, new int[]{1, 2}),
                Arguments.arguments(new int[]{3, 3}, 6, new int[]{0, 1})
        );
    }

    static private Set<Integer> toSet(int[] nums) {
        var set = new HashSet<Integer>();
        for (int i: nums) {
            set.add(i);
        }
        return set;
    }
}