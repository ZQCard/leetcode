package maximum_subarray

// 	Time Limit Exceeded 答案是对的,但是超时
func maximumSubArray(nums []int) int {
	length := len(nums)
	var maximum = nums[0]
	for i := 1; i <= length; i++ {
		for j := 0; j < length; j++ {
			if  j + i <= length{
				var res int
				temp := nums[j:j+i]
				for _, v := range temp{
					res += v
				}
				if maximum < res{
					maximum = res
				}
			}
		}
	}
	return maximum
}

func TheBest(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	copy(dp, nums)
	max := dp[0]
	for i := 1; i < size; i++ {
		if dp[i]+dp[i-1] > dp[i] {
			dp[i] += dp[i-1]

		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
