package removeDuplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	var cur int
	for i := 0; i < len(nums); i++ {
		if i + 1 < len(nums)&& nums[i] < nums[i + 1]{
			nums[cur] = nums[i]
			cur++
		}
	}
	if nums[cur] < nums[len(nums) - 1]{
		nums[cur] = nums[len(nums) - 1]
	}
	cur++
	return cur
}

func TheBest(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for k, v := range nums {
		if k > 0 && nums[length-1] != v {
			nums[length] = v
			length++
		}
	}
	return length
}