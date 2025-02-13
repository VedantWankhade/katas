package de.vedantwankha.kata.dsa;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class BST<E extends Comparable<E>> extends AbstractCollection<E> implements MutableCollection<E> {
    public static class Node<E> {
        private E data;
        private Node<E> left;
        private Node<E> right;

        Node(E e) {
            this.data = e;
            left = null;
            right = null;
        }

        @Override
        public String toString() {
            return "Node{" +
                    "data=" + data +
                    '}';
        }
    }

    private Node<E> root;
    private int size;

    public BST() {
        root = null;
        size = 0;
    }

    public BST(Collection<? extends E> c) {
        addAll(c);
    }

    @Override
    public void add(E e) {
        if (size == 0) {
            root = new Node<E>(e);
        } else {
            var currNode = root;
            while (true) {
                if (e.compareTo(currNode.data) > 0) {
                    if (currNode.right == null) {
                        currNode.right = new Node<E>(e);
                        break;
                    } else {
                        currNode = currNode.right;
                    }
                } else {
                    if (currNode.left == null) {
                        currNode.left = new Node<E>(e);
                        break;
                    } else {
                        currNode = currNode.left;
                    }
                }
            }
        }
        size++;
    }

    public ImmutableList<Node<E>> get(E key) {
        var n = root;
        var prev = root;
        while (n != null && n.data.compareTo(key) != 0) {
            prev = n;
            if (key.compareTo(n.data) <= 0) {
                n = n.left;
            } else {
                n = n.right;
            }
        }
        var ret = new ArrayList<Node<E>>();
        if (n != null) {
            ret.add(n);
            ret.add(prev);
            return ret;
        }
        ret.add(null);
        ret.add(null);
        return ret;
    }

    @Override
    public void clear() {
        root = null;
        size = 0;
    }

    @Override
    public void remove(E key) {
        // TODO))
        ImmutableList<Node<E>> l = get(key);
        var n = l.get(0);
        var p = l.get(1);
        if (n == null) {
            throw new NoSuchElementException("No such element");
        }
        if (p == null) {

        }
    }

    @Override
    public int size() {
        return this.size;
    }

    @Override
    public boolean isEmpty() {
        return size == 0;
    }

    public E max() {
        var current = root;
        while (current.right != null) {
            current = current.right;
        }
        return current.data;
    }

    public E min() {
        var current = root;
        while (current.left != null) {
            current = current.left;
        }
        return current.data;
    }

    @Override
    public Iterator<E> iterator() {
        return new InOrderIterator();
    }

    private class InOrderIterator implements Iterator<E> {
        private Stack<Node<E>> stack = new LinkedList<>();

        private InOrderIterator() {
            pushLeftSubtree(root);
        }

        private void pushLeftSubtree(Node<E> n) {
            while (n != null) {
                stack.push(n);
                n = n.left;
            }
        }

        @Override
        public boolean hasNext() {
            return !stack.isEmpty();
        }

        @Override
        public E next() {
            Node<E> n = stack.pop();
            if (n.right != null) {
                pushLeftSubtree(n.right);
            }
            return n.data;
        }
    }
}
