package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.*;

import java.util.EmptyStackException;

import static org.junit.jupiter.api.Assertions.*;

class TestArrayList {
    @Nested
    @DisplayName("Test as a Stack")
    class TestAsStack {
        @Test
        @DisplayName("Test stack abstraction")
        void testStack() {
            Stack<Integer> stack = new ArrayList<>();
            stack.push(1);
            stack.push(2);
            stack.push(3);
            stack.push(4);
            System.out.println(stack);
            assertEquals(4, stack.size());
            assertEquals(4, stack.pop());
            System.out.println(stack);
            assertEquals(3, stack.size());
            assertEquals(3, stack.pop());
            System.out.println(stack);
            assertEquals(2, stack.size());
            assertEquals(2, stack.pop());
            System.out.println(stack);
            assertEquals(1, stack.size());
            assertEquals(1, stack.pop());
            System.out.println(stack);
            assertThrows(EmptyStackException.class, stack::pop);

        }
    }

    @Nested
    @DisplayName("Test contains method")
    class TestContains {
        @Test
        @DisplayName("Test contains")
        void testContains() {
            ArrayList<Number> arr1 = new ArrayList<>();
            arr1.add(1);
            arr1.add(2);
            arr1.add(3);
            arr1.add(4);
            assertTrue(arr1.contains(2));
            assertFalse(arr1.contains(99));
        }
    }

    @Nested
    @DisplayName("Test remove methods")
    class TestRemove {
        List<Integer> list;

        @BeforeEach
        void setup() {
            list = new ArrayList<>();
            list.add(1);
            list.add(2);
            list.add(3);
            list.add(4);
            list.add(5);
        }

        @Test
        @DisplayName("Test remove e method")
        void TestRemoveE() {
            System.out.println(list);
            list.remove(Integer.valueOf(3));
            System.out.println(list);
            assertNotEquals(3, list.get(2));
        }

        @Test
        @DisplayName("Test remove at index method")
        void TestRemoveAtIndex() {
            System.out.println(list);
            list.remove(2);
            System.out.println(list);
            assertNotEquals(3, list.get(2));
        }
    }

    @Nested
    @DisplayName("Test size method")
    class TestSize {
        @Test
        @DisplayName("Test size method for empty ArrayList")
        void testEmpty() {
            ArrayList<Number> arr = new ArrayList<>();
            assertEquals(0, arr.size());

            arr.add(2);
            arr.remove(2);
            assertEquals(0, arr.size());
        }

        @Test
        @DisplayName("Test size method for after items inserted and deleted")
        void testDefaultConstructorArrayList() {
            ArrayList<Number> arr = new ArrayList<>();
            arr.add(2);
            assertEquals(1, arr.size());
            arr.add(4);
            assertEquals(2, arr.size());
            arr.set(1, 99);
            assertEquals(2, arr.size());
            arr.remove(99);
            assertEquals(1, arr.size());
            arr.remove(2);
            assertEquals(0, arr.size());
        }
    }

    @Nested
    @DisplayName("Test add method")
    class TestAdd {
        @Test
        @DisplayName("Test add method")
        void testAdd() {
            ArrayList<Number> arr = new ArrayList<>();
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);
            arr.add(1);
            arr.add(1);
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);
            arr.add(1);
            arr.add(1);
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);
            arr.add(1);
            arr.add(1);
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);
            arr.add(1);
            arr.add(1);
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);
            arr.add(1);
            arr.add(1);
            System.out.println(arr.size() + " " + arr.getCapacity() + " " + arr);

            for (Number number : arr) System.out.println(number);

        }

        @Test
        @DisplayName("Test add all method")
        void testAddAll() {
            ArrayList<Number> arr = new ArrayList<>();
            ArrayList<Number> arr1 = new ArrayList<>();
            arr1.add(1);
            arr1.add(2);
            arr1.add(3);
            arr1.add(4);
            arr.addAll(arr1);
            assertEquals(4, arr.size());
            System.out.println(arr);
        }

        @Test
        @DisplayName("Test add at index method")
        void testAddAtIndex() {
            ArrayList<Number> arr = new ArrayList<>();
            arr.add(1);
            arr.add(2);
            arr.add(3);
            arr.add(4);
            System.out.println(arr);
            arr.add(99, 2);
            System.out.println(arr);
        }

        @Test
        @DisplayName("Test add all method with already existing elements")
        void testAddAllExisting() {
            ArrayList<Number> arr = new ArrayList<>();
            ArrayList<Number> arr1 = new ArrayList<>();
            arr.add(22);
            arr.add(99);
            arr1.add(1);
            arr1.add(2);
            arr1.add(3);
            arr1.add(4);
            arr.addAll(arr1);
            assertEquals(6, arr.size());
            System.out.println(arr);
        }
    }

    @Nested
    @DisplayName("Test empty method")
    class TestEmpty {
        @Test
        @DisplayName("Test empty method after creation")
        void testEmpty() {
            ArrayList<Number> arr = new ArrayList<>();
            assertTrue(arr.isEmpty());
        }
    }

    @Nested
    @DisplayName("Test get method")
    class TestGet {
        @Test
        @DisplayName("Test get method")
        void testGet() {
            List<Number> arr1 = new ArrayList<>();
            arr1.add(1);
            arr1.add(2);
            arr1.add(3);
            arr1.add(4);
            for (int i = 0; i < arr1.size(); i++) {
                assertEquals(i + 1, arr1.get(i));
                System.out.println(arr1.get(i));
            }
        }
    }

    @Nested
    @DisplayName("Test toArray")
    class TestToArray {
        @Test
        @DisplayName("Test toArray")
        void testToArray() {
            List<Integer> l = new ArrayList<Integer>();
            l.add(1);
            l.add(2);
            l.add(3);
            var expt = new Integer[3];
            expt[0] = 1;
            expt[1] = 2;
            expt[2] = 3;
            Object[] act = l.toArray();
            for (int i = 0; i < expt.length; i++) {
                assertEquals(expt[i], act[i]);
                System.out.println(act[i]);
            }
        }
    }

    @Nested
    @DisplayName("Test abstraction")
    class TestAbstraction {
        @Test
        @DisplayName("Test List abstraction")
        void testListAbstraction() {
            var l = new ArrayList<Number>();
            l.add(1);
            l.add(2);
            l.add(3);
            List<Number> arr1 = new ArrayList<>(l);
            for (int i = 0; i < arr1.size(); i++) {
                assertEquals(i + 1, arr1.get(i));
                System.out.println(arr1.get(i));
            }
            arr1.add(22);
        }

        @Test
        @DisplayName("Test ImmutableList abstraction")
        void testImmutableListAbstraction() {
            var l = new ArrayList<Number>();
            l.add(1);
            l.add(2);
            l.add(3);
            ImmutableList<Number> arr1 = new ArrayList<>(l);
            for (int i = 0; i < arr1.size(); i++) {
                assertEquals(i + 1, arr1.get(i));
                System.out.println(arr1.get(i));
            }
            // the following should throw compile error
            // arr1.add(22);
        }
    }
}