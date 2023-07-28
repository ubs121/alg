/*
	 NxN хөлөг дээр 2 тэргийг хамгийн их нийлбэртэй байхаар байрлуулах
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// эхний мөр == хөлгийн хэмжээ
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	board := make([][]int, n+1) // хөлөг
	board[n] = make([]int, n+1)

	// өгөгдөл унших
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")

		board[i] = make([]int, n+1)

		for j := 0; j < len(arr); j++ {
			x, _ := strconv.Atoi(arr[j])
			board[i][j] = x

			// оролт хийх үед нийлбэрийг бодов
			board[i][n] += x
			board[n][j] += x
		}
	}

	// TODO: Greedy: мөр, багануудыг хамгийн ихээс нь бага руу нь эрэмбэлж нэг булан руу нь шахах байдлаар массивыг "дахин зохион" байгуулж болно. энэ нь үр дүнд нөлөөлөхгүй
	// TODO: үүний дараа уг булангаас диагоналийн дагуу доош давтан шалгаж болно

	max := 0
	total := 0

	for x1:=0; x1<n; x1++ {
		for x2:=x1+1; x2<n; x2++ {

			for y1:=0; y1<n; y1++ {

				for y2:=y1+1; y2<n; y2++ {

					total = board[x1][n] + board[x2][n] + board[n][y1]  + board[n][y2] - 2 * ( board[x1][y1]  + board[x2][y2] ) - board[x1][y2] - board[x2][y1]

					total = board[x1][n] + board[n][y1] - 2 * board[x1][y1]
					total = board[x2][n] + board[n][y2] - 2 * board[x2][y2]
				  total -= (board[x1][y2] + board[x2][y1])

					//fmt.Printf("(%d,%d) (%d,%d) - %d\n", x1, y1, x2, y2, total)

					// max утгатай байрлал олох (x1, y1), (x2, y2)
					if total > max {
						max = total
					}

				}
			}
		}
	}

	fmt.Println(max)

}
