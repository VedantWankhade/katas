package de.vedantwankha.kata.dsa;

import java.util.Iterator;

public class HashMap<K extends Comparable<K>, V> implements Map<K, V> {
    public static class Entry<K extends Comparable<K>, V> implements Map.Entry<K, V>, Comparable<Map.Entry<K, V>> {

        private final K key;
        private V val;

        public Entry(K key, V val) {
            this.key = key;
            this.val = val;
        }

        @Override
        public K getKey() {
            return key;
        }

        @Override
        public V getValue() {
            return val;
        }

        @Override
        public void setValue(V val) {
            this.val = val;
        }

        @Override
        public int compareTo(Map.Entry<K, V> o) {
            return this.key.compareTo(o.getKey());
        }

        @Override
        public boolean equals(Object o) {
            return compareTo((Entry<K, V>) o) == 0;
        }

        @Override
        public String toString() {
            return "Entry{" +
                    "key=" + key +
                    ", val=" + val +
                    '}';
        }
    }

    private LinkedList<Map.Entry<K, V>>[] table;
    private final float LOAD_FACTOR = 0.75f;
    private int capacity;
    private int size;

    public HashMap(int initialCapacity) {
        capacity = initialCapacity;
        size = 0;
        table = new LinkedList[capacity];
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    @Override
    public void clear() {
        size = 0;
        table = new LinkedList[capacity];
    }

    private int hash(K key) {
        return supplementalHash(key.hashCode()) & (capacity - 1);
    }

    private int supplementalHash(int h) {
        h ^= (h >>> 20) ^ (h >>> 12);
        return h ^ (h >>> 7) ^ (h >>> 4);
    }

    @Override
    public boolean containsKey(K key) {
        return get(key) != null;
    }

    @Override
    public V get(K key) {
        int idx = hash(key);
        if (table[idx] == null) {
            return null;
        }
        for (var e: table[idx]) {
            if (key.compareTo(e.getKey()) == 0) return e.getValue();
        }
        return null;
    }

    private void rehash() {
        Set<Map.Entry<K, V>> entries = entrySet();
        capacity *= 2;
        size = 0;
        table = new LinkedList[capacity];
        for (var e: entries) {
            put(e.getKey(), e.getValue());
        }
    }

    @Override
    public void put(K key, V val) {
        if (containsKey(key)) throw new RuntimeException("Entry with key " + key + " already exists");
        if ((float) size / capacity >= LOAD_FACTOR) {
            rehash();
        }
        int idx = hash(key);
        if (table[idx] == null) {
            table[idx] = new LinkedList<Map.Entry<K, V>>();
        }
        table[idx].add(new Entry<>(key, val));
        size++;
    }

    @Override
    public V remove(K key) {
        int idx = hash(key);
        V ret = null;
        if (table[idx] != null) {
            for (var e: table[idx]) {
                if (e.getKey().compareTo(key) == 0) {
                    ret = e.getValue();
                }
            }
            table[idx].remove(new Entry<>(key, null));
        }
        return ret;
    }

    @Override
    public Set<K> keySet() {
        Set<K> keys = new HashSet<>(size);
        for (var bucket: table) {
            if (bucket != null) {
                for (var e: bucket) {
                    keys.add(e.getKey());
                }
            }
        }
        return keys;
    }

    @Override
    public List<V> values() {
        ArrayList<V> vals = new ArrayList<>(size);
        for (var bucket: table) {
            if (bucket != null) {
                for (var e: bucket) {
                    vals.add(e.getValue());
                }
            }
        }
        return vals;
    }

    @Override
    public Set<Map.Entry<K, V>> entrySet() {
        Set<Map.Entry<K, V>> entries = new HashSet<>();
        for (var bucket: table) {
            if (bucket != null) {
                for (var e : bucket) {
                    entries.add(e);
                }
            }
        }
        return entries;
    }

    public String toStringChain() {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < table.length; i++) {
            sb.append(i).append(" ");
            if (table[i] == null) {
                sb.append("null");
            } else {
                String s = table[i].toString();
                sb.append(s);
            }
            sb.append("\n");
        }
        return sb.toString();
    }

    @Override
    public String toString() {
        return toStringChain();
    }

    @Override
    public Iterator<Map.Entry<K, V>> iterator() {
        return new HashMapIterator();
    }

    private class HashMapIterator implements Iterator<Map.Entry<K, V>> {
        private int tableCursor;
        private LinkedList<Map.Entry<K, V>> currentBucket;
        private Iterator<Map.Entry<K, V>> bucketIterator;

        HashMapIterator() {
            tableCursor = 0;
            currentBucket = table[tableCursor];
            bucketIterator = currentBucket == null ? null : currentBucket.iterator();
        }

        @Override
        public boolean hasNext() {
            if (bucketIterator != null && bucketIterator.hasNext()) {
                return true;
            } else {
                if (tableCursor + 1 < table.length) {
                    tableCursor++;
                    currentBucket = table[tableCursor];
                    bucketIterator = currentBucket == null ? null : currentBucket.iterator();
                    return hasNext();
                } else return false;
            }
        }

        @Override
        public Map.Entry<K, V> next() {
            return bucketIterator.next();
        }
    }
}