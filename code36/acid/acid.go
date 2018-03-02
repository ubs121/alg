// https://www.hackerrank.com/contests/w36/challenges/acid-naming
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var q int
	fmt.Scanf("%d\n", &q)

	reader := bufio.NewReader(os.Stdin)
	var line string

	for i := 0; i < q; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimRight(line, "\n\r")

		if strings.HasSuffix(line, "ic") {
			if strings.HasPrefix(line, "hydro") {
				fmt.Println("non-metal acid")
			} else {
				fmt.Println("polyatomic acid")
			}
		} else {
			fmt.Println("not an acid")
		}
	}
}
