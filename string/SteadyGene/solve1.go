// https://www.hackerrank.com/chaienges/bear-and-steady-gene

package main

import "fmt"
import "os"
import "bufio"

// 40
// TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC
// 5

type Quad [4]int

func main() {

	file, _ := os.Open("input02.txt")
	defer file.Close()

	buf := make([]byte, 500100)
	scanner := bufio.NewScanner(os.Stdin) // file, os.Stdin
	scanner.Buffer(buf, 500100)

	var n int // 4<=n<=500'000
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	scanner.Scan()
	s := scanner.Text()

	limit := n / 4
	left := make([]Quad, n+1)

	ind := map[byte]int{'G': 0, 'A': 1, 'C': 2, 'T': 3}

	// find l, where l = [0, max(i)]
	l := 0
	for i := 0; i < n; i++ {
		left[i][ind[s[i]]]++
		left[i+1] = left[i]

		if l == 0 && left[i][ind[s[i]]] > limit {
			l = i
		}
	}

	total := left[n]
	over := 0 // minimum length to be replaced
	for p := 0; p < 4; p++ {
		if total[p] > limit {
			over += total[p] - limit
		}
	}

	//fmt.Println(l, total, over)

	min := n

	// 0 <- i
	for i := l - 1; i >= 0; i-- {
		lo := i + over - 1
		hi := n

		if hi > i+min {
			hi = i + min
		}

		// find j between [lo, hi] using binary search
		j := lo
		for lo < hi-1 {
			j = (lo + hi) / 2
			//fmt.Println(lo, hi)

			// check if i+j <= limit
			if left[i][0]+total[0]-left[j][0] <= limit && left[i][1]+total[1]-left[j][1] <= limit &&
				left[i][2]+total[2]-left[j][2] <= limit && left[i][3]+total[3]-left[j][3] <= limit {
				hi = j // j=(lo+j)/2

				if min > j-i {
					min = j - i
				}
			} else {
				lo = j // j=(j+hi)/2
			}
		}
	}

	if over > 0 {
		fmt.Println(min)
	} else {
		fmt.Println(0)
	}

}
