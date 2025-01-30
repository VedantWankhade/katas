package de.vedantwankha.kata.dsa;

/**
 * Collection is a generic interface representing any kind of <b>immutable</b> collection / container.
 * It does not care if the items within are sequenced or unique. It is left to the lower level interfaces extending de.vedantwankha.kata.dsa.Collection.
 * It only declares the following behaviour:
 * <ul>
 *     <li>Check if element exists</li>
 *     <li>Check if the collection is empty</li>
 *     <li>Get size of the collection</li>
 *     <li>Get iterator [from {@link Iterable}]</li>
 * </ul>
 * @param <E> Type of element within
 */
public interface Collection<E> extends Iterable<E> {
    int size();
    boolean isEmpty();

    default Object[] toArray() {
//        E[] arr = (E[]) new Object[size()]; // this will throw ClassCastException somewhere else due to type erasure
        Object[] arr = new Object[size()];
        int i = 0;
        var itr = iterator();
        while (itr.hasNext()) {
            arr[i++] = itr.next();
        }
        return arr;
    }

    default boolean contains(E item) {
        for (E e: this) {
            if (e.equals(item)) return true;
        }
        return false;
    }
}
