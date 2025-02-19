package de.vedantwankha.kata.dsa;

import java.util.Iterator;

public class TreeMap<K extends Comparable<K>, V> implements Map<K, V>, Iterable<Map.Entry<K, V>> {
    @Override
    public Iterator<Map.Entry<K, V>> iterator() {
        return tree.iterator();
    }

    public static class Entry<K extends Comparable<K>, V> implements Map.Entry<K, V> {
        private final K key;
        private V val;

        public Entry(K key, V val) {
            this.val = val;
            this.key = key;
        }

        @Override
        public String toString() {
            return "Entry{" +
                    "key=" + key +
                    ", val=" + val +
                    '}';
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
    }

    public enum BSTType {
        BST, RBTREE
    }

    private final BSTType type;
    private BST<Map.Entry<K, V>> tree;

    public TreeMap(BSTType type) {
        this.type = type;
        tree = type == BSTType.BST ? new BST<>() : new RBTree<>();
    }

    @Override
    public int size() {
        return tree.size();
    }

    @Override
    public boolean isEmpty() {
        return tree.isEmpty();
    }

    @Override
    public void clear() {
        tree = type == BSTType.BST ? new BST<>() : new RBTree<>();
    }

    @Override
    public boolean containsKey(K key) {
        for (Map.Entry<K, V> e: this.tree) {
            if (key.compareTo(e.getKey()) == 0) return true;
        }
        return false;
    }

    @Override
    public V get(K key) {
        for (Map.Entry<K, V> e: this.tree) {
            if (key.compareTo(e.getKey()) == 0) return e.getValue();
        }
        return null;
    }

    @Override
    public void put(K key, V val) {
        var e = new Entry<K, V>(key, val);
        tree.add(e);
    }

    @Override
    public V remove(K key) {
        var v = this.get(key);
        this.tree.remove(new Entry<K, V>(key, null));
        return v;
    }

    @Override
    public Set<K> keySet() {
        // TODO))
        return null;
    }

    @Override
    public Collection<V> values() {
        ArrayList<V> list = new ArrayList<>();
        for (Map.Entry<K, V> e: this.tree) {
            list.add(e.getValue());
        }
        return list;
    }

    @Override
    public Set<Map.Entry<K, V>> entrySet() {
        // TODO))
        return null;
    }


    @Override
    public String toString() {
        StringBuilder out = new StringBuilder();
        out.append("[");
        var iter = this.iterator();
        while (iter.hasNext()) {
            out.append(iter.next());
            if (iter.hasNext())
                out.append(", ");
        }
        out.append("]");
        return out.toString();
    }
}
