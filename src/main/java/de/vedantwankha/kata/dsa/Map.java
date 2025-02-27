package de.vedantwankha.kata.dsa;

public interface Map<K extends Comparable<K>, V> extends Iterable<Map.Entry<K, V>> {
    interface Entry<K extends Comparable<K>, V> extends Comparable<Entry<K, V>> {
        K getKey();
        V getValue();
        void setValue(V val);
    }
    int size();
    boolean isEmpty();
    void clear();
    boolean containsKey(K key);
    V get(K key);
    void put(K key, V val);
    V remove(K key);
    Set<K> keySet();
    List<V> values();
    Set<Entry<K, V>> entrySet();
}
