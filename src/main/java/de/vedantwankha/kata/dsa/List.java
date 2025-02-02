package de.vedantwankha.kata.dsa;

public interface List<E> extends MutableCollection<E> {
    E get(int i);
    void add(E e, int i);
    void addFirst(E e);
    void addLast(E e);
    void remove(int idx);
    void set(int i, E e);
}
