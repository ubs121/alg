package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	isLucky := func(price string) (bool, int) {
		n4 := 0
		n7 := 0
		p := 0
		for i := 0; i < len(price); i++ {
			if price[i] == '4' {
				p = p*10 + 4
				n4++
			} else if price[i] == '7' {
				p = p*10 + 7
				n7++
			} else {
				return false, 0
			}
		}
		return (n4 == n7), p

	}

	reader := bufio.NewReader(os.Stdin)

	minName := "-1"
	minPrice := 777777777
	var sp []string
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		if line[len(line)-1] == '\n' {
			line = line[0 : len(line)-1] // remove '\n'
		}

		sp = strings.Split(line, " ")
		ok, p := isLucky(sp[1])

		if ok && p < minPrice {
			minName = sp[0]
			minPrice = p
		}
	}

	fmt.Printf("%s\n", minName)

}
