package de.vedantwankha.kata.dsa;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.io.File;
import java.util.Iterator;

public class GraphTest {
    @Nested
    @DisplayName("Test unweighted graphs")
    class UnweightedGraphTest {
        String[] vertices = {"Seattle", "San Francisco", "Los Angeles",
                "Denver", "Kansas City", "Chicago", "Boston", "New York",
                "Atlanta", "Miami", "Dallas", "Houston"};
        int[][] edges = {
                {0, 1}, {0, 3}, {0, 5},
                {1, 0}, {1, 2}, {1, 3},
                {2, 1}, {2, 3}, {2, 4}, {2, 10},
                {3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5},
                {4, 2}, {4, 3}, {4, 5}, {4, 7}, {4, 8}, {4, 10},
                {5, 0}, {5, 3}, {5, 4}, {5, 6}, {5, 7},
                {6, 5}, {6, 7},
                {7, 4}, {7, 5}, {7, 6}, {7, 8},
                {8, 4}, {8, 7}, {8, 9}, {8, 10}, {8, 11},
                {9, 8}, {9, 11},
                {10, 2}, {10, 4}, {10, 8}, {10, 11},
                {11, 8}, {11, 9}, {11, 10}
        };
        UnweightedGraph<String> g = new UnweightedGraph<>();

        @BeforeEach
        void init() {
            for (String v : vertices) {
                g.addVertex(v);
            }
            for (var e : edges) {
                g.addEdge(vertices[e[0]], vertices[e[1]]);
            }
        }

        @Test
        @DisplayName("Test add vertex and add edge")
        void testAddVertexEdge() {
            System.out.println(g);
        }

        @Test
        @DisplayName("Test add vertex and add edge")
        void testGetNeighbors() {
            for (var v : vertices) {
                System.out.println(v + ": " + g.getNeighbors(v));
            }
        }

        @Test
        @DisplayName("Test bfs path")
        void testBfsPath() {
            System.out.println(g.path("Seattle", "Seattle"));
            System.out.println(g.path("Seattle", "San Francisco"));
            System.out.println(g.path("Seattle", "Denver"));
            System.out.println(g.path("Seattle", "Boston"));
            System.out.println(g.path("Seattle", "Chicago"));
        }

        @Test
        @DisplayName("Test remove edge")
        void testRemoveEdge() {
            System.out.println(g.getNeighbors("Seattle"));
            System.out.println(g.getNeighbors("Denver"));
            g.removeEdge("Seattle", "Denver");
            System.out.println(g.getNeighbors("Seattle"));
            System.out.println(g.getNeighbors("Denver"));
        }

        @Test
        @DisplayName("Test remove vertex")
        void testRemoveVertex() {
            System.out.println(g);
            System.out.println(g.vertices());
            g.removeVertex("Seattle");
            System.out.println(g.vertices());
            System.out.println(g);
        }

        @Test
        @DisplayName("Test breadth first traversal")
        void testBreadthFirstIterator() {
            Iterator<String> it = ((UnweightedGraph) g).iterator("Seattle");
            while (it.hasNext()) {
                System.out.println(it.next());
            }
        }

        @Test
        @DisplayName("Test depth first traversal")
        void testDepthFirstIterator() {
            Iterator<String> it = ((UnweightedGraph) g).dfsIterator("Seattle");
            while (it.hasNext()) {
                System.out.println(it.next());
            }
        }
    }

    @Nested
    @DisplayName("Test weighted graphs")
    class WeightedGraphTest {
        String[] vertices = {"Seattle", "San Francisco", "Los Angeles",
                "Denver", "Kansas City", "Chicago", "Boston", "New York",
                "Atlanta", "Miami", "Dallas", "Houston"};
        int[][] edges = {
                {0, 1}, {0, 3}, {0, 5},
                {1, 0}, {1, 2}, {1, 3},
                {2, 1}, {2, 3}, {2, 4}, {2, 10},
                {3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5},
                {4, 2}, {4, 3}, {4, 5}, {4, 7}, {4, 8}, {4, 10},
                {5, 0}, {5, 3}, {5, 4}, {5, 6}, {5, 7},
                {6, 5}, {6, 7},
                {7, 4}, {7, 5}, {7, 6}, {7, 8},
                {8, 4}, {8, 7}, {8, 9}, {8, 10}, {8, 11},
                {9, 8}, {9, 11},
                {10, 2}, {10, 4}, {10, 8}, {10, 11},
                {11, 8}, {11, 9}, {11, 10}
        };
        WeightedGraph<String> g = new WeightedGraph<>();

        @BeforeEach
        void init() {
            for (String v : vertices) {
                g.addVertex(v);
            }
            for (var e : edges) {
                g.addEdge(vertices[e[0]], vertices[e[1]], Math.random());
            }
        }

        @Test
        @DisplayName("Test add vertex and add edge")
        void testAddVertexEdge() {
            System.out.println(g);
        }

        @Test
        @DisplayName("Test add vertex and add edge")
        void testGetNeighbors() {
            for (var v : vertices) {
                System.out.println(v + ": " + g.getNeighbors(v));
            }
        }

        @Test
        @DisplayName("Test remove edge")
        void testRemoveEdge() {
            System.out.println(g.getNeighbors("Seattle"));
            System.out.println(g.getNeighbors("Denver"));
            g.removeEdge("Seattle", "Denver");
            System.out.println(g.getNeighbors("Seattle"));
            System.out.println(g.getNeighbors("Denver"));
        }

        @Test
        @DisplayName("Test remove vertex")
        void testRemoveVertex() {
            System.out.println(g);
            System.out.println(g.vertices());
            g.removeVertex("Seattle");
            System.out.println(g.vertices());
            System.out.println(g);
        }
    }
}