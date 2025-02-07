package de.vedantwankha.kata.dsa;

import javax.naming.OperationNotSupportedException;

public interface MutableCollection<E> extends Collection<E> {
    void add(E item);
    void clear();
//    void remove(); TODO)) think about this - do I need this?
    void remove(E item);

    default void addAll(Collection<? extends E> c) {
        for (E e : c) {
            this.add(e);
        }
    }
}
