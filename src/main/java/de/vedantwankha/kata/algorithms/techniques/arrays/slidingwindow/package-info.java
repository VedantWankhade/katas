/**
 * <h1>Sliding Window Technique</h1>
 * <ul>
 * <li>Specific version of two pointers technique</li>
 * <li>Two pointers moving in the same direction and never overtake each other</li>
 * <li>Maintain a window of elements (between two pointers) that always satisfy the given condition</li>
 * <li>Ensured that each value is visited at most twice ensuring O(n)</li>
 * <li>Only works for questions which ask for subarray / substring (consecutive elements). And when you dont need to swap or change positions of elements</li>
 * <li>Works even when you need to swap or change the positions of elements (as opposed to sliding window)</li>
 * <li>Two types</li>
 * <ul>
 * 	<li>Fixed size: window stays constant size and moves one step at a time</li>
 * 	<li>Variable size: window expands or shrinks based on conditions</li>
 * 	</ul>
 * </ul>
 */
package de.vedantwankha.kata.algorithms.techniques.arrays.slidingwindow;