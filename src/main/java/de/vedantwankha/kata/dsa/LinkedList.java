package de.vedantwankha.kata.dsa;

import java.util.EmptyStackException;
import java.util.Iterator;

public class LinkedList<E> extends AbstractCollection<E> implements List<E>, ImmutableList<E>, Stack<E>, Deque<E> {
    private static class Node<E> {
        private E data;
        private Node<E> next;
        private Node<E> prev;

        Node(E e) {
            data = e;
            next = null;
            prev = null;
        }
    }

    private Node<E> head;
    private Node<E> tail;
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
    public E popLast() {
        E e = get(size() - 1);
        tail = tail.prev;
        size--;
        return e;
    }

    @Override
    public void pushFirst(E e) {
        push(e);
    }

    @Override
    public void pushLast(E e) {
        addLast(e);
    }

    @Override
    public E popFirst() {
        return pop();
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
        var newNode = new Node<E>(e);
        newNode.prev = currentNode.prev;
        newNode.next = currentNode.next;
        currentNode.prev = newNode;
        currentNode.next = newNode;
        size++;
    }

    @Override
    public void addFirst(E e) {
        var newNode = new Node<E>(e);
        if (head == null) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.next = head;
            head.prev = newNode;
            head = newNode;
        }
        size++;
    }

    @Override
    public void addLast(E e) {
        var node = new Node<E>(e);
        if (head == null) {
            head = node;
            tail = node;
        } else {
            tail.next = node;
            node.prev = tail;
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
    public E getFirst() {
        return this.get(0);
    }

    @Override
    public E getLast() {
        return this.get(this.size - 1);
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
        assert currentNode.next != null;
        currentNode.next = currentNode.next.next;
        size--;
    }

    @Override
    public void remove(E item) {
        // remove first occurrence
        if (head == null) throw new RuntimeException("Empty List");
        if (head.data.equals(item)) {
            head = head.next;
        } else {
            var current = head;
            while (current.next != null) {
                if (current.next.data.equals(item)) {
                    var toRemove = current.next;
                    current.next = toRemove.next;
                    if (toRemove.next != null) {
                        toRemove.next.prev = current;
                    } else {
                        tail = current;
                    }
                    size--;
                    break;
                }
                current = current.next;
            }
        }
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

        private Node<E> currentNode = head;
        private int cursor = 0;

        @Override
        public boolean hasNext() {
            if (currentNode == null)
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
