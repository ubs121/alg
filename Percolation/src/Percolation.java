
import edu.princeton.cs.algs4.WeightedQuickUnionUF;

// Percolation data type
public class Percolation {

    private final int n;
    private final int size;
    private boolean[] open; // open[i] - is site 'i' open?
    private final WeightedQuickUnionUF grid;
    private final WeightedQuickUnionUF grid2;
    private int openCount;
    private boolean percolated;
    private int lastSite;
    private final int start;
    private final int end;

    // create n-by-n state, with all sites blocked
    public Percolation(int n) {
        if (n <= 0) {
            throw new java.lang.IllegalArgumentException();
        }

        this.n = n;
        this.size = n * n;
        this.start = 0; // abstract root
        this.end = this.size + 1; // abstract end
        this.openCount = 0;
        this.open = new boolean[this.end + 1];
        this.grid = new WeightedQuickUnionUF(this.end + 1);
        this.grid2 = new WeightedQuickUnionUF(this.end + 1);

        this.open[start] = true;
        this.open[end] = false;
        this.lastSite = -1;

        // connect first row sites with 'start'
        for (int i = 1; i <= n; i++) {
            grid.union(start, i);
        }

        // connect last row sites with 'end'
        for (int i = n * (n - 1) + 1; i < this.end; i++) {
            grid2.union(i, end);
        }
    }

    private int xyTo1D(int row, int col) {
        // (1, 1) = 1, (1, 2) = 2, ...
        return (row - 1) * n + (col - 1) + 1;
    }

    // open site (row, col) if it is not open
    public void open(int row, int col) {
        validate(row, col);

        int id = xyTo1D(row, col);

        if (open[id]) {
            // already open
            return;
        }

        int selfId = grid.find(id); // reduces subsequent finds
        int selfId2 = grid2.find(id);

        if (row > 1) {
            int top = xyTo1D(row - 1, col);

            if (open[top]) {
                grid.union(selfId, top);
                grid2.union(selfId2, top);
            }
        }

        if (col > 1) { // left
            int left = xyTo1D(row, col - 1);

            if (open[left]) {
                grid.union(selfId, left);
                grid2.union(selfId2, left);
            }
        }

        if (row < n) {
            int bottom = xyTo1D(row + 1, col);

            if (open[bottom]) {
                grid.union(selfId, bottom);
                grid2.union(selfId2, bottom);
            }
        }

        if (col < n) {
            int right = xyTo1D(row, col + 1);

            if (open[right]) {
                grid.union(selfId, right);
                grid2.union(selfId2, right);
            }
        }

        this.open[id] = true; // open the site
        this.lastSite = id;
        this.openCount++;

        // update 'percolated'
        if (!percolated && isFull(lastSite) && grid2.connected(lastSite, end)) {
            percolated = true;
        }
    }

    // is site (row, col) openRandom?
    public boolean isOpen(int row, int col) {
        validate(row, col);
        return open[xyTo1D(row, col)];
    }

    // is site (row, col) full?
    public boolean isFull(int row, int col) {
        validate(row, col);
        return isFull(xyTo1D(row, col));
    }

    private boolean isFull(int id) {
        // TODO: check backwash (connected through 'end')
        return open[id] && grid.connected(id, start);
    }

    // number of openRandom sites
    public int numberOfOpenSites() {
        return this.openCount;
    }

    // does the system percolate?
    public boolean percolates() {
        return percolated;
    }

    private void validate(int row, int col) {
        if (row < 1 || row > n || col < 1 || col > n) {
            throw new java.lang.IllegalArgumentException();
        }
    }

    public static void main(String[] args) {
        Percolation grid = new Percolation(3);
        grid.open(1, 1);
        grid.open(2, 1);
        grid.open(3, 1);

        //grid.open(3, 3); // backwash

        System.out.println(grid.isFull(3, 3));
        System.out.println(grid.percolates());
    }

}
