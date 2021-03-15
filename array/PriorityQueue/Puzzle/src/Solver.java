
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.MinPQ;
import edu.princeton.cs.algs4.StdOut;
import java.util.Arrays;

/**
 *
 * @author ub
 */
public class Solver {

    private static final int UNKNOWN = -1;
    private static final int SOLVABLE = 1;
    private static final int UNSOLVABLE = 0;

    private int status; // -1 - not know, 0 - not solvable, 1 - solved
    private Board[] solution;

    // find a solution to the initial board (using the A* algorithm)
    public Solver(Board initial) {
        if (initial == null) {
            throw new java.lang.IllegalArgumentException();
        }
        search(initial);
    }

    // A* algorithm
    private void search(Board initial) {
        MinPQ<SearchNode> pq;
        MinPQ<SearchNode> pqTwin;

        status = UNKNOWN;
        pq = new MinPQ<>();
        pq.insert(new SearchNode(initial, 0, null));
        SearchNode current = null;

        Board twin = initial.twin();
        pqTwin = new MinPQ<>();
        pqTwin.insert(new SearchNode(twin, 0, null));
        SearchNode currentTwin = null;

        while (!pq.isEmpty() && status != UNSOLVABLE) {
            current = pq.delMin();

            if (current.board.isGoal()) {
                status = SOLVABLE;
                constructSolution(current);
                break;
            }

            // enqueue neighbors: OutOfMemoryError!!!
            for (Board child : current.board.neighbors()) {
                if (current.parent != null && child.equals(current.parent.board)) {
                    continue;
                }

                SearchNode newNode = new SearchNode(child, current.moves + 1, current);
                pq.insert(newNode);

            }

            // Detecting unsolvable puzzles, if current == currentTwin then it's solvable
            if (status == UNKNOWN && !pqTwin.isEmpty()) {
                currentTwin = pqTwin.delMin();

                if (currentTwin.board.isGoal()) {
                    status = UNSOLVABLE;
                }

                // twin's neighbors 
                for (Board nb : currentTwin.board.neighbors()) {
                    if (currentTwin.parent != null && nb.equals(currentTwin.parent.board)) {
                        continue;
                    }

                    pqTwin.insert(new SearchNode(nb, currentTwin.moves + 1, currentTwin));
                }
            }
        }

    }

    // is the initial board solvable?
    public boolean isSolvable() {
        return status == SOLVABLE;
    }

    // min number of moves to solve initial board; -1 if unsolvable
    public int moves() {
        if (isSolvable() && solution != null) {
            return solution.length - 1;
        }
        return -1;
    }

    // sequence of boards in a shortest solution; null if unsolvable
    public Iterable<Board> solution() {
        if (isSolvable() && solution != null) {
            return () -> Arrays.asList(solution).iterator();
        }
        return null;
    }

    private void constructSolution(SearchNode last) {
        solution = new Board[last.moves + 1];
        int k = last.moves;
        SearchNode p = last;
        do {
            solution[k--] = p.board;
            p = p.parent;
        } while (p != null);

    }

    // SearchNode, helper class
    private class SearchNode implements Comparable<SearchNode> {

        private final Board board;
        private final int moves;     // distance from the initial node
        private final SearchNode parent; // parent node

        public SearchNode(Board b, int move, SearchNode parent) {
            this.board = b;
            this.parent = parent;
            this.moves = move;
        }

        public int compareTo(SearchNode that) {
            return (board.manhattan() - that.board.manhattan()) + (moves - that.moves);
        }

        public boolean hasParent(Board b) {
            SearchNode p = parent;

            while (p != null) {
                if (p.board.equals(b)) {
                    return true;
                }
                p = p.parent;
            }
            return false;
        }
    }

    // solve a slider puzzle (given below)   
    public static void main(String[] args) {
        // create initial board from file
        In in = new In(args[0]);
        int n = in.readInt();
        int[][] blocks = new int[n][n];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                blocks[i][j] = in.readInt();
            }
        }
        Board initial = new Board(blocks);

        // solve the puzzle
        Solver solver = new Solver(initial);

        // print solution to standard output
        if (!solver.isSolvable()) {
            StdOut.println("No solution possible");
        } else {
            StdOut.println("Minimum number of moves = " + solver.moves());
            for (Board board : solver.solution()) {
                StdOut.println(board);
            }
        }
    }
}
