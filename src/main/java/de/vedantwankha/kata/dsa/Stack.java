package de.vedantwankha.kata.dsa;

public interface Stack<E> extends Collection<E> {
    E pop();
    void push(E e);
}
