package de.vedantwankha.kata.dsa;

public interface List<E> extends MutableCollection<E> {
    E get(int i);
    void set(int i, E e);
}
