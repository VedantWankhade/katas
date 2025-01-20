package de.vedantwankha.kata.dsa;

/**
 * Collection is a generic interface representing any kind of collection / container.
 * It does not care if the items within are sequenced or unique. It is left to the lower level interfaces extending de.vedantwankha.kata.dsa.Collection.
 * It only declares the following behaviour:
 * <ul>
 *     <li>Adding element</li>
 *     <li>Add all elements from another {@link Collection}</li>
 *     <li>Removing element</li>
 *     <li>Seeing if element exists</li>
 *     <li>Getting size of the collection</li>
 *     <li>Removing all elements</li>
 *     <li>Give iterator [from {@link Iterable}]</li>
 * </ul>
 * @param <E> Type of element within
 */
public interface Collection<E> extends Iterable<E> {
    int size();
    boolean isEmpty();
    boolean contains(E e);
    void add(E e);
    void addAll(Collection<? extends E> c);
    void remove(E e);
    E[] toArray();
    E get(int i);
    void set(int i, E e);
}
