package main

import "fmt"

var pow2 map[int]int

// 2-н зэрэгтүүд хадгалах
func init_pow() {
	pow2 = make(map[int]int)

	p := 1
	for i := 0; i < 64; i++ {
		pow2[i] = p
		p = p << 1
	}
}

func xorCount(x int) int {
	i := x
	d := 0
	for i > 0 {
		d += 1
		i = i / 2
	}
	return pow2[d] - x - 1
}

func xorCount2(x int) int {
	n := 0
	for a := 1; a < x; a++ {
		if a^x > x {
			n++
		}
	}

	return n
}

func main() {
	init_pow()

	//var q int
	//fmt.Scanf("%d", &q)

	x := 100
	s1 := xorCount(x)
	s2 := xorCount2(x)
	fmt.Println("s1=", s1, "s2=", s2)

}
