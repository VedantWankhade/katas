package de.vedantwankha.kata.dsa;

public interface MutableCollection<E> extends Collection<E> {
    void add(E item);
    void clear();
    void remove(E item);

    default void addAll(Collection<? extends E> c) {
        for (E e : c) {
            this.add(e);
        }
    }
}
