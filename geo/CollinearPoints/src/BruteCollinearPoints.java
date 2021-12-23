
import java.util.Arrays;

public class BruteCollinearPoints {

    private final LineSegment[] seg;
    private int n = 0; // number of collinear segments

    // finds all line segments containing 4 points
    public BruteCollinearPoints(Point[] points) {
        validate(points);

        Point[] ps = Arrays.copyOf(points, points.length);
        Segment[] colls = new Segment[ps.length]; // temporary array

        for (int i = 0; i < ps.length; i++) { // point 1
            for (int j = i + 1; j < ps.length; j++) { // point 2

                double sl1 = ps[i].slopeTo(ps[j]);
                if (Double.NEGATIVE_INFINITY == sl1) {
                    throw new java.lang.IllegalArgumentException();
                }

                for (int x = j + 1; x < ps.length; x++) { // point 3

                    double sl2 = ps[i].slopeTo(ps[x]);

                    if (Double.compare(sl1, sl2) != 0) {
                        continue;
                    }

                    for (int y = x + 1; y < ps.length; y++) { // point 4

                        double sl3 = ps[i].slopeTo(ps[y]);

                        if (Double.compare(sl2, sl3) == 0) {
                            colls[n] = new Segment(ps[i], ps[j]);
                            colls[n].merge(ps[x]);
                            colls[n].merge(ps[y]);
                            n++;

                            // extend temporary array
                            if (n == colls.length) {
                                colls = Arrays.copyOf(colls, colls.length * 2);
                            }
                        }
                    }
                }
            }

        }

        seg = cleanup(Arrays.copyOf(colls, n));
    }

    // the number of line segments
    public int numberOfSegments() {
        return n;
    }

    // the line segments
    public LineSegment[] segments() {
        return Arrays.copyOf(seg, n);
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


    private LineSegment[] cleanup(Segment[] segs) {
        Arrays.sort(segs, (Segment s1, Segment s2) -> {
            return s1.compareTo(s2);
        });

        // remove duplicates & copy
        int m = 0;
        LineSegment[] ret = new LineSegment[segs.length];

        int i = 0;
        while (i < segs.length) {
            ret[m++] = new LineSegment(segs[i].start, segs[i].end);

            // skip same segments
            do {
                i++;
            } while (i < segs.length && segs[i].compareTo(segs[i - 1]) == 0);
        }
        this.n = m;
        return ret;
    }

    // helper class
    private class Segment {

        Point start, end;

        Segment(Point p, Point q) {
            if (p.compareTo(q) < 0) {
                start = p;
                end = q;
            } else {
                start = q;
                end = p;
            }

        }

        void merge(Point p) {
            if (p.compareTo(start) < 0) {
                start = p;
            }
            if (end.compareTo(p) < 0) {
                end = p;
            }
        }

        int compareTo(Segment that) {
            if (start == that.start) {
                return end.compareTo(that.end);
            }
            return start.compareTo(that.start);
        }
    }

}
