// https://www.hackerrank.com/challenges/coin-change

package main

import (
	"fmt"
	"sort"
)

var (
	// solution store ( amount -> ( first_k_coins -> number_of_variants) )
	mem map[int]map[int]int
)

func solve(n int, c []int) int {

	// check from the solution store
	if arrSol, ok := mem[n]; ok && arrSol != nil {
		if k, ok := arrSol[len(c)]; ok {
			return k
		}
	}

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	k := 0

	// бага дугаартай зоосноос эхлэн задлах
	for i := 0; i < len(c); i++ {
		//fmt.Printf("(%d-%d)", n, c[i])

		k = k + solve(n-c[i], c[i:]) // remove duplicates c[i:]
	}

	// remember in the solution store
	if _, exists := mem[n]; !exists {
		mem[n] = make(map[int]int)
	}

	mem[n][len(c)] = k

	return k
}

func main() {
	var n, m int

	fmt.Scanf("%d %d", &n, &m)

	c := make([]int, m) // зоосууд
	mem = make(map[int]map[int]int)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &c[i])
	}

	sort.Ints(c)

	r := solve(n, c)

	fmt.Println(r)
}
