package main

import (
	"alg/sort"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	/* элементийн тоо */
	n := 0

	/* Сул зайгаар тусгаарлагдсан тоон дарааллыг унших */
	for i := 0; i < len(s); n++ {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		fmt.Sscanf(s[i:], "%d", &array[n])
		for i < len(s) && s[i] != ' ' {
			i++
		}
	}

	/* массивыг эрэмбэлэх */
	sort.QuickSort(array[:], 0, n-1)

	/* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}
}
