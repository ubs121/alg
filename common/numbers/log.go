package numbers

import "math"

func Log2Int(n int) int {
	return int(math.Log(float64(n)) / math.Log(2))
}
