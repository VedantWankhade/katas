package de.vedantwankha.kata.dsa;

import java.util.Arrays;
import java.util.Iterator;

public class ArrayList<E> implements Collection<E> {
    private final int GROWTH_FACTOR = 2;

    public int getCapacity() {
        return capacity;
    }

    private int capacity, size;
    private E[] items;


    public ArrayList() {
        this(4);
    }

    public ArrayList(int capacity) {
        this.capacity = capacity;
        this.items = (E[]) new Object[this.capacity];
    }

    public ArrayList(Collection<? extends E> col) {
        this.capacity = col.size();
        this.items = (E[]) new Object[this.capacity];
        E[] colArr = col.toArray();
        for (int i = 0; i < this.capacity; i++) {
            this.add(colArr[i]);
        }
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public boolean contains(E e) {
        return false;
    }

    @Override
    public void add(E e) {
        if (this.size == this.capacity) resize();
        this.items[this.size] = e;
        this.size++;
    }

    private void resize(int capacity) {
        E[] newArr = (E[]) new Object[capacity];
        for (int i = 0; i < this.size; i++) {
            newArr[i] = this.get(i);
        }
        this.items = newArr;
        this.capacity = capacity;
    }

    private void resize() {
        resize(GROWTH_FACTOR * this.capacity);
    }

    /**
     * extend here to save space and time
     * @param c
     */
    @Override
    public void addAll(Collection<? extends E> c) {
        if (this.getCapacity() < this.size() + c.size()) {
            resize(this.size + c.size());
        }
        System.out.println(c.size());
        System.out.println(this.capacity);
        for (int i = this.size, j = 0; i < this.capacity; i++, j++) {
            this.set(i, c.get(j));
            size++;
        }
    }

    /**
     * Removes the first occurance of e
     * @param e element
     */
    @Override
    public void remove(E e) {
        for (int i = 0; i < this.size; i++) {
            if (e.equals(this.get(i))) {
                int j = i;
                while (j < this.size) {
                    this.set(j, this.get(++j));
                }
            }
        }
        this.size--;
    }

    @Override
    public E[] toArray() {
        E[] arr = (E[]) new Object[this.size];
        for (int i = 0; i < this.size; i++) {
            arr[i] = this.get(i);
        }
        return arr;
    }

    @Override
    public E get(int i) {
        if (i < this.size)
            return this.items[i];
        return null;
    }

    @Override
    public void set(int i, E e) {
        if (i < this.capacity) {
            this.items[i] = e;
        } else throw new ArrayIndexOutOfBoundsException("Given index " + i + " should be less than " + this.capacity);
    }

    @Override
    public Iterator<E> iterator() {
        return new ArrayListIterator();
    }

    @Override
    public String toString() {
        StringBuilder out = new StringBuilder();
        out.append("[");
        var iter = this.iterator();
        while (iter.hasNext()) {
            out.append(iter.next());
            if (iter.hasNext())
                out.append(", ");
        }
        out.append("]");
        return out.toString();
    }

    private class ArrayListIterator implements Iterator<E> {

        private int currentIndex = 0;

        @Override
        public boolean hasNext() {
            return currentIndex !=  size;
        }

        @Override
        public E next() {
            return get(this.currentIndex++);
        }
    }
}
