package remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0{
		return 0
	}

	var i = 0
	for _, v := range nums{
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}