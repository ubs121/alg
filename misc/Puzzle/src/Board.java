
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;
import java.util.ArrayList;
import java.util.List;

/**
 *
 * @author ub
 */
public class Board {

    private final byte n;
    private final int hamm;
    private final int manh;
    private final int hash;

    private final short[][] tiles;
    private Iterable<Board> nbs;
    private byte x0, y0; // blank square position
    private Board twin;

    // construct a board from an n-by-n array of blocks
    // (where blocks[i][j] = block in row i, column j)
    public Board(int[][] blocks) {
        this.n = (byte) blocks.length;
        this.tiles = copy(blocks);

        int m = 0;
        int h = 0;
        int result = 1;

        int val;
        int x, y; // goal position
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                val = blocks[i][j];
                if (val != 0) {
                    x = (val - 1) / n;
                    y = (val - 1) % n;

                    // manhattan: distance from goal
                    m += Math.abs(i - x) + Math.abs(j - y);

                    // hamming: count wrong positions
                    if (x != i || y != j) {
                        h++;
                    }
                } else {
                    x0 = (byte) i;
                    y0 = (byte) j;
                }

                // calculate hash code
                result = 31 * result + val;

            }

        }

        manh = m;
        hamm = h;
        hash = result;
    }

    private Board(Board parent, int dx, int dy) {
        this.n = (byte) parent.tiles.length;
        this.tiles = copy(parent.tiles);

        tiles[parent.x0][parent.y0] = tiles[parent.x0 + dx][parent.y0 + dy];
        tiles[parent.x0 + dx][parent.y0 + dy] = 0;

        int m = 0;
        int h = 0;
        int result = 1;

        int val;
        int x, y; // goal position
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                val = tiles[i][j];
                if (val != 0) {
                    x = (val - 1) / n;
                    y = (val - 1) % n;

                    // manhattan: distance from goal
                    m += Math.abs(i - x) + Math.abs(j - y);

                    // hamming: count wrong positions
                    if (x != i || y != j) {
                        h++;
                    }
                } else {
                    x0 = (byte) i;
                    y0 = (byte) j;
                }

                // calculate hash code
                result = 31 * result + val;

            }

        }

        manh = m;
        hamm = h;
        hash = result;
    }

    // board dimension n
    public int dimension() {
        return n;
    }

    // number of blocks out of place
    public int hamming() {
        return hamm;
    }

    // sum of Manhattan distances between blocks and goal
    public int manhattan() {
        return manh;
    }

    // is this board the goal board?
    public boolean isGoal() {
        return (this.hamm == 0);
    }

    // a board that is obtained by exchanging any pair of blocks
    public Board twin() {
        if (twin == null) {

            int x1, y1, x2, y2;
            do {
                // choose 2 cells randomly
                x1 = StdRandom.uniform(n);
                y1 = StdRandom.uniform(n);
                x2 = StdRandom.uniform(n);
                y2 = StdRandom.uniform(n);

                if (tiles[x1][y1] != 0 && tiles[x2][y2] != 0 && (x1 != x2 || y1 != y2)) {
                    break;
                }
            } while (true);

            int[][] cp = new int[n][n];
            for (int i = 0; i < n; i++) {
                cp[i] = new int[n];
                for (int j = 0; j < n; j++) {
                    cp[i][j] = tiles[i][j];
                }
            }

            // exchange
            int tmp = cp[x1][y1];
            cp[x1][y1] = cp[x2][y2];
            cp[x2][y2] = tmp;

            twin = new Board(cp);
        }
        return twin;
    }

    // all neighboring boards
    public Iterable<Board> neighbors() {
        if (nbs == null) {
            List<Board> nbsList = new ArrayList<>(2);

            if (0 < x0) { // left
                nbsList.add(new Board(this, -1, 0));
            }

            if (0 < y0) { // up
                nbsList.add(new Board(this, 0, -1));
            }

            if (x0 < n - 1) { // right
                nbsList.add(new Board(this, 1, 0));
            }

            if (y0 < n - 1) { // down
                nbsList.add(new Board(this, 0, 1));
            }

            nbs = () -> nbsList.iterator();
        }

        return nbs;
    }

    private short[][] copy(int[][] blocks) {
        short[][] cp = new short[n][];
        for (int i = 0; i < n; i++) {
            cp[i] = new short[n];
            for (int j = 0; j < n; j++) {
                cp[i][j] = (short) blocks[i][j];
            }
        }
        return cp;
    }

    private short[][] copy(short[][] blocks) {
        short[][] cp = new short[n][];
        for (int i = 0; i < n; i++) {
            cp[i] = new short[n];
            System.arraycopy(blocks[i], 0, cp[i], 0, n);
        }
        return cp;
    }

    // does this board equal y?
    @Override
    public boolean equals(Object y) {
        if (this == y) {
            return true;
        }

        if (y instanceof Board) {
            Board that = (Board) y;
            return hash == that.hash;
        }

        return false;
    }

    // string representation of this board (in the output format specified below)
    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append(n).append("\n");
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                sb.append(" ").append(tiles[i][j]);
            }
            sb.append("\n");
        }

        return sb.toString();
    }

    // unit tests (not graded)
    public static void main(String[] args) {
        int[][] blk = new int[][]{{0, 1, 3}, {4, 2, 5}, {7, 8, 6}};

        Board b = new Board(blk);
        for (Board n : b.neighbors()) {
            StdOut.println(n);
        }

        int[][] goal = new int[][]{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}};
        StdOut.printf("Goal ? %b\n", new Board(goal).isGoal());

        Board b1 = new Board(blk);
        Board b2 = new Board(blk);
        StdOut.printf("Equal ? %b\n", b2.equals(b1));

        int[][] blk3 = new int[][]{{0, 1, 3}, {4, 2, 5}, {7, 8, 6}};
        Board b3 = new Board(blk3);
        StdOut.printf("Equal hash ? %b\n", b1.equals(b3));
        StdOut.printf("Equal manhattan ? %b\n", b1.manhattan() == b3.manhattan());
        StdOut.printf("Equal hamming ? %b\n", b1.hamming() == b3.hamming());
    }
}
