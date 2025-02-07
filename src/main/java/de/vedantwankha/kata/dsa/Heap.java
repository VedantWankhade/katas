package de.vedantwankha.kata.dsa;

import java.util.Comparator;
import java.util.Iterator;

public class Heap<E extends Comparable<E>> extends AbstractCollection<E> implements MutableCollection<E> {
    private final ArrayList<E> heap;
    private final Comparator<? super E> comparator;

    public Heap(Comparator<? super E> comparator) {
        this.heap = new ArrayList<>();
        this.comparator = comparator;
    }

    public Heap() {
        this((e1, e2) -> e1.compareTo(e2));
    }

    @Override
    public void add(E item) {
        heap.addLast(item);
        int currentIndex = heap.size() - 1;
        int parentIndex = (currentIndex - 1) / 2;
        while (currentIndex > 0 && comparator.compare(item, heap.get(parentIndex)) > 0) {
            heap.set(currentIndex, heap.get(parentIndex));
            heap.set(parentIndex, item);
            currentIndex = parentIndex;
        }
    }

    public E extract() {
        E ret = this.heap.getFirst();
        this.heap.set(0, this.heap.get(this.size() - 1));
        this.heap.remove(this.size() - 1);
        int currentIndex = 0;
        while(true) {
            int leftChildIndex = 2 * currentIndex + 1;
            int rightChildIndex = 2 * currentIndex + 2;
            int largestIndex = currentIndex;
            if (leftChildIndex < this.size() && comparator.compare(this.heap.get(largestIndex), this.heap.get(leftChildIndex)) < 0) {
                largestIndex = leftChildIndex;
            }
            if (rightChildIndex < this.size() && comparator.compare(this.heap.get(largestIndex), this.heap.get(rightChildIndex)) < 0) {
                largestIndex = rightChildIndex;
            }
            if (largestIndex == currentIndex) break;
            E temp = this.heap.get(largestIndex);
            this.heap.set(largestIndex, this.heap.get(currentIndex));
            this.heap.set(currentIndex, temp);
            currentIndex = largestIndex;
        }
        return ret;
    }

    @Override
    public void clear() {
        this.heap.clear();
    }

    @Override
    public void remove(E item) {
        throw new UnsupportedOperationException("Method remove not supported");
    }

    @Override
    public int size() {
        return this.heap.size();
    }

    @Override
    public boolean isEmpty() {
        return this.heap.isEmpty();
    }

    @Override
    public Object[] toArray() {
        return this.heap.toArray();
    }

    @Override
    public boolean contains(E item) {
        return this.heap.contains(item);
    }

    @Override
    public Iterator<E> iterator() {
        return this.heap.iterator();
    }

}
