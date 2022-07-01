// https://coursera.cs.princeton.edu/algs4/assignments/collinear/specification.php
// The problem. Given a set of n distinct points in the plane, find every (maximal) line segment that connects a subset of 4 or more of the points.
package geometric

import (
	"fmt"
	"math"
	"testing"
)

func TestCollinear(t *testing.T) {
	type testCase struct {
		points   []*point
		expected int
	}
	testCases := []testCase{
		{[]*point{{19000, 10000}, {18000, 10000}, {32000, 10000}, {21000, 10000}, {1234, 5678}, {14000, 10000}}, 1},
		{[]*point{{10000, 0}, {0, 10000}, {3000, 7000}, {7000, 3000}, {20000, 21000}, {3000, 4000}, {14000, 15000}, {6000, 7000}}, 2},
	}

	for _, tc := range testCases {
		got := CollinearPoints(tc.points)
		if tc.expected != len(got) {
			t.Errorf("testCase %v: exp %d, got %d", tc.points, tc.expected, len(got))
			fmt.Printf("%v", got)
		}
	}
}

func CollinearPoints(points []*point) [][]*point {
	var result [][]*point
	lines := lineSet{}
	for _, p := range points {
		pSlopes := make(map[float64][]*point) // p.slope => [point1,point2,...]
		for _, q := range points {
			if p == q {
				continue
			}
			pqSlope := p.slopeTo(q)                        // determine the slope it makes with p
			pSlopes[pqSlope] = append(pSlopes[pqSlope], q) // add into a segment
		}

		for sl, ps := range pSlopes {
			if len(ps) > 1 && !lines.contains(p, sl) { // found at least 3 points that make 2 pairs
				seg := append(ps, p)
				result = append(result, seg)
				for _, q := range seg {
					lines.add(q, sl) // mark all lines
				}
			}
		}
	}
	return result
}

type point struct {
	x, y float64
}

func (p *point) slopeTo(that *point) float64 {
	if that.x == p.x {
		if that.y == p.y {
			return math.Inf(-1)
		}
		return math.Inf(1)
	}
	if that.y == p.y {
		return +0.0
	}
	return (that.y - p.y) / (that.x - p.x)
}

func (p point) String() string {
	return fmt.Sprintf("(%.2f,%.2f)", p.x, p.y)
}

type lineSet struct {
	lines map[*point]map[float64]bool
}

// add a line through 'p'
func (l *lineSet) add(p *point, slope float64) {
	if l.lines == nil {
		l.lines = make(map[*point]map[float64]bool)
	}
	m, exists := l.lines[p]
	if !exists {
		m = map[float64]bool{}
	}
	m[slope] = true
	l.lines[p] = m
}

// checks if the set contains a line
func (l *lineSet) contains(p *point, slope float64) bool {
	if m, exists := l.lines[p]; exists {
		if _, slopeExists := m[slope]; slopeExists {
			return true
		}
	}
	return false
}
