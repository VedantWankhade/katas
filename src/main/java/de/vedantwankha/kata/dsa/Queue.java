package de.vedantwankha.kata.dsa;

public interface Queue<E> extends Collection<E> {
    void pushLast(E e);
    E popFirst();
}
