package de.vedantwankha.kata.dsa;

public interface Deque<E> extends Queue<E> {
    E popLast();
    void pushFirst(E e);
}
