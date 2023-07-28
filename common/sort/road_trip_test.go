package sort

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

//  in a straight line,  to travel this distance on a single tank of gasoline
// but you only know distances to each city from your starting point.

func findDistances(inp string) []int {
	// parse the input
	dists := parseInput(inp)

	if len(dists) == 0 {
		return nil
	}

	// sort distances
	sort.Ints(dists)

	// find distances between nearest two cities
	ret := make([]int, len(dists))
	ret[0] = dists[0] // first city distance
	for i := 1; i < len(dists); i++ {
		ret[i] = dists[i] - dists[i-1]
	}
	return ret
}

func parseInput(inp string) []int {
	pairs := strings.Split(inp, ";")
	var dists []int
	for _, p := range pairs {
		pair := strings.Split(p, ",")
		if len(pair) < 2 {
			continue // empty entry, caused by strings.Split()
		}
		v, err := strconv.ParseInt(pair[1], 10, 64)
		if err != nil {
			panic(err)
		}
		dists = append(dists, int(v))
	}
	return dists
}

func joinInts(arr []int) string {
	var tmpArr []string
	for _, v := range arr {
		tmpArr = append(tmpArr, strconv.Itoa(v))
	}
	return strings.Join(tmpArr, ",")
}

func TestRoadTrip(t *testing.T) {
	testCases := map[string]string{
		"Rkbs,5453; Wdqiz,1245; Rwds,3890; Ujma,5589; Tbzmo,1303;":     "1245,58,2587,1563,136",
		"Vgdfz,70; Mgknxpi,3958; Nsptghk,2626; Wuzp,2559; Jcdwi,3761;": "70,2489,67,1135,197",
	}
	for tc, exp := range testCases {
		got := findDistances(tc)
		gotStr := joinInts(got)
		if gotStr != exp {
			t.Errorf("%s: exp %s, got %s", tc, exp, gotStr)
		}
	}
}
