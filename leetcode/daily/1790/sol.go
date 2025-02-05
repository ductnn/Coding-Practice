package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }

    var diff []int
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diff = append(diff, i)
        }
    }

    if len(diff) != 2 {
        return false
    }

    i, j := diff[0], diff[1]
    return s1[i] == s2[j] && s1[j] == s2[i]
}

func main() {
    s1 := "bank" 
    s2 := "kanb"
   
    fmt.Print(areAlmostEqual(s1, s2))
}
