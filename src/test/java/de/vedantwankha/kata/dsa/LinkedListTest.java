package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class TestLinkedList {
    @Nested
    @DisplayName("Test constructors")
    class TestConstructors {
        @Test
        @DisplayName("Default constructor")
        void TestDefaultConstructor() {
            LinkedList<String> l = new LinkedList<>();
            assertEquals(0, l.size());
        }

        @Test
        @DisplayName("Default with collection")
        void testConstructorWithCollection() {
            List<String> al = new ArrayList<>();
            al.add("1");
            al.add("2");
            al.add("3");
            al.add("4");
            LinkedList<String> ll = new LinkedList<>(al);
            assertEquals(4, ll.size());
        }
    }

    @Nested
    @DisplayName("Test get method")
    class TestGet {
        @Test
        @DisplayName("Test get at index method")
        void testGetAtIndex() {
            List<String> ll = new LinkedList<>();
            ll.add("1");
            ll.add("2");
            ll.add("3");
            ll.add("4");
            ll.add("5");
            assertEquals("1", ll.get(0));
            assertEquals("2", ll.get(1));
            assertEquals("5", ll.get(4));
        }
    }

    @Nested
    @DisplayName("Test add methods")
    class TestAdd {

        @Test
        @DisplayName("Test addFirst method")
        void testAddFirst() {
            List<Integer> ll1 = new LinkedList<>();
            ll1.add(1);
            ll1.add(2);
            ll1.add(3);
            System.out.println(ll1);
            ll1.addFirst(99);
            System.out.println(ll1);

            List<Integer> ll2 = new LinkedList<>();
            System.out.println(ll2);
            ll2.addFirst(99);
            System.out.println(ll2);
            ll2.addFirst(88);
            System.out.println(ll2);
        }

        @Test
        @DisplayName("Default addLast method")
        void TestDefaultConstructor() {
            List<String> al = new ArrayList<>();
            al.add("1");
            al.add("2");
            al.add("3");
            al.add("4");
            LinkedList<String> ll = new LinkedList<>(al);
            System.out.println(ll);
            ll.addLast("99");
            System.out.println(ll);

            List<Integer> ll2 = new LinkedList<>();
            System.out.println(ll2);
            ll2.addLast(99);
            System.out.println(ll2);
            ll2.addLast(88);
            System.out.println(ll2);
        }

        @Test
        @DisplayName("Test add at index method")
        void testAtIndex() {
            List<Integer> ll = new LinkedList<>();
            ll.add(1);
            ll.add(2);
            System.out.println(ll);
            ll.add(99, 1);
            System.out.println(ll);
            assertEquals(99, ll.get(1));
        }
    }

    @Test
    @DisplayName("Test iterator")
    void testIterator() {
        List<Integer> ll = new LinkedList<>();
        ll.add(1);
        ll.add(2);
        ll.add(3);
        ll.add(4);
        ll.add(5);
        System.out.println(ll);
        for (Integer i: ll) {
            System.out.println(i);
        }
    }
}