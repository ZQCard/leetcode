package search_insert_position

func searchInsert(nums []int, target int) int {
	var temp []int
	var length = len(nums)

	if length == 0{
		temp = append(temp, target)
		return 0
	}

	if target <= nums[0]{
		temp = append(temp, target)
		temp = append(temp, nums[:]...)
		return 0
	}

	if target > nums[length-1]{
		temp = append(temp, nums[:]...)
		temp = append(temp, target)
		return length
	}

	for i := 0; i < length - 1; i++ {
		if nums[i] < target && target <= nums[i + 1]{
			temp = append(temp, nums[0:i+ 1]...)
			temp = append(temp, target)
			temp = append(temp, nums[i+ 1:]...)
			return i + 1
		}
	}
	return 0
}

// 最优解 二分法
func TheBest(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	var middle int

	for left <= right {
		middle = (left + right) / 2
		if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			return middle
		}
	}
	return left
}
