package main

import (
	"fmt"
	"leetcode/problem"
)

func main()  {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(problem.TwoSumTwoMap(nums, target))
}
