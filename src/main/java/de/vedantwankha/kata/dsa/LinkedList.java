package de.vedantwankha.kata.dsa;

import java.util.EmptyStackException;
import java.util.Iterator;

public class LinkedList<E> extends AbstractCollection<E> implements List<E>, ImmutableList<E>, Stack<E> {
    private class Node {
        private E data;
        private Node next;
        private Node prev;

        Node(E e) {
            data = e;
            next = null;
            prev = null;
        }
    }

    private Node head;
    private Node tail;
    private int size;

    public LinkedList() {
        head = null;
        tail = null;
        size = 0;
    }

    public LinkedList(Collection<? extends E> c) {
        this();
        for (E e: c) {
            this.add(e);
        }
    }

    @Override
    public void add(E e) {
        addLast(e);
    }

    @Override
    public void add(E e, int idx) {
        if (idx > size) {
            throw new IndexOutOfBoundsException("Index " + idx + " should be less than or equal to the size of list " + size);
        } else if (idx == size) {
            addLast(e);
            return;
        } else if (idx == 0) {
            addFirst(e);
            return;
        }

        var currentNode = head;
        int i = 0;
        while (currentNode.next != null && i < idx - 1) {
            i++;
            currentNode = currentNode.next;
        }
        var newNode = new Node(e);
        newNode.next = currentNode.next;
        currentNode.next = newNode;
        size++;
    }

    @Override
    public void addFirst(E e) {
        var newNode = new Node(e);
        if (head == null) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.next = head;
            head = newNode;
        }
        size++;
    }

    @Override
    public void addLast(E e) {
        var node = new Node(e);
        if (head == null) {
            head = node;
            tail = node;
        } else {
            tail.next = node;
            tail = node;
        }
        size++;
    }

    @Override
    public void clear() {
        head = tail = null;
        size = 0;
    }

    @Override
    public E get(int idx) {
        if (head == null) throw new IndexOutOfBoundsException("List is empty");
        if (idx >= size) throw new IndexOutOfBoundsException("Given index " + idx + " should be less than size of the list " + size);
        int i = 0;
        var currentNode = head;
        while (currentNode.next != null && i < idx) {
            i++;
            currentNode = currentNode.next;
        }
        if (i == idx) return currentNode.data;
        else throw new IndexOutOfBoundsException("Given index " + idx + " should be less than size of the list " + size);
    }

    @Override
    public void remove(int idx) {
        if (idx >= size()) throw new IndexOutOfBoundsException("Given index " + idx + " should be less than the size of list " + size());
        if (idx == 0) {
            head = head.next;
            size--;
            return;
        }
        var currentNode = head;
        int i = 0;
        while (currentNode.next != null && i < idx - 1) {
            i++;
            currentNode = currentNode.next;
        }
        currentNode.next = currentNode.next.next;
        size--;
    }

    @Override
    public void remove(E item) {
        // TODO))
    }

    @Override
    public void set(int i, E e) {
        // TODO))
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
    public E pop() {
        if (size() == 0) throw new EmptyStackException();
        E e = get(0);
        remove(0);
        return e;
    }

    @Override
    public void push(E e) {
        addFirst(e);
    }

    @Override
    public Iterator<E> iterator() {
        return new LinkedListIterator();
    }

    private class LinkedListIterator implements Iterator<E> {

        private Node currentNode = head;
        private int cursor = 0;

        @Override
        public boolean hasNext() {
            if (cursor >= size)
                return false;
            return true;
        }

        @Override
        public E next() {
            var ret = currentNode.data;
            currentNode = currentNode.next;
            cursor++;
            return ret;
        }
    }
}
