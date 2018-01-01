
import java.util.Iterator;

public class Deque<Item> implements Iterable<Item> {

    private Item[] arr;         // array of items
    private int n;            // number of elements on stack
    private int front;
    private int rear;

    // construct an empty deque
    public Deque() {
        int cap = 2;
        arr = (Item[]) new Object[cap];
        front = cap / 2;
        rear = cap / 2;
    }

    // is the deque empty?
    public boolean isEmpty() {
        return n == 0;
    }

    // return the number of items on the deque
    public int size() {
        return n;
    }

    // resize the underlying array holding the elements
    private void resize(int cap) {
        int middle = cap / 2;
        int newFront = middle - n / 2;
        int newRear = newFront + n;

        Item[] dest = (Item[]) new Object[cap];
        System.arraycopy(arr, front, dest, newFront, n);
        arr = dest;
        front = newFront;
        rear = newRear;
    }

    // add the item to the front
    public void addFirst(Item item) {
        if (item == null) {
            throw new java.lang.IllegalArgumentException();
        }

        if (front == 0) {
            resize(2 * arr.length);
        }

        arr[--front] = item;
        n++;
    }

    // add the item to the end
    public void addLast(Item item) {
        if (item == null) {
            throw new java.lang.IllegalArgumentException();
        }

        if (rear == arr.length) {
            resize(2 * arr.length);    // double size of array if necessary
        }

        arr[rear++] = item;
        n++;
    }

    // remove and return the item from the front
    public Item removeFirst() {
        if (isEmpty()) {
            throw new java.util.NoSuchElementException();
        }

        Item item = arr[front++];
        arr[front - 1] = null; // to avoid loitering
        n--;

        // shrink size of array if necessary
        if (n > 0 && n == arr.length / 4) {
            resize(arr.length / 2);
        }
        return item;
    }

    // remove and return the item from the end
    public Item removeLast() {
        if (isEmpty()) {
            throw new java.util.NoSuchElementException();
        }

        Item item = arr[--rear];
        if (rear + 1 < arr.length) {
            arr[rear + 1] = null; // to avoid loitering
        }
        n--;

        // shrink size of array if necessary
        if (n > 0 && n == arr.length / 4) {
            resize(arr.length / 2);
        }
        return item;
    }

    // return an iterator over items in order from front to end
    public Iterator<Item> iterator() {
        return new DequeForwardIterator();
    }

    private class DequeForwardIterator implements Iterator<Item> {
        private Item[] items;
        private int i;

        public DequeForwardIterator() {
            items = (Item[]) new Object[n]; 
            System.arraycopy(arr, front, items, 0, n);
            i = 0;
        }

        public boolean hasNext() {
            return i < items.length;
        }

        public void remove() {
            throw new java.lang.UnsupportedOperationException();
        }

        public Item next() {
            if (!hasNext()) {
                throw new java.util.NoSuchElementException();
            }
            return items[i++];
        }
    }

}
