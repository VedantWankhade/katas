package de.vedantwankha.kata.dsa;

public interface Graph<V extends Comparable<V>> extends Iterable<V> {
    /** Returns the number of vertices */
    int size();

    /** Returns the vertices in the graph */
    Set<V> vertices();

    /** Return the object for the specified vertex index */
    V getVertex(int index);

    /** Return the index for the specified vertex object */
//    int getIndex(V v);

    /** Return teh neighbors of the vertex of specified index */
    Set<V> getNeighbors(V vertex);

    /** Return the degree for a specified vertex */
    int getDegree(V vertex);

    /** Return string representation of edges */
//    String edgesToString();

    /** Clear the graph */
    void clear();

    /** Add a vertex to the graph */
    boolean addVertex(V vertex);

    /** Add an edge to the graph */
    boolean addEdge(Edge<V> e);

    /** Add an edge to the graph */
    boolean addEdge(V u, V v);

    /** Remove a vertex v from the graph */
    boolean removeVertex(V v);

    /** Remove an edge (u, v) from the graph */
    boolean removeEdge(V u, V v);

    /** Find the shortest path from src to dest vertex */
    List<V> path(V src, V dest);

    interface Edge<V> {
        V src();
        V dest();
    }
}
