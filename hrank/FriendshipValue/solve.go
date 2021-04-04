package main

import (
	"fmt"
)

type student struct {
	id  int   // student id
	adj []int // friends

	// node data
	color int // White=0, Gray=1, Black=2
}

type set []student

// do BFS from 'root'
func totalBFS(data set, root int) int64 {
	n := int64(1) // node count
	r := int64(0) // repeated connections

	var Q []int
	var u, v int

	Q = append(Q, root)
	for len(Q) > 0 {
		Q, u = Q[1:], Q[0] // dequeue

		for i := 0; i < len(data[u].adj); i++ {
			v = data[u].adj[i]

			if data[v].color == 0 {
				data[v].color = 1

				Q = append(Q, v) // enqueue

				n++ // new node
			} else if data[v].color == 1 {
				r++ // count repeated connection
			}
		}

		data[u].color = 2
	}

	//fmt.Printf("%d, %d\n", n, r)

	// calculate the total
	total := ((n - 1) * n * (n + 1)) / 3
	total += n * (n - 1) * r

	return total
}

func main() {
	var q int
	fmt.Scanf("%d\n", &q)

	var n, m int
	var x, y int

	// iterate over test cases
	for T := 0; T < q; T++ {
		fmt.Scanf("%d %d\n", &n, &m)

		data := make(set, n+1)

		for i := 0; i < m; i++ {
			fmt.Scanf("%d %d\n", &x, &y)

			data[x].adj = append(data[x].adj, y)
			data[y].adj = append(data[y].adj, x)
		}

		var total, max int64
		for root := 1; root <= n; root++ {
			if data[root].color == 0 {
				total = totalBFS(data, root)

				if total > max {
					max = total
				}
			}
		}

		fmt.Println(max)

	}

}
