package main

import "fmt"

type node struct {
	parent int
	adj    []int
}

type tree []node

var data tree

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	data = make(tree, n+1)

	var u, v int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &u, &v)

		data[u].adj = append(data[u].adj, v)
		data[v].adj = append(data[v].adj, u)
	}

	// build tree (fix reference)
	dfs(1)

	cuts := 0

	// a function that counts children
	var countChildren func(root int) int

	countChildren = func(root int) int {
		u := data[root]

		c := 0 // number of children of 'u', including itself
		for i := 0; i < len(u.adj); i++ {
			c += countChildren(u.adj[i])
		}

		if (c+1)%2 == 0 {
			// count it
			cuts++
		}

		return c + 1
	}

	c := countChildren(1)

	if c%2 == 0 {
		cuts-- // rollback, because root=1 is counted in cuts
	}

	// print answer
	fmt.Println(cuts)

}

func dfs(root int) {
	var u, v int
	visited := make([]byte, len(data))

	var stack []int
	stack = append(stack, root)
	for len(stack) > 0 {
		stack, u = stack[:len(stack)-1], stack[len(stack)-1] // pop

		i := 0
		for i < len(data[u].adj) {
			v = data[u].adj[i]

			if visited[v] == 0 {
				visited[v] = 1
				data[v].parent = u

				stack = append(stack, v) // push
				i++
			} else {
				// remove 'v' from u.children
				data[u].adj = data[u].adj[i+1:]
			}
		}

		visited[u] = 2
	}
}
