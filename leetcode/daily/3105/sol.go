package main

import "fmt"

func longestMonotonicSubarray(nums []int) int {
    res := 1
    tmp := 1

    for i, v := range nums[1:] {
        if nums[i] < v {
            tmp++
            res = max(res, tmp)
        } else {
            tmp = 1
        }
    }
    
    tmp = 1
    for i, v := range nums[1:] {
        if nums[i] > v {
            tmp++
            res = max(res, tmp)
        } else {
            tmp = 1
        }
    }

    return res
}

func main() {
    nums := []int{1,4,3,3,2}

    fmt.Println(longestMonotonicSubarray(nums))
}
