package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

@DisplayName("Test bst")
class BSTTest {
    @Test
    @DisplayName("Test add method")
    void testAdd() {
        BST<Integer> t = new BST<>();
        t.add(5);
        t.add(2);
        t.add(1);
        t.add(6);
        t.add(7);
        t.add(3);
        System.out.println("Size: " + t.size());
        System.out.println(t);
    }

    @Test
    @DisplayName("Test search")
    void testSearch() {
        BST<Integer> t = new BST<>();
        t.add(5);
        t.add(2);
        t.add(1);
        t.add(6);
        t.add(7);
        t.add(3);
        ImmutableList<BST.Node<Integer>> n = t.get(1);
        System.out.println(n.get(0));
        System.out.println(n.get(1));

        ImmutableList<BST.Node<Integer>> n1 = t.get(99);
        System.out.println(n1.get(0));
        System.out.println(n1.get(1));
    }

    @Test
    @DisplayName("Test min and max")
    void testMinMax() {
        BST<Integer> t = new BST<>();
        t.add(5);
        t.add(2);
        t.add(1);
        t.add(6);
        t.add(7);
        t.add(3);
        System.out.println(t.max());
        System.out.println(t.min());
    }
}
