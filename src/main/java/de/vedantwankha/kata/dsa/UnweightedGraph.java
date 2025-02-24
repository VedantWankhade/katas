package de.vedantwankha.kata.dsa;

import java.util.Iterator;

public class UnweightedGraph<V extends Comparable<V>> implements Graph<V>, Iterable<Map.Entry<V, List<Graph.Edge<V>>>> {
    public static class UnweightedEdge<V extends Comparable<V>> implements Edge<V> {
        protected final V src;
        protected final V dest;

        public UnweightedEdge(V src, V dest) {
            this.src = src;
            this.dest = dest;
        }

        @Override
        public V src() {
            return src;
        }

        @Override
        public V dest() {
            return dest;
        }

        @Override
        public String toString() {
            return "UnweightedEdge{" +
                    "src=" + src +
                    ", dest=" + dest +
                    '}';
        }

        @Override
        public boolean equals(Object o) {
            if (o instanceof UnweightedGraph.UnweightedEdge<?>) {
                return ((UnweightedEdge<V>) o).src().compareTo(src) == 0 && ((UnweightedEdge<V>) o).dest().compareTo(dest) == 0;
            }
            throw new ClassCastException("Not a object of UnweightedEdge");
        }
    }

    protected Map<V, List<Edge<V>>> adjList;
    private int size;

    public UnweightedGraph(int initialCapacity) {
        size = 0;
        adjList = new HashMap<>(initialCapacity);
    }

    public UnweightedGraph() {
        this(8);
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public Set<V> vertices() {
        return adjList.keySet();
    }

    @Override
    public V getVertex(int index) {
        return null;
    }

    @Override
    public Set<V> getNeighbors(V vertex) {
        Set<V> neighbors = new HashSet<>();
        var l = adjList.get(vertex);
        for (Edge<V> e: l) {
            neighbors.add(e.dest());
        }
        return neighbors;
    }

    @Override
    public int getDegree(V vertex) {
        return 0;
    }

    @Override
    public void clear() {
        size = 0;
        adjList.clear();
    }

    @Override
    public boolean addVertex(V vertex) {
        if (adjList.containsKey(vertex)) return false;
        adjList.put(vertex, new LinkedList<>());
        size++;
        return true;
    }

    @Override
    public boolean addEdge(Edge<V> e) {
        if (!adjList.containsKey(e.src())) throw new RuntimeException("Vertex " + e.src() + " does not exist");
        if (!adjList.containsKey(e.dest())) throw new RuntimeException("Vertex " + e.dest() + " does not exist");

        adjList.get(e.src()).add(e);
        return true;
    }

    @Override
    public boolean addEdge(V u, V v) {
        return addEdge(new UnweightedEdge<>(u, v));
    }

    @Override
    public boolean removeVertex(V v) {
        if (!adjList.containsKey(v)) throw new RuntimeException("Vertex " + v + " does not exist");
        adjList.remove(v);
        for (var edgeList: adjList) {
            edgeList.getValue().remove(new UnweightedEdge<>(edgeList.getKey(), v));
        }
        size--;
        return true;
    }

    @Override
    public boolean removeEdge(V u, V v) {
        if (!adjList.containsKey(u)) throw new RuntimeException("Vertex " + u + " does not exist");
        if (!adjList.containsKey(v)) throw new RuntimeException("Vertex " + v + " does not exist");

        var bucket = adjList.get(u);
        bucket.remove(new UnweightedEdge<>(u, v));
        return true;
    }

    @Override
    public Iterator<Map.Entry<V, List<Edge<V>>>> iterator() {
        return adjList.iterator();
    }

    @Override
    public String toString() {
        return adjList.toString();
    }

//    @Override
//    public UnweightedGraph<V>.SearchTree dfs(V v) {
//        return null;
//    }
//
//    @Override
//    public UnweightedGraph<V>.SearchTree bfs(V v) {
//        return null;
//    }
}
