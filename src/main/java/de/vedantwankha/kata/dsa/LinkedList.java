package de.vedantwankha.kata.dsa;

public class LinkedList<E> implements List<E>, ImmutableList<E> {
    private class Node {
        private E data;
        private Node next;

        Node(E e) {
            this.data = e;
            this.next = null;
        }
    }

    private Node head;
    private int size;

    public LinkedList() {
        this.head = null;
        this.size = 0;
    }

    public LinkedList(Collection<? extends E> c) {
        this();
        for (E e: c) {
            add(e);
        }
    }

    @Override
    public void add(E e) {
        if (this.size == 0)
            this.head = new Node(e);
        else {

        }
    }
}
