package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

@DisplayName("Test bst general methods")
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
        t.add(4);
        assertEquals(7, t.size());
        int j = 1;
        for (int i: t) {
            assertEquals(j, i);
            j++;
        }
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
        assertEquals(1, n.get(0).getData());
        assertEquals(2, n.get(1).getData());

        ImmutableList<BST.Node<Integer>> n1 = t.get(99);
        assertNull(n1.get(0));
        assertNull(n1.get(1));

        ImmutableList<BST.Node<Integer>> n2 = t.get(5);
        assertEquals(5, n2.get(0).getData());
        assertNull(n2.get(1));
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
        assertEquals(7, t.max().get(0).getData());
        assertEquals(1, t.min().get(0).getData());

        BST<Integer> t2 = new BST<>();
        t2.add(5);
        assertEquals(5, t2.max().get(0).getData());
        assertEquals(5, t2.min().get(0).getData());
    }
}

@DisplayName("Test remove method")
class TestRemove {
    BST<Integer> t;

    @BeforeEach()
    void init() {
         t = new BST<>();
         t.add(5);
         t.add(2);
         t.add(1);
         t.add(6);
         t.add(7);
         t.add(3);
    }

    @Test
    @DisplayName("Test remove leaf")
    void testRemoveLeaf() {
        assertTrue(t.contains(1));
        System.out.println(t);
        t.remove(1);
        System.out.println(t);
        assertFalse(t.contains(1));

        assertTrue(t.contains(3));
        System.out.println(t);
        t.remove(3);
        System.out.println(t);
        assertFalse(t.contains(3));

        assertTrue(t.contains(7));
        System.out.println(t);
        t.remove(7);
        System.out.println(t);
        assertFalse(t.contains(7));
    }

    @Test
    @DisplayName("Test remove leaf if it is root")
    void testRemoveLeafRoot() {
        var t1 = new BST<Integer>();
        t1.add(5);
        System.out.println(t1);
        assertTrue(t1.contains(5));
        t1.remove(5);
        System.out.println(t1);
        assertFalse(t1.contains(5));
    }

    @Test
    @DisplayName("Test remove node with one child subtree")
    void testRemoveOneSubtree() {
        assertTrue(t.contains(6));
        System.out.println(t);
        t.remove(6);
        assertFalse(t.contains(6));
        System.out.println(t);
    }

    @Test
    @DisplayName("Test remove node with one child subtree if its root")
    void testRemoveOneSubtreeRoot() {
        var t1 = new BST<Integer>();
        t1.add(5);
        t1.add(6);
        assertTrue(t1.contains(5));
        System.out.println(t1);
        t1.remove(5);
        assertFalse(t1.contains(5));
        System.out.println(t1);
    }

    @Test
    @DisplayName("Test remove node with two subtrees")
    void testRemoveTwoSubtree() {
        assertTrue(t.contains(2));
        System.out.println(t);
        t.remove(2);
        assertFalse(t.contains(2));
        System.out.println(t);
    }

    @Test
    @DisplayName("Test remove node with two subtrees if its root")
    void testRemoveTwoSubtreeRoot() {
        var t1 = new BST<Integer>();
        t1.add(5);
        t1.add(6);
        t1.add(2);
        assertTrue(t1.contains(5));
        System.out.println(t1);
        t1.remove(5);
        assertFalse(t1.contains(5));
        System.out.println(t1);
    }
}
