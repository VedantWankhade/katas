package de.vedantwankha.kata.dsa;

import java.util.Iterator;

public class HashSet<E extends Comparable<E>> implements Set<E> {
    private HashMap<E, Boolean> map;

    public HashSet(int initialCapacity) {
        map = new HashMap<>(initialCapacity);
    }

    public HashSet() {
        this(8);
    }

    @Override
    public void add(E item) {
        map.put(item, true); // could also set null to save a bit of space instead of a bool value
    }

    @Override
    public void clear() {
        map.clear();
    }

    @Override
    public void remove(E item) {
        map.remove(item);
    }

    @Override
    public int size() {
        return map.size();
    }

    @Override
    public boolean isEmpty() {
        return map.isEmpty();
    }

    @Override
    public Object[] toArray() {
        return new Object[0];
    }

    @Override
    public boolean contains(E item) {
        return map.containsKey(item);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        var iter = iterator();
        sb.append("[");
        while (iter.hasNext()) {
            sb.append(iter.next());
            sb.append(iter.hasNext() ? ", " : "");
        }
        sb.append("]");
        return sb.toString();
    }

    @Override
    public Iterator<E> iterator() {
        return new HashSetIterator();
    }

    private class HashSetIterator implements Iterator<E> {
        private final Iterator<Map.Entry<E, Boolean>> mapIterator;

        HashSetIterator() {
            mapIterator = map.iterator();
        }

        @Override
        public boolean hasNext() {
            return mapIterator.hasNext();
        }

        @Override
        public E next() {
            return mapIterator.next().getKey();
        }
    }
}
