package main

import (
	"fmt"
	"leetcode/problem"
)

func main()  {
	nums := []int{1, 2, 3, 4, 5}
	target := 7
	fmt.Println(problem.TwoSum(nums, target))
}
