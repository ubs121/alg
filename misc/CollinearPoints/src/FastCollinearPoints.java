
import java.util.Arrays;

public class FastCollinearPoints {

    private final LineSegment[] segments;
    private final int n; // number of collinear segments

    // finds all line segments containing 4 or more points
    public FastCollinearPoints(Point[] points) {
        validate(points);

        Point[] ps = Arrays.copyOf(points, points.length);
        Point[][] colls = new Point[points.length / 4][2];

        int m = 0;

        for (Point p : points) {

            // sort points by slope to 'p'
            Arrays.sort(ps, p.slopeOrder());

            // duplicate point
            if (ps.length > 1 && p.compareTo(ps[1]) == 0) {
                throw new java.lang.IllegalArgumentException();
            }

            // count equal segments
            int q = 1; // skip first point, because points[0] == p
            while (q < ps.length - 1) {

                Point lo = p, hi = ps[q];
                if (p.compareTo(ps[q]) > 0) {
                    lo = ps[q];
                    hi = p;
                }

                double pq = lo.slopeTo(hi);

                int r = q + 1;
                while (r < ps.length && Double.compare(pq, p.slopeTo(ps[r])) == 0) {
                    if (lo.compareTo(ps[r]) > 0) {
                        lo = ps[r];
                    }
                    if (hi.compareTo(ps[r]) < 0) {
                        hi = ps[r];
                    }
                    r++;
                }

                if (r - q > 2) { // new collinear segment found
                    colls[m][0] = lo;
                    colls[m][1] = hi;
                    m++;

                    // extend array
                    if (m == colls.length) {
                        Point[][] tmp = new Point[colls.length * 2][2];
                        System.arraycopy(colls, 0, tmp, 0, m);
                        colls = tmp;
                    }
                    q = r;
                } else {
                    q++;
                }
            }
        }

        // cut by m
        colls = Arrays.copyOf(colls, m);

        // sort
        Arrays.sort(colls, (Point[] s1, Point[] s2) -> {
            if (s1[0].compareTo(s2[0]) == 0) {
                return s1[1].compareTo(s2[1]);
            }
            return s1[0].compareTo(s2[0]);
        });

        // remove duplicates & copy into 'segments'
        segments = new LineSegment[m];
        int i = 0, k = 0;
        while (i < colls.length) {
            segments[k++] = new LineSegment(colls[i][0], colls[i][1]);

            do {
                i++;
            } while (i < colls.length && compareSegments(colls[i], colls[i - 1]) == 0);
        }
        this.n = k;
    }

    private void validate(Point[] ps) {
        if (ps == null) {
            throw new java.lang.IllegalArgumentException();
        }

        for (Point p : ps) {
            if (p == null) {
                throw new java.lang.IllegalArgumentException();
            }
        }
    }

    private int compareSegments(Point[] s1, Point[] s2) {
        if (s1[0].compareTo(s2[0]) == 0) {
            return s1[1].compareTo(s2[1]);
        }
        return s1[0].compareTo(s2[0]);
    }

    // the number of line segments
    public int numberOfSegments() {
        return n;
    }

    // the line segments
    public LineSegment[] segments() {
        return Arrays.copyOf(segments, n);
    }

}
