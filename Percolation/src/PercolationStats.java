
import edu.princeton.cs.algs4.StdRandom;
import edu.princeton.cs.algs4.StdStats;

public class PercolationStats {

    private static final double CONFIDENCE_95 = 1.96;
    
    private final int n;
    private final int trials;
    private double[] sample;

    
            
    public PercolationStats(int n, int trials) {   // perform trials independent experiments on an n-by-n grid
        if (n <= 0 || trials <= 0) {
            throw new java.lang.IllegalArgumentException();
        }

        this.n = n;
        this.trials = trials;
        this.sample = new double[trials];
    }

    // sample mean of percolation threshold
    public double mean() {
        return StdStats.mean(sample);
    }

    // sample standard deviation of percolation threshold
    public double stddev() {
        return StdStats.stddev(sample);
    }

    // low  endpoint of 95% confidence interval
    public double confidenceLo() {
        return this.mean() - CONFIDENCE_95 * this.stddev() / Math.sqrt(trials);
    }

    // high endpoint of 95% confidence interval
    public double confidenceHi() {
        return this.mean() + CONFIDENCE_95 * this.stddev() / Math.sqrt(trials);
    }

    private void run() {

        for (int t = 0; t < trials; t++) {
            Percolation grid = new Percolation(n);

            int randSite;

            // repeat until it percolates
            while (!grid.percolates()) {
                // choose a site uniformly at random among all blocked sites
                while (true) {
                    randSite = StdRandom.uniform(n * n) + 1;
                    
                    // conversion
                    int row = (randSite - 1) / n + 1;
                    int col = (randSite - 1) % n + 1;
                    
                    if (!grid.isOpen(row, col)) {
                        // open the site
                        grid.open(row, col);

                        break;
                    }
                }

            }

            sample[t] = (double) grid.numberOfOpenSites() / (n * n);

        }

        System.out.printf("mean                     = %f\n", this.mean());
        System.out.printf("stddev                   = %f\n", this.stddev());
        System.out.printf("95%% confidence interval = [%f,%f]\n", this.confidenceLo(), this.confidenceHi());

    }

    public static void main(String[] args) {    // test client (described below)
        if (args.length < 2) {
            throw new java.lang.IllegalArgumentException();
        }

        int n, trials;
        n = Integer.parseInt(args[0]);
        trials = Integer.parseInt(args[1]);

        // run experiment
        PercolationStats percolationStats = new PercolationStats(n, trials);
        percolationStats.run();
    }
}
