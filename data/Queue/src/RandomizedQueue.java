
import edu.princeton.cs.algs4.StdRandom;
import java.util.Iterator;

public class RandomizedQueue<Item> implements Iterable<Item> {

    private Item[] arr;         // array of items
    private int n;            // number of elements on stack

    // construct an empty randomized queue
    public RandomizedQueue() {
        arr = (Item[]) new Object[2];
        n = 0;
    }

    // is the randomized queue empty?
    public boolean isEmpty() {
        return n == 0;
    }

    // return the number of items on the randomized queue
    public int size() {
        return n;
    }

    // resize the underlying array holding the elements
    private void resize(int capacity) {
        arr = java.util.Arrays.copyOf(arr, capacity);
    }

    // add the item
    public void enqueue(Item item) {
        if (item == null) {
            throw new java.lang.IllegalArgumentException();
        }

        if (n == arr.length) {
            resize(2 * arr.length);    // double size of array if necessary
        }

        arr[n++] = item;  // add item
    }

    // remove and return a random item
    public Item dequeue() {
        if (isEmpty()) {
            throw new java.util.NoSuchElementException();
        }

        int pick = StdRandom.uniform(n);

        Item item = arr[pick];
        arr[pick] = arr[n - 1]; // swap with last element
        arr[n - 1] = null;      // to avoid loitering
        n--;

        // shrink size of array if necessary
        if (n > 0 && n == arr.length / 4) {
            resize(arr.length / 2);
        }
        return item;
    }

    // return a random item (but do not remove it)
    public Item sample() {
        if (isEmpty()) {
            throw new java.util.NoSuchElementException();
        }

        int pick = StdRandom.uniform(n);
        return arr[pick];
    }

    // return an independent iterator over items in random order
    public Iterator<Item> iterator() {
        return new RandomizedQueueIterator();
    }

    private class RandomizedQueueIterator implements Iterator<Item> {

        private Item[] shuffle;
        private int i;

        public RandomizedQueueIterator() {
            shuffle = java.util.Arrays.copyOf(arr, n);
            i = n - 1;
        }

        public boolean hasNext() {
            return i >= 0;
        }

        public void remove() {
            throw new java.lang.UnsupportedOperationException();
        }

        public Item next() {
            if (!hasNext()) {
                throw new java.util.NoSuchElementException();
            }

            int pick = StdRandom.uniform(i + 1);

            Item item = shuffle[pick];
            // swap with i'th element
            shuffle[pick] = shuffle[i];
            shuffle[i] = item;

            i--;
            return item;
        }
    }
}
