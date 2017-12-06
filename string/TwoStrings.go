package main

import "fmt"

func checkIntersect(s1, s2 string) bool {
    m1 := map[rune]bool{}

    for i := 0; i < len(s1); i++ {
        m1[s1[i]] = true
    }

    for i := 0; i < len(s2); i++ {
        if _, ok := m1[s2[i]]; ok {
            return true
        }
    }

    return false
}

func main() {
    var n int
    fmt.Scanf("%d", &n)

    var s1, s2 string
    for i := 0; i < n; i++ {
        fmt.Scanln(&s1)
        fmt.Scanln(&s2)

        if checkIntersect(s1, s2) {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
