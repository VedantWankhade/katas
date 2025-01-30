package de.vedantwankha.kata.dsa;

abstract public class AbstractCollection<E> implements Collection<E> {
    @Override
     public Object[] toArray() {
//      E[] arr = (E[]) new Object[size()]; // this will throw ClassCastException somewhere else due to type erasure
        Object[] arr = new Object[size()];
        int i = 0;
        var itr = iterator();
        while (itr.hasNext()) {
            arr[i++] = itr.next();
        }
        return arr;
    }

    @Override
    public boolean contains(E item) {
        for (E e: this) {
            if (e.equals(item)) return true;
        }
        return false;
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
