package de.vedantwankha.kata.dsa;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class BST<E extends Comparable<E>> extends AbstractCollection<E> implements MutableCollection<E> {
    public static class Node<E> {
        private final E data;
        private Node<E> left;
        private Node<E> right;

        public E getData() {
            return data;
        }

        public Node(E data, Node<E> left, Node<E> right) {
            this.data = data;
            this.left = left;
            this.right = right;
        }

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
        var ret = new ArrayList<Node<E>>();
        if (root.data.compareTo(key) == 0) {
            ret.add(root);
            ret.add(null);
        } else {
            Node<E> n = root;
            Node<E> prev = null;
            while (n != null && n.data.compareTo(key) != 0) {
                prev = n;
                if (key.compareTo(n.data) <= 0) {
                    n = n.left;
                } else {
                    n = n.right;
                }
            }
            if (n != null) {
                ret.add(n);
                ret.add(prev);
                return ret;
            }
            ret.add(null);
            ret.add(null);
        }
        return ret;
    }

    @Override
    public void clear() {
        root = null;
        size = 0;
    }

    @Override
    public void remove(E key) {
        if (root == null) throw new NoSuchElementException("Empty tree");
        ImmutableList<Node<E>> l = get(key);
        var n = l.get(0);
        var p = l.get(1);
        if (n == null) {
            throw new NoSuchElementException("No such element");
        }
        if (n.left == null || n.right == null) {
            Node<E> maybeChild = null;
            // the node we want to delete is a node that is leaf node or only has one child node
            // get the only child even if it is null (in case of leaf)
            maybeChild = n.left == null ? n.right : n.left;

            if (p == null) {
                // that node is root
                root = maybeChild;
            } else if (key.compareTo(p.data) <= 0) {
                // that node is not root
                p.left = maybeChild;
            } else {
                p.right = maybeChild;
            }
        } else {
            Node<E> max, maxParent = null;
            ImmutableList<Node<E>> maxTuple = max(n.left);
            max = maxTuple.get(0);
            maxParent = maxTuple.get(1);
            Node<E> newNode;
            if (maxParent == null) {
                // left child of the node is the largest in node's left subtree
                newNode = new Node<E>(max.data, null, n.right);
            } else {
                newNode = new Node<>(max.data, n.left, n.right);
                maxParent.right = max.left;
            }
            if (p == null) {
                root = newNode;
            } else if (key.compareTo(p.data) <= 0) {
                p.left = newNode;
            } else {
                p.right = newNode;
            }
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

    public ImmutableList<Node<E>> max() {
        return max(root);
    }

    private ImmutableList<Node<E>> max(Node<E> node) {
        Node<E> current = node;
        Node<E> parent = null;
        while (current.right != null) {
            parent = current;
            current = current.right;
        }
        var ret = new ArrayList<Node<E>>();
        ret.add(current);
        ret.add(parent);
        return ret;
    }

    public ImmutableList<Node<E>> min() {
        var current = root;
        var parent = root;
        while (current.left != null) {
            parent = current;
            current = current.left;
        }
        var ret = new ArrayList<Node<E>>();
        ret.add(current);
        ret.add(parent);
        return ret;
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
