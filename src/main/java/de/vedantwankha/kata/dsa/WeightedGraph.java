package de.vedantwankha.kata.dsa;

public class WeightedGraph<V extends Comparable<V>> extends UnweightedGraph<V> implements Graph<V>, Iterable<Map.Entry<V, List<Graph.Edge<V>>>> {
    public static class WeightedEdge<V extends Comparable<V>> extends UnweightedEdge<V> {
        private double weight;
        public WeightedEdge(V src, V dest, double w) {
            super(src, dest);
            weight = w;
        }

        @Override
        public String toString() {
            return "WeightedEdge{" +
                    "weight=" + weight +
                    ", src=" + src +
                    ", dest=" + dest +
                    '}';
        }
    }

    @Override
    public boolean addEdge(V u, V v) {
        return addEdge(new WeightedEdge<>(u, v, Double.NaN));
    }

    public boolean addEdge(V u, V v, double w) {
        return addEdge(new WeightedEdge<>(u, v, w));
    }
}
