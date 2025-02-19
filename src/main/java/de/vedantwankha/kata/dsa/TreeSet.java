package de.vedantwankha.kata.dsa;

public class TreeSet<E extends Comparable<E>> extends BST<E> implements Set<E> {
    @Override
    public void add(E e) {
        if (!this.contains(e)) super.add(e);
    }
}
