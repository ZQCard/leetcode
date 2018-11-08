package problem

// 56ms
func TwoSum(nums []int, target int) []int{
	var res  []int
	for key, value := range nums {
		for k, v := range nums{
			if k <= key {
				continue
			}
			if value + v == target {
				res = append(res, key)
				res = append(res, k)
				return res
			}
		}
	}
	return res
}

// 44ms
func TwoSum2(nums []int, target int) []int  {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 8ms
func TwoSumTwoMap(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[i] = nums[i]
	}

	numMap2 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap2[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - numMap[i]
		if k, ok := numMap2[complement]; k != i && ok{
			return []int{i, k}
		}
	}
	return nil
}
