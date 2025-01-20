package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

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
}

@DisplayName("Test empty method")
class TestEmpty {
    @Test
    @DisplayName("Test empty method after creation")
    void testEmpty() {
        ArrayList<Number> arr = new ArrayList<>();
        assertTrue(arr.isEmpty());
    }
}